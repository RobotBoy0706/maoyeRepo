package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"miao/Utils/DB"
	"miao/Utils/R"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var (
	rolecardidmap      map[string]int
	rolecardidmapMutex sync.Mutex
	
	npccardidmap      map[string]int
	npccardidmapMutex sync.Mutex
)

const (
	rolecardidfile = "./id"
)

func GetRolecardid(rg string) string {
	rolecardidmapMutex.Lock()
	last_id, ok := rolecardidmap[rg]
	if !ok {
		last_id = 1
	} else {
		last_id += 1
	}
	rolecardidmap[rg] = last_id
	
	SaveConfig(rolecardidmap)
	rolecardidmapMutex.Unlock()
	
	ret := fmt.Sprintf("%s_%d", rg, last_id)
	
	return ret
}

func GetNpccardid(rg string) string {
	npccardidmapMutex.Lock()
	last_id, ok := npccardidmap[rg]
	if !ok {
		last_id = 1
	} else {
		last_id += 1
	}
	npccardidmap[rg] = last_id
	
	SaveConfig(npccardidmap)
	npccardidmapMutex.Unlock()
	
	ret := fmt.Sprintf("%s_%d", rg, last_id)
	
	return ret
}

/**
* @api POST /api/v1/rolecard/info 新增一个技能/武器/职业/角色卡(body参数的userid如果不传，则填为跟query参数的userid一样)
* 只有管理员能创建公用的,只有管理员能为他人创建
* @apiGroup rolecard
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery name string 必填，请求的表名称，比如角色卡：rolecard；NPC卡：npccard；技能：skill; 武器:weapon; 职业：job； 技能子选项: skselect 不同表决定了body数据中的字段
* @apiQuery rg string 选填，如果name为rolecard时，则为必填，选择角色卡的规则，比如coc7

* @apiRequest json
* @apiExample json
如果为skill：
	name string
	ini string
	grow string
	pro string
	interest string
	total string
	num int
	kind string
	bz int
	introduce string
	sub string  子项名称，如kexue
	userid int64  所属用户，如果是公用的，则填0
如果为job：
	value string
	name string
	skills []json  技能集合,和skill表字段一样
	ext int 自定义技能数量
	fourt []json 具有如下字段:
		onum string
		okind string
		tnum string
		tkind string
		hnum string
		hkind string
		fnum string
		fkind string
	intro []json  介绍说明，具有下面三个字段：
		honesty string
		propoint string
		proskill
	pro [8]int 影响系数，8个值的列表
	userid int64  所属用户，如果是公用的，则填0
如果为weapons：
	name string
	skill string
	dam string
	tho string
	range string
	round string
	num string
	price string
	err int
	time string
	type string   类型，比如cg：常规，sq：手枪等
	userid int64  所属用户，如果是公用的，则填0
如果为rolecard
	job json 职业信息，和job表字段一样
	bz []json 技能信息，和skill字段一样
	jobwt []string
	selskval []string
	name json  有如下字段
		touxiang     (可以直接把角色卡的头像的文件内容保存在这里，比如base64)
		job
		jobval
		player
		chartname
		time
		ages
		sex
		address
		hometown
	attribute json 有如下字段
		str,
		con,
		siz,
		dex,
		app,
		int1,
		pow,
		edu,
		luck,
		mov,
		build,
		db
	hp json 有如下字段
		have string
		total string
	mp  json 有如下字段
		have string
		total string
	san json 有如下字段
		have string
		total string
	weapons  []string
	things []string
	money []string
	story json 具有如下字段
		miaoshu string
		xinnian string
		zyzr string
		feifanzd string
		bgzw string
		tedian string
		bahen string
		kongju string
		story string
	more json 有如下字段
		jingli string
		huoban string
		kesulu string
	drawpic
		attrpic []json
		namepic []json
		skillspic []json
		thingspic []json
		storypic []json
		story string
		touxiang string
	mind string
	health string
	touniang string
	chartid     //这个字段由后端自动分配，创建成果后会返回#格式为规则简称_顺延卡号,比如coc7_2427,代表着这是一张coc7版规则下的角色卡，角色卡号是2427。如果玩家把卡转移转移给别人，自己的卡不会消失，同时别人会生成一张coc7_2846的角色卡，简单来说就是卡的赠予之后，角色卡号会更新，单个卡号就对应一个用户就好"

* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiParam data json 新建立的信息，如果是rolecard，则有chartid字段表示分配到的id，其他的如skill等用_id作为唯一标识
* @apiExample json
* {
    "errcode":0,
	"errmsg":"ok",
	"data":{
		...
	}
}

* @apiError 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* 参数错误
* {
    "errcode":30002,
    "errmsg":"invalid parameter"
* }
* @apiExample json
* 数据库操作失败
* {
    "errcode":30009,
    "errmsg":"database operation error"
* }
*/

