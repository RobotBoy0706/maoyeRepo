package main

import (
	"crypto/md5"
	_ "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/globalsign/mgo/bson"
	"io"
	"io/ioutil"
	"miao/Utils/DB"
	"miao/Utils/ErrCode"
	"miao/Utils/R"
	"miao/Utils/common"
	"net/http"
	"os"
	"path"
	_ "regexp"
	"strconv"
	_ "strconv"
	"strings"
	_ "sync"
)

var (
	AdvertisementTable string
	picturepath        string
)

const (
	maxUploadSize = 50 * 1024 * 1024 // File size max: 50MB
	// rootDefault   = "~/data/resource"
	//envkey = "RESOURCE_STORE_PATH"
)



type Advertisement struct {
	Id          int64  `json:"id" gorm:"index"`
	Title       string `json:"title" gorm:"index"`
	Text        string `json:"text"`
	Type        string `json:"type"`
	Picturename string `json:"picturename" form:"picturename"`
	Md5         string `json:"md5" form:"md5" gorm:"index"`

	Starttime string `json:"starttime"`
	Endtime   string `json:"endtime"`
}

/**
* @api POST /api/v1/advertisement/advertisement 新增广告
* @apiGroup advertisement
* @apiQuery userid int 必填，执行请求的用户id（只有管理员用户才能新增,管理员id在config.json中admin处设置）
* @apiQuery token string 必填，执行请求的用户token

* @apiRequest multipart/form-data
* @apiParam title  string 必填，标题
* @apiParam text  string 必填，文本
* @apiParam md5  string 必填，资源md5
* @apiParam file  file 必填，文件数据流
* @apiParam starttime string 必填 开始生效时间 必须用2021-01-02 15:04:05的格式
* @apiParam endtime string 必填 结束生效时间 必须用2021-01-02 15:04:05的格式
* @apiParam type string 必填 类型 广告/公告
*{
    "data": {
        "id": 28,
        "title": "abcd",
        "text": "广告，公告管理系统",
        "type": "广告",
        "picturename": "接口信息.doc",
        "md5": "f77e06096252c8a27236f137e0f0dd89",
        "starttime": "2021-01-02 15:04:05",
        "endtime": "2022-01-02 15:04:05"
    },
    "errcode": 0,
    "errmsg": "ok"
*}

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
* 不是管理员
* {
    "errcode":30005,
    "errmsg":"no permission"
* }
* @apiExample json
* 操作数据库失败
* {
    "errcode":30009,
    "errmsg":"database operation error"
* }
*/

func PostAdvertisement(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	//userid, err := strconv.Atoi(c.Query("userid"))
	//if err != nil {
	//	R.RJson(c, "INVALID_PARAM")
	//	return
	//}
	//if !IsAdmin(int64(userid)) {
	//	R.RJson(c, "NO_PERMISSION")
	//	return
	//}
	ad := Advertisement{}
	if err := c.ShouldBindWith(&ad, binding.Form); err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}

	ec, em := getPictureByForm(c, &ad)
	if ec != 0 {
		R.Jr(c, http.StatusOK, gin.H{"errcode": ec, "errmsg": em})
		return
	}
	ad.Title = c.Request.FormValue("title")
	ad.Text = c.Request.FormValue("text")
	ad.Type = c.Request.FormValue("type")
	//以这种格式"2006-01-02 15:04:05"
	ad.Starttime = c.Request.FormValue("starttime")
	ad.Endtime = c.Request.FormValue("endtime")

	err := DB.CreateM(&ad, AdvertisementTable)
	if err != nil {
		DB.Db.Table(AdvertisementTable).Where("`id`=?", ad.Id).Delete(nil)
		R.RJson(c, "ERROR_DB")
		return
	}
	R.RData(c, ad)

}

