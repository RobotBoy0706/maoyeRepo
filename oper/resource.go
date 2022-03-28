package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"miao/Utils/DB"
	"miao/Utils/ErrCode"
	"miao/Utils/R"
	"miao/Utils/common"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	
	// "encoding/json"
	
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	// rootDefault   = "~/data/resource"
	envkey = "RESOURCE_STORE_PATH"
)

var (
	ResourceTable     string
	ResourceinfoTable string
	
	pathRoot string
)

type Resource struct {
	Id  int64  `json:"id" gorm:"index"`
	Rid string `json:"rid" form:"rid" gorm:"index"`
	// Rid      string `json:"rid" form:"rid" binding:"required" `
	// Type     string `json:"type" form:"type" binding:"required" `
	Type     string `json:"type" form:"type" `
	Filename string `json:"filename" form:"filename" gorm:"index"`
	Md5      string `json:"md5" form:"md5"  gorm:"index"`
	Resume   string `json:"resume" form:"resume"`
	Bindtime int64  `json:"bindtime"`
	Layerid  int    `json:"layer_id" form:"layer_id"  gorm:"-"`
	State    int    `json:"state" form:"state"  gorm:"-"`
	
	Layer string `json:"layer" form:"layer"  gorm:"-"`
}

type Resourceinfo struct {
	Md5        string `json:"md5" gorm:"index" gorm:"index"`
	Userid     int64  `json:"userid"` //上传者ID
	Createtime int64  `json:"createtime"`
}

func verifyMd5Sum(needSumData []byte, md5Value string) bool {
	
	//check md5sum from file
	//Buf, err := ioutil.ReadFile(fileName)
	//if err != nil {
	//	fmt.Println(err)
	//}
	md5fromfile := fmt.Sprintf("%x", md5.Sum(needSumData))
	
	if strings.EqualFold(md5fromfile, md5Value) {
		return true
	}
	
	return false
}

func getFileByForm(c *gin.Context, fileInfo *Resource) (ec int, em string) {
	
	err := c.Request.ParseMultipartForm(maxUploadSize)
	if err != nil {
		fmt.Println("file overlength: ", err)
		ec, em = ErrCode.GetReturnCode("INVALID_PARAM")
		return
	}
	
	var fileBuf []byte
	base64File := c.Request.PostFormValue("base64File")
	if base64File == "" {
		//得到上传的文件
		file, fileFs, err := c.Request.FormFile("file")
		if err != nil {
			fmt.Println("no key:'file' found: ", err)
			ec, em = ErrCode.GetReturnCode("OK")
			fileInfo.Md5 = ""
			return
		}
		if fileInfo.Filename != fileFs.Filename && fileInfo.Filename == "" {
			fileInfo.Filename = fileFs.Filename
		}
		
		fileBuf, err = ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("read file error, ", err)
			ec, em = ErrCode.GetReturnCode("INVALID_PARAM")
			return
		}
	} else {
		fileBuf, err = base64.StdEncoding.DecodeString(base64File)
		if err != nil {
			fmt.Println("no key:'base64file' found: ", err)
			ec, em = ErrCode.GetReturnCode("OK")
			fileInfo.Md5 = ""
			return
		}
		if fileInfo.Filename == "" {
			fileInfo.Filename = fileInfo.Md5
		}
	}
	if md5SumRes := verifyMd5Sum(fileBuf, fileInfo.Md5); !md5SumRes {
		ec, em = ErrCode.GetReturnCode("ERROR_MD5")
		return
	}
	
	//以文件Md5为文件名称保存文件
	storeFile := path.Join(pathRoot, fileInfo.Md5)
	
	// 如果文件存在，则不用保存
	// isExist, err := FileExists(storeFile)
	// if err != nil {
	// 	fmt.Println("FileExists API err: ", err)
	// 	return ErrCode.GetReturnCode("INTERNAL_ERROR")
	// }
	// if isExist {
	// 	fmt.Println("FileExists return")
	// 	return ErrCode.GetReturnCode("OK")
	// }
	
	// 文件不存在，保存
	err = ioutil.WriteFile(storeFile, fileBuf, 0664)
	if err != nil {
		fmt.Println("store file error: ", err)
		ec, em = ErrCode.GetReturnCode("INTERNAL_ERROR")
		return
	}
	
	ec, em = ErrCode.GetReturnCode("OK")
	
	return
}