func PostRoleInfo(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	
	name := c.Query("name")
	if name == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	} else if name != "job" && name != "skill" && name != "rolecard" && name != "skselect" && name != "weapon" && name != "npccard" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	id := c.Query("userid")
	if id == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	userid, err := strconv.Atoi(id)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	var sampleInfo map[string]interface{}
	err = c.BindJSON(&sampleInfo)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	sampleInfo["_id"] = bson.NewObjectId()
	
	useridString, ok := sampleInfo["userid"].(string)
	if ok {
		if useridString != "" {
			uString, err := strconv.Atoi(useridString)
			if err != nil {
				fmt.Println("userid is string, and can not change to int: ", useridString)
				R.RJson(c, "INVALID_PARAM")
				return
			}
			sampleInfo["userid"] = uString
		} else {
			sampleInfo["userid"] = userid
		}
	} else {
		//useridInt, ok := sampleInfo["userid"].(float64)
		//if !ok {
		//	sampleInfo["userid"] = userid
		//} else {
		//	if int(useridInt) == 0 && !IsAdmin(int64(userid)) {
		//		R.RJson(c, "NO_PERMISSION")
		//		return
		//	}
		//	if int(useridInt) != 0 && int(useridInt) != userid && !IsAdmin(int64(userid)) {
		//		R.RJson(c, "NO_PERMISSION")
		//		return
		//	}
		//}
	}
	
	if name == "npccard" {
		rg := c.Query("rg")
		if rg == "" {
			R.RJson(c, "INVALID_PARAM")
			return
		}
		sampleInfo["npcid"] = GetNpccardid(rg)
	}
	
	if name == "rolecard" {
		rg := c.Query("rg")
		if rg == "" {
			R.RJson(c, "INVALID_PARAM")
			return
		}
		sampleInfo["chartid"] = GetRolecardid(rg)
	}
	
	err = DB.MongoCreate(Mgodb, name, sampleInfo)
	if err != nil {
		R.RJson(c, "ERROR_DB")
		return
	}
	
	R.RData(c, sampleInfo)
}

type toolids struct {
	Ids []string `json:"ids"`
}

/**
* @api DELETE /api/v1/rolecard/info 删除一个技能/武器/职业/角色卡(每个人只能删除自己的,管理员可以删除公用及其他人的)
* @apiGroup rolecard
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery name string 必填，请求的表名称，比如角色卡：rolecard；技能：skill; 武器:weapon; 职业：job；NPC卡：npccard； 技能子选项: skselect 不同表决定了body数据中的字段

* @apiRequest json
* @apiParam ids array 必填，要删除的id集
* @apiExample json
* {
    "ids":["123456789012345678901234","123456789012345678901235"]
* }

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
* 参数错误
* {
    "errcode":30002,
    "errmsg":"invalid parameter"
* }
* @apiExample json
* 数据库操作失败
* {
    "errcode":30009,
    "errmsg":"database operation error"
* }
*/
func DeleteRoleInfo(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	
	name := c.Query("name")
	if name == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	} else if name != "job" && name != "skill" && name != "rolecard" && name != "skselect" && name != "weapon" && name != "npccard" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	id := c.Query("userid")
	if id == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	userid, err := strconv.Atoi(id)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	tools := toolids{}
	err = c.Bind(&tools)
	if err != nil || len(tools.Ids) == 0 {
		fmt.Println(tools)
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	for _, id := range tools.Ids {
		query := bson.M{"_id": bson.ObjectIdHex(id), "userid": userid}
		if true {
			query = bson.M{"_id": bson.ObjectIdHex(id)}
		}
		fmt.Println(query)
		err := DB.MongoDelete(Mgodb, name, query)
		if err != nil {
			fmt.Println(err)
			// R.RJson(c, "ERROR_DB")
			continue
		}
	}
	
	R.RJson(c, "OK")
}