/**
* @api PUT /api/v1/advertisement/advertisement 新增广告
* @apiGroup advertisement
* @apiQuery userid int 必填，执行请求的用户id（只有管理员用户才能新增,管理员id在config.json中admin处设置）
* @apiQuery token string 必填，执行请求的用户token

* @apiRequest multipart/form-data
* @apiParam title  string 必填，标题
* @apiParam text  string 必填，文本
* @apiParam md5  string 必填，资源md5
* @apiParam file  file 必填，文件数据流
* @apiParam starttime string 必填 开始生效时间 必须用2021-01-02 15:04:05的格式
* @apiParam endtime string 必填 结束生效时间 必须用2021-01-02 15:04:05的格式
* @apiParam type string 必填 类型 广告/公告
*{
    "data": {
        "id": 28,
        "title": "abcd",
        "text": "广告，公告管理系统",
        "type": "广告",
        "picturename": "接口信息.doc",
        "md5": "f77e06096252c8a27236f137e0f0dd89",
        "starttime": "2021-01-02 15:04:05",
        "endtime": "2022-01-02 15:04:05"
    },
    "errcode": 0,
    "errmsg": "ok"
*}

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
* 不是管理员
* {
    "errcode":30005,
    "errmsg":"no permission"
* }
* @apiExample json
* 操作数据库失败
* {
    "errcode":30009,
    "errmsg":"database operation error"
* }
*/
func PutAdvertisement(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	//userid, err := strconv.Atoi(c.Query("userid"))
	//if err != nil {
	//	R.RJson(c, "INVALID_PARAM")
	//	return
	//}
	//if !IsAdmin(int64(userid)) {
	//	R.RJson(c, "NO_PERMISSION")
	//	return
	//}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	quretad := Advertisement{}
	DB.GetM(&quretad, int64(id), AdvertisementTable)
	if quretad.Id == 0 {
		R.RJson(c, "NOT_FOUND")
		return
	}
	ad := Advertisement{}
	if err := c.ShouldBindWith(&ad, binding.Form); err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	ec, em := getPictureByForm(c, &ad)
	if ec != 0 {
		R.Jr(c, http.StatusOK, gin.H{"errcode": ec, "errmsg": em})
		return
	}
	ad.Id = int64(id)
	ad.Title = c.Request.FormValue("title")
	ad.Text = c.Request.FormValue("text")
	ad.Type = c.Request.FormValue("type")
	//以这种格式"2006-01-02 15:04:05"存
	ad.Starttime = c.Request.FormValue("starttime")
	ad.Endtime = c.Request.FormValue("endtime")

	DB.UpdateByFieldM(&ad, "id", quretad.Id, AdvertisementTable)
	if quretad.Md5 != "" {
		storeFile := path.Join(picturepath, quretad.Md5)
		os.Remove(storeFile)
	}
	R.RData(c, ad)
}

/**
* @api DELETE /api/v1/advertisement/advertisement 删除广告
* @apiGroup advertisement
* @apiQuery userid int 必填，执行请求的用户id（只有管理员用户才能删除,管理员id在config.json中admin处设置）
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery id string 必填，要删除的广告id

*{
    "data": {
        "id": 28,
        "title": "abcd",
        "text": "广告，公告管理系统",
        "type": "广告",
        "picturename": "接口信息.doc",
        "md5": "f77e06096252c8a27236f137e0f0dd89",
        "starttime": "2021-01-02 15:04:05",
        "endtime": "2022-01-02 15:04:05"
    },
    "errcode": 0,
    "errmsg": "ok"
*}

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
* 不是管理员
* {
    "errcode":30005,
    "errmsg":"no permission"
* }
* @apiExample json
* 操作数据库失败
* {
    "errcode":30009,
    "errmsg":"database operation error"
* }
*/
func DeleteAdvertisement(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	//userid, err := strconv.Atoi(c.Query("userid"))
	//if err != nil {
	//	R.RJson(c, "INVALID_PARAM")
	//	return
	//}
	//if !IsAdmin(int64(userid)) {
	//	R.RJson(c, "NO_PERMISSION")
	//	return
	//}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	ad := Advertisement{}
	DB.GetM(&ad, int64(id), AdvertisementTable)
	DB.DeleteM(int64(id), AdvertisementTable)
	if ad.Md5 != "" {
		storeFile := path.Join(picturepath, ad.Md5)
		os.Remove(storeFile)
	}
	R.RData(c, ad)
}