func bindResource(id int64, rid string) {
	re := Resource{}
	DB.GetM(&re, id, ResourceTable)
	if re.Id == 0 {
		return
	}
	re.Rid = rid
	DB.UpdateByFieldM(&re, "id", id, ResourceTable)
}

/**
* @api POST /api/v1/resource/resource 新增一个资源
* 只有管理员能创建类型type为站内新闻的资源
* @apiGroup resource
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiRequest multipart/form-data
* @apiParam type  string 必填，资源类型，比如地图、线索、模组等，如果是站内新闻，则房间ID不填，type填"站内新闻"
* @apiParam rid  string 选填，房间id
* @apiParam filename  string 必填，资源文件名
* @apiParam md5  string 必填，资源md5
* @apiParam resume  string 选填，资源描述
* @apiParam file  file 选填，文件数据流
* @apiParam base64File  string 选填，base64文件数据流
* @apiParam layer_id  int 选填,图层ID
* @apiParam state  int 选填,槽位ID

* @apiSuccess 200 OK(上传过程会发回progress如：{"progress": "24%"})
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiParam data json id:资源ID，md5:资源md5，用这个值可以在后台获取文件数据
* @apiExample json
* {
    "errcode":0,
    "errmsg":"ok",
    "data":{"id":1, "md5":"134123412341234123412234123"}
}

* @apiError 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* 服务器内部错误
* {
    "errcode":30001,
    "errmsg":"system internal error"
* }
* @apiExample json
* 用户不合法
* {
    "errcode":30305,
    "errmsg":"auth failed"
* }
* @apiExample json
* 参数错误
* {
    "errcode":30002,
    "errmsg":"invalid parameter"
* }
* @apiExample json
* 操作数据库失败
* {
    "errcode":30009,
    "errmsg":"database operation error"
* }
*/
func PostResource(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	
	userid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	userid = 999999
	
	rq := Resource{}
	if err := c.ShouldBindWith(&rq, binding.Form); err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	if rq.Type == "站内新闻" {
		R.RJson(c, "NO_PERMISSION")
		return
	}
	
	ec, em := getFileByForm(c, &rq)
	if ec != 0 {
		R.Jr(c, http.StatusOK, gin.H{"errcode": ec, "errmsg": em})
		return
	}
	
	rq.Bindtime = time.Now().Unix()
	
	DB.CreateM(&rq, ResourceTable)
	
	rqis := []Resourceinfo{}
	if rq.Md5 != "" {
		DB.GetByFieldM(&rqis, "md5", rq.Md5, ResourceinfoTable)
		if len(rqis) == 0 {
			rqi := Resourceinfo{}
			
			rqi.Md5 = rq.Md5
			rqi.Createtime = time.Now().Unix()
			rqi.Userid = int64(userid)
			
			DB.CreateM(&rqi, ResourceinfoTable)
			rqis = append(rqis, rqi)
		}
	}
	
	R.RData(c, rq)
}