/**
* @api POST /api/v1/rolecard/get 按特征查询某个一个技能/武器/职业/角色卡（每个人最多只能查到自己的和公用的）
* @apiGroup rolecard
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery name string 必填，请求的表名称，比如角色卡：rolecard；技能：skill; 武器:weapon; 职业：job； 技能子选项: skselect 不同表决定了body数据中的字段
* @apiQuery flag int  选填，0：查询到自己的和公用的； 1：只查询公用的；2：只查询自己的
* @apiQuery rid string  选填，房间id，如果name是rolecard的话，该字段用于房主可以查看玩家的角色卡

* @apiRequest json
* @apiParam ... json 必填，查询的特征，需要是json，且相应的表有对应的字段
* @apiExample json
下面查询name为"会计"的技能
* {
    "name":"会计"
* }

* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiParam data json 查询到的
* @apiExample json
* {
    "errcode":0,
    "errmsg":"ok"
}

* @apiError 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* 参数错误
* {
    "errcode":30002,
    "errmsg":"invalid parameter"
* }
* @apiExample json
* 数据库操作失败
* {
    "errcode":30009,
    "errmsg":"database operation error"
* }
*/
func GetRoleInfo(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	
	name := c.Query("name")
	if name == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	} else if name != "job" && name != "skill" && name != "rolecard" && name != "skselect" && name != "weapon" && name != "npccard" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	id := c.Query("userid")
	if id == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	userid, err := strconv.Atoi(id)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	flag := c.Query("flag")
	if flag == "" {
		flag = "0"
	}
	
	var sampleInfo map[string]interface{}
	err = c.BindJSON(&sampleInfo)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	var rets []interface{}
	
	if flag == "1" || flag == "0" {
		var ret []interface{}
		sampleInfo["userid"] = nil
		err := DB.MongoGetAll(Mgodb, name, sampleInfo, &ret)
		if err != nil {
			R.RJson(c, "ERROR_DB")
			return
		}
		rets = append(rets, ret...)
	}
	if flag == "2" || flag == "0" {
		var ret []interface{}
		sampleInfo["userid"] = userid
		err := DB.MongoGetAll(Mgodb, name, sampleInfo, &ret)
		if err != nil {
			R.RJson(c, "ERROR_DB")
			return
		}
		rets = append(rets, ret...)
	}
	if len(rets) == 0 {
		rid := c.Query("rid")
		// 判断是否是房主并且要操作的角色卡在房间中
		if rid != "" && name == "rolecard" {
			rs := []Room{}
			DB.GetByFieldM(&rs, "id", rid, RoomTable)
			if len(rs) == 0 {
				R.RJson(c, "NOT_FOUND")
				return
			}
			
			chartid, ok := sampleInfo["chartid"]
			if ok && rs[0].Userid == int64(userid) {
				rr := []Roomrole{}
				sql_cmd := fmt.Sprintf("SELECT * FROM %s WHERE rid = '%s' AND JSON_CONTAINS(roleid, '[\"%v\"]');", RoomroleTable, rid, chartid)
				fmt.Println(sql_cmd)
				err := DB.Db.Raw(sql_cmd).Scan(&rr)
				if err.Error != nil {
					fmt.Println("find chart error: ", err.Error)
				}
				if len(rr) != 0 {
					delete(sampleInfo, "userid")
					fmt.Println("find chart: ", sampleInfo)
					var ret []interface{}
					err := DB.MongoGetAll(Mgodb, name, sampleInfo, &ret)
					if err != nil {
						fmt.Println("find chart error: ", err)
						R.RJson(c, "ERROR_DB")
						return
					}
					rets = append(rets, ret...)
					fmt.Println("find chart: ", ret)
				} else {
					fmt.Println("find chart empty")
				}
			}
		}
	}
	
	R.RData(c, rets)
}

/**
* @api POST /api/v1/rolecard/share 按特征查询某个一个技能/武器/职业/角色卡（不检验userid和token）
* @apiGroup rolecard
* @apiQuery name string 必填，请求的表名称，比如角色卡：rolecard(默认)；技能：skill; 武器:weapon; 职业：job； 技能子选项: skselect 不同表决定了body数据中的字段

* @apiRequest json
* @apiParam ... json 必填，查询的特征，需要是json，且相应的表有对应的字段
* @apiExample json
下面查询name为"会计"的技能
* {
    "name":"会计"
* }

* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiParam data json 查询到的信息
* @apiExample json
* {
	"data":{...}
    "errcode":0,
    "errmsg":"ok"
}

* @apiError 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* 参数错误
* {
    "errcode":30002,
    "errmsg":"invalid parameter"
* }
* @apiExample json
* 数据库操作失败
* {
    "errcode":30009,
    "errmsg":"database operation error"
* }
*/
func GetRoleShare(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	} else if name != "job" && name != "skill" && name != "rolecard" && name != "skselect" && name != "weapon" && name != "npccard" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	var sampleInfo map[string]interface{}
	err := c.BindJSON(&sampleInfo)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	var rets []interface{}
	
	var ret []interface{}
	err = DB.MongoGetAll(Mgodb, name, sampleInfo, &ret)
	if err != nil {
		R.RJson(c, "ERROR_DB")
		return
	}
	rets = append(rets, ret...)
	
	R.RData(c, rets)
}