/**
* @api GET /api/v1/advertisement/advertisement 获取广告
* @apiGroup advertisement
* @apiQuery userid int 必填，执行请求的用户id（这里可以用非管理员id，不过非管理员只能看到未过期的广告，管理员可以看到全部广告,管理员id在config.json中admin处设置）
* @apiQuery token string 必填，执行请求的用户token
* 返回的是请求有权限看到的广告列表
*{
    "data": {
        "id": 28,
        "title": "abcd",
        "text": "广告，公告管理系统",
        "type": "广告",
        "picturename": "接口信息.doc",
        "md5": "f77e06096252c8a27236f137e0f0dd89",
        "starttime": "2021-01-02 15:04:05",
        "endtime": "2022-01-02 15:04:05"
    },
    "errcode": 0,
    "errmsg": "ok"
*}
*{
    "data": {
        "id": 30,
        "title": "abcd",
        "text": "广告，公告管理系统",
        "type": "广告",
        "picturename": "接口信息.doc",
        "md5": "f77e06096252c8a27236f137e0f0dd89",
        "starttime": "2021-01-02 15:04:05",
        "endtime": "2022-01-02 15:04:05"
    },
    "errcode": 0,
    "errmsg": "ok"
*}

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
func GetAdvertisement(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	//userid, err := strconv.Atoi(c.Query("userid"))
	//if err != nil {
	//	R.RJson(c, "INVALID_PARAM")
	//	return
	//}
	isAdmin := true
	//if !IsAdmin(int64(userid)) {
	//	isAdmin = false
	//}
	ads := []Advertisement{}
	DB.GetAllM(&ads, AdvertisementTable)
	if isAdmin {
		R.RData(c, ads)
	} /* else {
		ads2 := []Advertisement{}
		for i := 0; i < len(ads); i++ {
			nowtime := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
			//endtime, _ :=time.ParseInLocation("2006-01-02 15:04:05", bt.Endtime, time.Local)
			t1, err1 := time.Parse("2006-01-02 15:04:05", nowtime)
			t2, err2 := time.Parse("2006-01-02 15:04:05", ads[i].Endtime)
			if err1 == nil && err2 == nil && t2.After(t1) {
				ads2 = append(ads2, ads[i])
			}
		}
		R.RData(c, ads2)
	}*/
	return
}

/**
* @api GET /api/v1/advertisement/advertisementPicture 获取广告图片
* @apiGroup advertisement
* @apiQuery userid int 必填，执行请求的用户id（这里可以用非管理员id，不过非管理员只能看到未过期的广告，管理员可以看到全部广告,管理员id在config.json中admin处设置）
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery md5 string 必填，图片的md5值，这个可以通过上面的获取广告信息中md5得到

* 返回的是图片

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
func GetAdvertisementPicture(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	md5 := c.Query("md5")
	if md5 == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	storeFile := path.Join(picturepath, md5)
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
	header["Content-type"] = []string{"application/octet-stream"}
	header["Content-Length"] = []string{fmt.Sprintf("%d", filesize)}
	header["Content-Disposition"] = []string{fmt.Sprintf(`attachment; filename="%s"` + path.Base(storeFile))}

	file, err := os.Open(storeFile)
	if err != nil {
		fmt.Println("Open API err: ", err)
		R.RJson(c, "INTERNAL_ERROR")
		return
	}
	defer file.Close()

	io.Copy(c.Writer, file)

}

func verifyMd5Sum2(needSumData []byte, md5Value string) bool {
	md5fromfile := fmt.Sprintf("%x", md5.Sum(needSumData))

	if strings.EqualFold(md5fromfile, md5Value) {
		return true
	}

	return false
}

func getPictureByForm(c *gin.Context, fileInfo *Advertisement) (ec int, em string) {
	//根据resource的getFileByForm改的
	err := c.Request.ParseMultipartForm(maxUploadSize)
	if err != nil {
		fmt.Println("file overlength: ", err)
		ec, em = ErrCode.GetReturnCode("INVALID_FILE")
		return
	}

	//得到上传的文件
	file, fileFs, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println("no key:'file' found: ", err)
		ec, em = ErrCode.GetReturnCode("INVALID_FILE")
		fileInfo.Md5 = ""
		return
	}

	fileInfo.Picturename = fileFs.Filename

	fileBuf, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("read file error, ", err)
		ec, em = ErrCode.GetReturnCode("INTERNAL_ERROR")
		return
	}
	if md5SumRes := verifyMd5Sum2(fileBuf, fileInfo.Md5); !md5SumRes {
		ec, em = ErrCode.GetReturnCode("ERROR_MD5")
		return
	}

	//以文件Md5为文件名称保存文件
	storeFile := path.Join(picturepath, fileInfo.Md5)

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

func AdvertisementInit() {
	DB.CreateTableM(Advertisement{})
	AdvertisementTable = DB.GetTableNameM(Advertisement{})
	DB.Db.Table(AdvertisementTable).AutoMigrate(&Advertisement{})

	// 初始化文件保存路径
	picturepath = Picturedir
	fmt.Println("Data path: ", picturepath)

	exist, _ := common.PathExists(picturepath)
	if !exist {
		err := os.MkdirAll(picturepath, 0711)
		if err != nil {
			fmt.Println("create root path error")
			os.Exit(1)
		}
	}
}