/**
* @api PUT /api/v1/resource/resource 修改一个资源
* 只有管理员能修改类型type为站内新闻的资源;只有管理员和房主才能修改地图
* @apiGroup resource
* @apiRequest json
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery id int 必填，执行需要修改的资源id

* @apiRequest multipart/form-data
* @apiParam type  string 必填，资源类型，比如地图、线索、模组等，如果是站内新闻，则房间ID不填，type填"站内新闻"
* @apiParam filename  string 必填，资源文件名
* @apiParam md5  string 必填，资源md5
* @apiParam resume  string 选填，资源描述
* @apiParam file  file 选填，文件数据流
* @apiParam base64File  string 选填，base64文件数据流
* @apiParam layer_id  int 选填，图层id
* @apiParam state  int 选填，槽id

* @apiSuccess 200 OK(上传过程会发回progress如：{"progress": "24%"})
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiParam data json id:资源ID，md5:资源md5，用这个值可以在后台获取文件数据
* @apiExample json
* {
    "errcode":0,
    "errmsg":"ok",
    "data":{"id":1, "md5":"134123412341234123412234123"}
}

* @apiError 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* 服务器内部错误
* {
    "errcode":30001,
    "errmsg":"system internal error"
* }
* @apiExample json
* 用户不合法
* {
    "errcode":30305,
    "errmsg":"auth failed"
* }
* @apiExample json
* 参数错误
* {
    "errcode":30002,
    "errmsg":"invalid parameter"
* }
* @apiExample json
* 操作数据库失败
* {
    "errcode":30009,
    "errmsg":"database operation error"
* }
*/
func PutResource(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	
	userid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	resourceid, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	userid = 999999
	queryResource := Resource{}
	DB.GetM(&queryResource, int64(resourceid), ResourceTable)
	if queryResource.Id == 0 {
		R.RJson(c, "NOT_FOUND")
		return
	}
	
	rq := Resource{}
	if err := c.ShouldBindWith(&rq, binding.Form); err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	if rq.Type == "站内新闻" || queryResource.Type == "站内新闻" {
		R.RJson(c, "NO_PERMISSION")
		return
	}
	
	if rq.Type != "" {
		queryResource.Type = rq.Type
	}
	if rq.Resume != "" {
		queryResource.Resume = rq.Resume
	}
	if rq.Md5 != "" {
		queryResource.Md5 = rq.Md5
		queryResource.Filename = rq.Filename
	}
	
	queryResource.Bindtime = time.Now().Unix()
	
	DB.UpdateByFieldM(&rq, "id", queryResource.Id, ResourceTable)
	
	rqis := []Resourceinfo{}
	if rq.Md5 != "" {
		DB.GetByFieldM(&rqis, "md5", rq.Md5, ResourceinfoTable)
		if len(rqis) == 0 {
			rqi := Resourceinfo{}
			
			rqi.Md5 = rq.Md5
			rqi.Createtime = time.Now().Unix()
			rqi.Userid = int64(userid)
			
			DB.CreateM(&rqi, ResourceinfoTable)
			rqis = append(rqis, rqi)
		}
	}
	
	R.RData(c, queryResource)
}

/**
* @api DELETE /api/v1/resource/resource 删除一个资源
* 只有管理员能删除类型type为站内新闻的资源
* @apiRequest query
* @apiGroup resource
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery id int64 必填，需要删除的资源的id
* @apiQuery layer_id int64 必填，需要删除的layer id
* @apiQuery state int64 必填，槽id


* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* {
    "errcode":0,
    "errmsg":"ok"
}

* @apiError 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* 服务器内部错误
* {
    "errcode":30001,
    "errmsg":"system internal error"
* }
* @apiExample json
* 用户不合法
* {
    "errcode":30305,
    "errmsg":"auth failed"
* }
* @apiExample json
* 参数错误
* {
    "errcode":30002,
    "errmsg":"invalid parameter"
* }
* @apiExample json
* 操作数据库失败
* {
    "errcode":30009,
    "errmsg":"database operation error"
* }
*/
func DeleteResource(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	
	resourceId := c.Query("id")
	if resourceId == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	resourceid, err := strconv.Atoi(resourceId)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	rq := Resource{}
	DB.GetM(&rq, int64(resourceid), ResourceTable)
	
	if rq.Type == "站内新闻" {
		R.RJson(c, "NO_PERMISSION")
		return
	}
	
	DB.DeleteM(int64(resourceid), ResourceTable)
	
	if rq.Md5 != "" {
		rqs := []Resource{}
		DB.GetByFieldM(&rqs, "md5", rq.Md5, ResourceTable)
		if len(rqs) == 0 {
			DB.Db.Table(ResourceinfoTable).Where("md5 = ?", rq.Md5).Delete(nil)
			storeFile := path.Join(pathRoot, rq.Md5)
			os.Remove(storeFile)
		}
	}
	
	R.RData(c, rq)
}