/**
* @api GET /api/v1/rolecard/count 获取记录数量
* @apiGroup rolecard
* @apiQuery userid  int	必填，用户ID
* @apiQuery token   string	必填，登录token
* @apiQuery name string 必填，请求的表名称，比如角色卡：rolecard；技能：skill; 武器:weapon; 职业：job； 技能子选项: skselect 不同表决定了body数据中的字段
* @apiQuery flag int  选填，0：查询到自己的和公用的； 1：只查询公用的；2：只查询自己的
* @apiQuery filter string 选填，过滤的关键词：例如{"status":"idle","machineid":1}

* @apiSuccess 200 OK
* @apiParam errcode int     错误代码
* @apiParam errmsg  string  错误信息
* @apiParam data   int     查询到的记录数量
* @apiExample json
* {
*	   "errcode": 0,
*	   "errmsg": "ok",
*	   "data": 10
* }

* @apiError 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* 接口所需要的参数有误
* {"errcode":30002,"errmsg":"invalid parameter"}
* @apiExample json
* 操作数据库出现错误
* {"errcode":30001,"errmsg":"system internal error"}
 */
func GetCount(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	
	name := c.Query("name")
	if name == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	} else if name != "job" && name != "skill" && name != "rolecard" && name != "skselect" && name != "weapon" && name != "npccard" && name != "article" && name != "invitation" && name != "article_comment" && name != "invitation_comment" && name != "report"&& name != "manager_report" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	flag := c.Query("flag")
	if flag == "" {
		flag = "0"
	}
	
	needSearchCharMap := make(map[string]interface{})
	
	filter := c.Query("filter")
	if filter != "" {
		f := make(map[string]interface{})
		err := json.Unmarshal([]byte(filter), &f)
		if err == nil {
			switch name {
			case "invitation_comment", "article_comment":
				bussID, ok := f["buss_id"].(string)
				if !ok {
					R.RJson(c, "INVALID_PARAM")
					return
				}
				f["buss_id"] = bson.ObjectIdHex(bussID)
			}
			for k, v := range f {
				if k == "_id" {
					id_key_str, ok := v.(string)
					if !ok {
						R.RJson(c, "INVALID_PARAM")
						return
					}
					needSearchCharMap[k] = bson.ObjectIdHex(id_key_str)
				} else {
					needSearchCharMap[k] = v
				}
			}
		} else {
			fmt.Printf("decode fileter err=%v\n", err)
		}
	}
	
	//needSearchCharMap["type"]= "精品"
	//fmt.Printf("name=%v,needSearchCharMap=%#v,filter=%s\n", name, needSearchCharMap, filter)
	n, err := DB.MongoCount(Mgodb, name, needSearchCharMap)
	if err != nil {
		R.RJson(c, "ERROR_DB")
		return
	}
	
	R.RData(c, n)
}