func ResourceInit() {
	DB.CreateTableM(Resource{})
	ResourceTable = DB.GetTableNameM(Resource{})
	
	DB.CreateTableM(Resourceinfo{})
	ResourceinfoTable = DB.GetTableNameM(Resourceinfo{})
	
	// 初始化文件保存路径
	pathRoot = Datadir
	fmt.Println("Data path: ", pathRoot)
	
	exist, _ := common.PathExists(pathRoot)
	if !exist {
		err := os.MkdirAll(pathRoot, 0711)
		if err != nil {
			fmt.Println("create root path error")
			os.Exit(1)
		}
	}
}



/**
* @api GET /api/v1/resource/data GET方法获取资源(支持Content-Range，如果resource的type是音乐的话，Content-Type会设置成audio/mpeg)
* @apiGroup resource
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiParam md5 string 必填，资源md5值
* @apiParam timestamp int 选填，请求时时间戳
* @apiParam download int 选填，是否下载，1为下载，0为不下载，默认为0
* @apiParam file_name string 选填，文件名


* @apiSuccess 200 OK
* @apiExample BIN
* 返回文件数据
 */
func GetResourceData(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}

	md5 := c.Query("md5")
	if md5 == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}

	fileName := c.Query("file_name")

	dflag := 0
	download := c.Query("download")
	if download == "0" || download == "" {
		dflag = 0
	} else {
		dflag = 1
	}

	rqis := []Resource{}
	DB.GetByFieldM(&rqis, "md5", md5, ResourceTable)
	if len(rqis) == 0 {
		rqis = append(rqis, Resource{Filename: fileName})
	}

	mime := ""
	filenameReal := rqis[0].Filename
	ext := path.Ext(filenameReal)
	if len(ext) != 0 && ext[0] == '.' {
		mime = GetMime(ext[1:])
	}

	storeFile := path.Join(pathRoot, md5)
	isExist, err := common.FileExists(storeFile)
	if err != nil {
		fmt.Println("FileExists API err: ", err)
		R.RJson(c, "INTERNAL_ERROR")
		return
	}
	if !isExist {
		R.RJson(c, "NOT_FOUND")
		return
	}

	// 获得文件长度
	fileInfo, err := os.Stat(storeFile)
	if err != nil {
		fmt.Println("Stat API err: ", err)
		R.RJson(c, "INTERNAL_ERROR")
		return
	}
	filesize := fileInfo.Size()

	header := c.Writer.Header()
	rg := c.Request.Header["Range"]
	start_byte := 0
	end_byte := 0
	fmt.Println(len(rg), rg)
	if len(rg) != 0 {
		fmt.Sscanf(rg[0], "bytes=%d-%d", &start_byte, &end_byte)
		c.Writer.WriteHeader(http.StatusPartialContent)
	}
	fmt.Println(start_byte, end_byte)
	if end_byte == 0 {
		end_byte = int(filesize) - 1
	}

	if rqis[0].Type == "音乐" || rqis[0].Type == "music" {
		if mime != "" {
			header["Content-Type"] = []string{mime}
		} else {
			header["Content-Type"] = []string{"audio/mpeg"}
		}
	} else if dflag == 1 {
		header["Content-Type"] = []string{"application/octet-stream"}
	} else {
		header["Content-Type"] = []string{mime}
	}
	header["Content-Length"] = []string{fmt.Sprintf("%d", int64(end_byte-start_byte+1))}
	if dflag == 1 {
		header["Content-Disposition"] = []string{fmt.Sprintf(`attachment; filename="%s"`, filenameReal)}
	}
	if end_byte > start_byte && end_byte != 0 && int64(end_byte) < filesize {
		sss := fmt.Sprintf("bytes %d-%d/%d", start_byte, end_byte, filesize)
		header["Content-Range"] = []string{sss}
		header["Accept-Ranges"] = []string{"bytes"}
	}

	file, err := os.Open(storeFile)
	if err != nil {
		fmt.Println("Open API err: ", err)
		R.RJson(c, "INTERNAL_ERROR")
		return
	}
	defer file.Close()

	file.Seek(int64(start_byte), 0)
	io.CopyN(c.Writer, file, int64(end_byte-start_byte+1))
}