/**
* @api GET /api/v1/rolecard/list 获取工具列表按特征查询某个一个技能/武器/职业/角色卡（每个人最多只能查到自己的和公用的）
* @apiGroup rolecard
* @apiQuery userid  int	必填，用户ID
* @apiQuery token   string	必填，登录token
* @apiQuery name string 必填，请求的表名称，比如角色卡：rolecard；技能：skill; 武器:weapon; 职业：job； 技能子选项: skselect 不同表决定了body数据中的字段
* @apiQuery position   string  取上一次查询结果数组中第一个元素的“_id”作为参数
* @apiQuery page   int  进行操作所在的页
* @apiQuery nextpage   int  需要跳转的页
* @apiQuery pagesize   int  必填，需要返回的数据条数
* @apiQuery father string 选填，可以是product、order等（表示把这个工具放到哪一个对象下面），如果不填则为product
* @apiQuery filter string 选填，过滤的关键词：例如name=rolecard，filter为{"status":"idle","userid":1}可以查到用户1的所拥有的角色卡

* @apiSuccess 200 OK
* @apiParam errcode int     错误代码
* @apiParam errmsg  string  错误信息
* @apiParam data   int     查询到的工具记录数量
* @apiExample json
* {
*	   "errcode": 0,
*	   "errmsg": "ok",
*	   "data": [
            {
				//...
            }
        ]
* }

* @apiError 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* 接口所需要的参数有误
* {"errcode":30002,"errmsg":"invalid parameter"}
* @apiExample json
* 操作数据库出现错误
* {"errcode":30001,"errmsg":"system internal error"}
* /

func GetList(c *gin.Context) {
	//if Check(c) == false {
	//	R.RJson(c, "AUTH_FAILED")
	//	return
	//}
	
	name := c.Query("name")
	if name == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	} else if name != "job" && name != "skill" && name != "rolecard" && name != "skselect" && name != "weapon" && name != "npccard" && name != "article" && name != "invitation" && name != "article_comment" && name != "invitation_comment" && name != "report" && name != "manager_report" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	// index := c.Query("index") // []string
	position := c.Query("position")
	page_str := c.Query("page")
	nextpage_str := c.Query("nextpage")
	pagesize_str := c.Query("pagesize")
	fullText := c.Query("full_text")
	sort := c.Query("sort")
	if sort == "" {
		sort = "-create_time"
	}
	
	if name == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	page, err := strconv.Atoi(page_str)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	nextpage, err := strconv.Atoi(nextpage_str)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	pagesize, err := strconv.Atoi(pagesize_str)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	needSearchCharMap := make(map[string]interface{})
	db := Mgodb.C(name)
	if fullText != "" {
		//err = db.EnsureIndexKey("$text:$**")//{"name":{$regex:"83"}}
		//if err != nil {
		//	fmt.Printf("EnsureIndexKey fail err=%v",err)
		//}
		needSearchCharMap["title"] = bson.M{
			"$regex": fullText,
		}
	}
	filter := c.Query("filter")
	if filter != "" {
		f := make(map[string]interface{})
		err := json.Unmarshal([]byte(filter), &f)
		if err == nil {
			switch name {
			case "invitation_comment", "article_comment":
				bussID, ok := f["buss_id"].(string)
				if !ok {
					R.RJson(c, "INVALID_PARAM")
					return
				}
				f["buss_id"] = bson.ObjectIdHex(bussID)
			}
			for k, v := range f {
				if k == "_id" {
					id_key_str, ok := v.(string)
					if !ok {
						R.RJson(c, "INVALID_PARAM")
						return
					}
					needSearchCharMap[k] = bson.ObjectIdHex(id_key_str)
				} else {
					needSearchCharMap[k] = v
				}
			}
		}
	}
	fmt.Printf("needSearchCharMap=%#v\n", needSearchCharMap)
	var sampleInfoArr []map[string]interface{}
	sortField := make([]string, 0)
	var skip int
	if page-nextpage <= 0 {
		if nextpage-1 < 0 {
			skip = 0
		} else {
			skip = (nextpage - 1) * pagesize
		}
		sortField = append(sortField, sort)
		if position != "" && position != "0" {
			needSearchCharMap["_id"] = map[string]interface{}{"$lte": bson.ObjectIdHex(position)}
		}
		
		err = db.Find(needSearchCharMap).
			Skip(skip).
			Limit(pagesize).Sort(sortField...).All(&sampleInfoArr)
	} else if page-nextpage > 0 {
		if nextpage-1 < 0 {
			skip = 0 * pagesize
		} else {
			skip = (nextpage - 1) * pagesize
		}
		sortField = append(sortField, sort)
		if position != "" {
			needSearchCharMap["_id"] = map[string]interface{}{"$gt": bson.ObjectIdHex(position)}
		}

		err = db.Find(needSearchCharMap).
			Skip(skip).
			Limit(pagesize).Sort(sortField...).All(&sampleInfoArr)
		
		//因排序倒置,倒转数组后再发送,统一操作逻辑
		infoStart := 0
		infoEnd := len(sampleInfoArr) - 1
		for infoStart < infoEnd {
			sampleInfoArr[infoStart], sampleInfoArr[infoEnd] =
				sampleInfoArr[infoEnd], sampleInfoArr[infoStart]
			infoStart++
			infoEnd--
		}
	}
	
	//end:
	if err != nil {
		fmt.Printf("mongo find err=%v\n", err)
		R.RJson(c, "INTERNAL_ERROR")
		return
	}
	
	R.RData(c, sampleInfoArr)
}

/*
* @api PUT /api/v1/rolecard/update 按特征修改某个一个技能/武器/职业/角色卡（每个人最多只能修改到自己的）
* 当有在房间的角色卡修改时，会通过websocket通知房间所有玩家, websocket cmd为updateOneRoleCard，extend中有修改后的整个角色卡的信息
* @apiGroup rolecard
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery name string 必填，请求的表名称，比如角色卡：rolecard；技能：skill; 武器:weapon; 职业：job； 技能子选项: skselect 不同表决定了body数据中的字段
* @apiQuery filter string 选填，匹配的关键词, 即要修改哪个，比如要修改rolecard的chartid=coc7_1234的角色卡，则filter={"chartid":"coc7_1234"}
* @apiQuery rid string 选填，角色卡所在房间。一般情况下，修改用户只能修改自己拥有的角色卡，但当该角色卡在房间中，则该房间的房主也可以修改该角色卡；另外当该字段有填，则修改后会把角色卡信息推送到该房间的websocket上

* @apiRequest json
* @apiExample json
* {
	"order":"123",
	"step":"middle"
}

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
* 参数错误
* {
    "errcode":30002,
    "errmsg":"invalid parameter"
* }
* @apiExample json
* 数据库操作失败
* {
    "errcode":30009,
    "errmsg":"database operation error"
* }
*/
func PutUpdate(c *gin.Context) {
	//if Check(c) == false {
	//	R.RJson(c, "AUTH_FAILED")
	//	return
	//}
	
	rid := c.Query("rid")
	
	name := c.Query("name")
	if name == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	} else if name != "job" && name != "skill" && name != "rolecard" && name != "skselect" && name != "weapon" && name != "npccard" && name != "article" && name != "invitation" && name != "article_comment" && name != "invitation_comment" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	id := c.Query("userid")
	if id == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	userid, err := strconv.Atoi(id)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	filter := c.Query("filter")
	if filter == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	queryData := make(map[string]interface{})
	f := make(map[string]interface{})
	err = json.Unmarshal([]byte(filter), &f)
	if err == nil {
		for k, v := range f {
			if k == "_id" {
				id_key_str, ok := v.(string)
				if !ok {
					R.RJson(c, "INVALID_PARAM")
					return
				}
				queryData[k] = bson.ObjectIdHex(id_key_str)
			} else {
				queryData[k] = v
			}
		}
	}
	
	// 判断是否是房主并且要操作的角色卡在房间中
	if rid != "" && name == "rolecard" {
		chartid, ok := queryData["chartid"]
		fmt.Println(chartid)
		if ok && chartid != "" {
			rs := []Room{}
			DB.GetByFieldM(&rs, "id", rid, RoomTable)
			if len(rs) == 0 {
				R.RJson(c, "NOT_FOUND")
				return
			}
			
			if rs[0].Userid == int64(userid) {
				rr := []Roomrole{}
				sql_cmd := fmt.Sprintf("SELECT * FROM %s WHERE rid = '%s' AND JSON_CONTAINS(roleid, '[\"%s\"]');", RoomroleTable, rid, chartid)
				fmt.Println(sql_cmd)
				DB.Db.Raw(sql_cmd).Scan(&rr)
				
			}
		}
	}
	
	var sampleInfo map[string]interface{}
	err = c.BindJSON(&sampleInfo)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	_, ok := sampleInfo["_id"]
	if ok {
		delete(sampleInfo, "_id")
	}
	_, ok = sampleInfo["chartid"]
	if ok {
		delete(sampleInfo, "chartid")
	}

	if name =="article"{
		sampleInfo["pub_time"] = time.Now().Unix()
	}
	
	data := bson.M{"$set": sampleInfo}
	
	err = DB.MongoUpdate(Mgodb, name, queryData, data)
	if err != nil {
		R.RJson(c, "ERROR_DB")
		return
	}
	
	if rid != "" && name == "rolecard" {
		result, err := DB.MongoGetOneByData(Mgodb, "rolecard", queryData)
		if err != nil {
			R.RJson(c, "ERROR_DB")
			return
		}
		wsSendUpdateData(int64(userid), rid, "updateOneRoleCard", result)
	}
	
	R.RJson(c, "OK")
}

func wsSendUpdateData(userid int64, rid, cmd string, data interface{}) string {
	return ""
}

/**
* @api PUT /api/v1/rolecard/npcUpdate 修改npccard
* @apiGroup rolecard
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery name string 必填，填npccard

* @apiRequest json
* @apiParam old json
* @apiParam new json
* @apiExample json
* {
	"old":{
		"title":"aaa"
	},
	"new":{
		"title":"bbb"
	}
}
* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* {
    "errcode":0,
    "errmsg":"ok"
}
*/
func PutNpccard(c *gin.Context) {
	//鉴权
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	//参数检验
	id := c.Query("userid")
	if id == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	name := c.Query("name")
	if name == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	} else if name != "npccard" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	//组装更新查询条件
	var sampleInfo map[string]interface{}
	err := c.BindJSON(&sampleInfo)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	data := bson.M{"$set": sampleInfo["new"]}
	
	//更新MongoDB文档数据
	err = DB.MongoUpdate(Mgodb, name, sampleInfo["old"], data)
	if err != nil {
		R.RJson(c, "ERROR_DB")
		return
	}
	
	//返回结果
	R.RJson(c, "OK")
}

/**
* @api GET /api/v1/rolecard/GetNpcById 根据_id获取npccard
* @apiGroup rolecard
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery name string 必填，填npccard
* @apiQuery ids string 选填，_id数组

* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* {
    "data":""...
}
*/
func GetNpccardById(c *gin.Context) {
	/*if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}*/
	id := c.Query("npcId")
	name := c.Query("name")
	var sampleInfoArr []map[string]interface{}
	
	if id != "" {
		/*var ids_int []int64
		
		err := json.Unmarshal([]byte(ids), &ids_int)
		if err != nil {
			R.RJson(c, "INVALID_PARAM")
			return
		}*/
		var one map[string]interface{}
		temp := make(map[string]interface{})
		/*for _, id := range ids_int {*/
		temp["npcid"] = id
		err := Mgodb.C(name).Find(temp).One(&one)
		if err != nil {
			R.RJson(c, "ERROR_DB")
			return
		}
		if one != nil {
			sampleInfoArr = append(sampleInfoArr, one)
		}
		//}
		
	} else {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	if len(sampleInfoArr) == 0 {
		R.RData(c, sampleInfoArr)
		return
	}
	
	R.RData(c, sampleInfoArr)
}

/*
* @api POST /api/v1/rolecard/give 把角色卡赠与某人（赠与后会产生一个角色卡拷贝，但拷贝的角色卡id和原来不同）
* 要赠与的角色卡是本人的才可以赠与他人
* @apiGroup rolecard
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery chartid string 必填，标明赠送的角色卡chartid
* @apiQuery uid int 必填，标明赠送谁

* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiParam data json 角色卡拷贝的信息，包括id等
* @apiExample json
* {
    "errcode":0,
	"errmsg":"ok",
	"data":{
		"chartid":"coc7_1236"
	}
}

* @apiError 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* 参数错误
* {
    "errcode":30002,
    "errmsg":"invalid parameter"
* }
* @apiExample json
* 数据库操作失败
* {
    "errcode":30009,
    "errmsg":"database operation error"
* }
*/
func PostRolecardGive(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	
	uid := c.Query("userid")
	userid, _ := strconv.Atoi(uid)
	uidGive := c.Query("uid")
	useridGive, _ := strconv.Atoi(uidGive)
	
	chartid := c.Query("chartid")
	if chartid == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	chartid_list := strings.SplitN(chartid, "_", 2)
	if len(chartid_list) != 2 {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	query := make(map[string]interface{})
	query["chartid"] = chartid
	result, err := DB.MongoGetOneByData(Mgodb, "rolecard", query)
	if err != nil {
		R.RJson(c, "ERROR_DB")
		return
	}
	
	rr, ok := result.(bson.M)
	if !ok {
		R.RJson(c, "NOT_FOUND")
		return
	}
	
	userid_query, ok := rr["userid"].(float64)
	if !ok {
		R.RJson(c, "NOT_FOUND")
		return
	}
	if int(userid_query) != userid {
		R.RJson(c, "NO_PERMISSION")
		return
	}
	id := GetRolecardid(chartid_list[0])
	
	rr["chartid"] = id
	rr["userid"] = useridGive
	rr["_id"] = bson.NewObjectId()
	
	err = DB.MongoCreate(Mgodb, "rolecard", rr)
	if err != nil {
		R.RJson(c, "ERROR_DB")
		return
	}
	
	R.RData(c, id)
	return
}

/*
* @api GET /api/v1/rolecard/exist 判断某用户是否存在某角色卡
* @apiGroup rolecard
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery chartid string 必填，角色卡chartid
* @apiQuery uid int 必填，要判断的用户

* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiParam data bool true:uid的用户存在charid的角色卡，false为不存在
* @apiExample json
* {
    "errcode":0,
	"errmsg":"ok",
	"data":true
}

* @apiError 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* 参数错误
* {
    "errcode":30002,
    "errmsg":"invalid parameter"
* }
* @apiExample json
* 数据库操作失败
* {
    "errcode":30009,
    "errmsg":"database operation error"
* }
*/
func GetRolecardExist(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	
	uid := c.Query("userid")
	userid, _ := strconv.Atoi(uid)
	
	chartid := c.Query("chartid")
	if chartid == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	query := make(map[string]interface{})
	query["chartid"] = chartid
	result, err := DB.MongoGetOneByData(Mgodb, "rolecard", query)
	if err != nil {
		R.RJson(c, "ERROR_DB")
		return
	}
	
	rr, ok := result.(bson.M)
	if !ok {
		R.RData(c, false)
		return
	}
	
	userid_query, ok := rr["userid"].(float64)
	if !ok {
		R.RData(c, false)
		return
	}
	if int(userid_query) != userid {
		R.RData(c, false)
		return
	}
	
	R.RData(c, true)
	return
}

/*
* @api GET /api/v1/rolecard/all 获取job、skill、weapons、skselect全部信息，类似all.json
* @apiGroup rolecard
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery select []string 选填，选择要返回哪些信息，默认返回job、skill、weapons、npccard、skselect公用库信息，如果select=["job","skill"]则只返回job和skill信息
* @apiQuery needcommon bool 选填，是否要返回公有的，false为不返回公有的，默认为true
* @apiQuery needprivate bool 选填，是否要返回私有的，false为不返回私有的，默认为false

* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiParam data json job、skill、weapons、skselect信息合集（可以根据每一项的userid来判断是私有的还是公有的，userid不为0则为私有）
* @apiExample json
* {
    "errcode":0,
	"errmsg":"ok",
	"data":{
			"job":[...],
			"skill":[...],
			"weapons":[...],
			"skselect":[...]
		}
	}
}

* @apiError 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* 参数错误
* {
    "errcode":30002,
    "errmsg":"invalid parameter"
* }
* @apiExample json
* 数据库操作失败
* {
    "errcode":30009,
    "errmsg":"database operation error"
* }
*/
func GetRolecardAll(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	
	uid := c.Query("userid")
	userid, _ := strconv.Atoi(uid)
	
	sel := c.Query("select")
	if sel == "" {
		sel = "[\"job\",\"skill\",\"weapons\",\"skselect\"]"
	}
	
	needprivate := c.Query("needprivate")
	if needprivate != "true" {
		needprivate = "false"
	}
	needcommon := c.Query("needcommon")
	if needcommon != "false" {
		needcommon = "true"
	}
	
	common := make(map[string][]map[string]interface{})
	
	needSearchCharMap := make(map[string]interface{})
	
	chose := []string{}
	err := json.Unmarshal([]byte(sel), &chose)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	
	if needcommon == "true" {
		needSearchCharMap["userid"] = 0
		
		for _, c := range chose {
			var sampleInfoArr []map[string]interface{}
			Mgodb.C(c).Find(needSearchCharMap).Sort("_id").All(&sampleInfoArr)
			if sampleInfoArr != nil {
				common[c] = append(common[c], sampleInfoArr...)
			}
		}
	}
	
	if needprivate == "true" {
		needSearchCharMap["userid"] = userid
		
		for _, c := range chose {
			var sampleInfoArr []map[string]interface{}
			Mgodb.C(c).Find(needSearchCharMap).Sort("_id").All(&sampleInfoArr)
			if sampleInfoArr != nil {
				common[c] = append(common[c], sampleInfoArr...)
			}
		}
	}
	
	R.RData(c, common)
}

func LoadConfig() (*map[string]int, error) {
	file, err := os.Open(rolecardidfile)
	if err != nil {
		fmt.Println("<ERR> rolecardid open file: ", err.Error())
		return nil, err
	}
	
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	
	c := make(map[string]int)
	err = decoder.Decode(&c)
	if err != nil {
		fmt.Println("<ERR> CONFIG FIle decode: ", err.Error())
		
		return nil, err
	}
	return &c, nil
}

func SaveConfig(data map[string]int) error {
	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println("<ERR> CONFIG data marshal error: ", err.Error())
		return err
	}
	
	var out bytes.Buffer
	json.Indent(&out, d, "", "\t")
	
	err = ioutil.WriteFile(rolecardidfile, out.Bytes(), 0666)
	if err != nil {
		fmt.Println("<ERR> CONFIG FIle WriteFile: ", err.Error())
		return err
	}
	return nil
}

func RolecardInit() {
	m, err := LoadConfig()
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		rolecardidmap = make(map[string]int)
	} else {
		rolecardidmap = *m
	}
	
	index := mgo.Index{
		Key:        []string{"chartid"},
		Unique:     false,
		DropDups:   false,
		Background: true, // See notes.
		Sparse:     false,
	}
	err = Mgodb.C("rolecard").EnsureIndex(index)
	fmt.Println("set index:", err)
}

func NpccardInit() {
	m, err := LoadConfig()
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		npccardidmap = make(map[string]int)
	} else {
		npccardidmap = *m
	}
	
	index := mgo.Index{
		Key:        []string{"npcid"},
		Unique:     false,
		DropDups:   false,
		Background: true, // See notes.
		Sparse:     false,
	}
	err = Mgodb.C("npccard").EnsureIndex(index)
	fmt.Println("set index:", err)
}
