package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"miao/Utils/DB"
	"miao/Utils/R"
	"strconv"
	"time"
)

type Watch struct {
	Id         int    `json:"id" gorm:"index"`
	Userid     int64  `json:"userid"`
	Date       string `json:"date"`
	Type       int    `json:"type"` //1交流区 2文章
	CollectID  string `json:"collect_id"`
	Createtime int64  `json:"createtime"`
	Updatetime int64  `json:"updatetime" gorm:"index"`
}

type Like struct {
	Id         int    `json:"id" gorm:"index"`
	Userid     int64  `json:"userid"`
	Type       int    `json:"type"` //1交流区 2文章 3交流区评论点赞  4 文章评论点赞
	CollectID  string `json:"collect_id"`
	Createtime int64  `json:"createtime"`
	Updatetime int64  `json:"updatetime" gorm:"index"`
}

var (
	WatchTable string
	LikeTable  string
)

func ArticleInit() {
	DB.CreateTableM(Watch{})
	WatchTable = DB.GetTableNameM(Watch{})
	DB.Db.Table(WatchTable).AutoMigrate(&Watch{})

	DB.CreateTableM(Like{})
	LikeTable = DB.GetTableNameM(Like{})
	DB.Db.Table(LikeTable).AutoMigrate(&Like{})
}

/**
* @api GET /api/v1/collection/detail 获取mongo collection详细
* @apiGroup 文章帖子
* @apiQuery userid  int	必填，用户ID
* @apiQuery token   string	必填，登录token
* @apiQuery name string 必填，请求的表名称，比如角色卡：rolecard；技能：skill; 武器:weapon; 职业：job； 技能子选项: skselect 不同表决定了body数据中的字段
* @apiQuery id   string  取上一次查询结果数组中第一个元素的“_id”作为参数
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
*/
func AricleDetail(c *gin.Context) {
	//if Check(c) == false {
	//	R.RJson(c, "AUTH_FAILED")
	//	return
	//}

	name := c.Query("name")
	if name == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	} else if name != "job" && name != "skill" && name != "rolecard" && name != "skselect" && name != "weapon" && name != "npccard" && name != "article" && name != "invitation" && name != "article_comment" && name != "invitation_comment"&& name != "report"{
		R.RJson(c, "INVALID_PARAM")
		return
	}
	id := c.Query("id")
	if id == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}

	needSearchCharMap := make(map[string]interface{})
	db := Mgodb.C(name)
	needSearchCharMap["_id"] = bson.ObjectIdHex(id)
	fmt.Printf("needSearchCharMap=%#v\n", needSearchCharMap)
	var sampleInfoArr []map[string]interface{}
	err := db.Find(needSearchCharMap).All(&sampleInfoArr)

	//end:
	if err != nil {
		fmt.Printf("mongo find err=%v\n", err)
		R.RJson(c, "INTERNAL_ERROR")
		return
	}

	if name =="report" &&len(sampleInfoArr)>0{
		where := make(map[string]interface{})

		bussName,_:=sampleInfoArr[0]["buss_name"].(string)
		where["_id"]=sampleInfoArr[0]["buss_id"]
		if bussName==""{
			fmt.Printf("report bussName is nil \n")
			R.RJson(c, "INTERNAL_ERROR")
			return
		}
		var infos []map[string]interface{}
		err := Mgodb.C(bussName).Find(where).All(&infos)

		//end:
		if err != nil {
			fmt.Printf("mongo find err=%v\n", err)
			R.RJson(c, "INTERNAL_ERROR")
			return
		}
		if len(infos)>0{
			sampleInfoArr[0]["info"] = infos[0]
		}
	}


	R.RData(c, sampleInfoArr)
}


/**
* @api GET /api/v1/report/update 举报审核
* @apiGroup 文章帖子
* @apiQuery userid  int	必填，用户ID
* @apiQuery token   string	必填，登录token
* @apiQuery name string 必填，请求的表名称 例“report”
* @apiQuery id   string  举报业务id
* @apiQuery stat string 必填 2通过 3不通过 4删除

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
*/
func ReportUpdate(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}

	name := c.Query("name")
	if name == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	} else if name != "job" && name != "skill" && name != "rolecard" && name != "skselect" && name != "weapon" && name != "npccard" && name != "article" && name != "invitation" && name != "article_comment" && name != "invitation_comment"&& name != "report"{
		R.RJson(c, "INVALID_PARAM")
		return
	}
	id := c.Query("id")
	if id == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}

	statStr:=c.Query("stat")
	stat,_:=strconv.Atoi(statStr)
	if stat == 0{
		R.RJson(c, "INVALID_PARAM")
		return
	}

	needSearchCharMap := make(map[string]interface{})
	db := Mgodb.C(name)
	needSearchCharMap["_id"] = bson.ObjectIdHex(id)
	fmt.Printf("needSearchCharMap=%#v\n", needSearchCharMap)
	var sampleInfoArr []map[string]interface{}
	err := db.Find(needSearchCharMap).All(&sampleInfoArr)

	//end:
	if err != nil {
		fmt.Printf("mongo find err=%v\n", err)
		R.RJson(c, "INTERNAL_ERROR")
		return
	}

	if name =="report" &&len(sampleInfoArr)>0{
		where := make(map[string]interface{})

		bussName,_:=sampleInfoArr[0]["buss_name"].(string)
		where["_id"]=sampleInfoArr[0]["buss_id"]
		if bussName==""{
			fmt.Printf("report bussName is nil \n")
			R.RJson(c, "INTERNAL_ERROR")
			return
		}
		var infos []map[string]interface{}
		err := Mgodb.C(bussName).Find(where).All(&infos)

		//end:
		if err != nil {
			fmt.Printf("mongo find err=%v\n", err)
			R.RJson(c, "INTERNAL_ERROR")
			return
		}
		if len(infos)==0{
			fmt.Printf("mongo find is nil")
			R.RJson(c, "INTERNAL_ERROR")
			return
		}
		originBussID,_:=infos[0]["buss_id"].(string)
		var resportStat int //1举报不通过 2 举报通过
		switch stat {
		case 2:
			resportStat=2
			update:=map[string]interface{}{
				"$set": map[string]interface{}{
					"report_stat":     resportStat,
					"pub_time": time.Now().Unix(),
					"update_time": time.Now().Unix(),
				}}
			err = DB.MongoUpdate(Mgodb, bussName, where, update)
			if err != nil {
				R.RJson(c, "ERROR_DB")
				return
			}


		case 3:
			resportStat=1
			update:=map[string]interface{}{
				"$set": map[string]interface{}{
					"$set": map[string]interface{}{
						"report_stat":     resportStat,
						"stat":     1,

				}}}

			err = DB.MongoUpdate(Mgodb, bussName, where, update)
			if err != nil {
				R.RJson(c, "ERROR_DB")
				return
			}
			switch bussName {
			case "invitaion_comment":
				err = DB.MongoUpdate(Mgodb, "invitation", map[string]interface{}{"invid":originBussID}, map[string]interface{}{"$inc": map[string]interface{}{
					"comment_count": -1,
				}})
				if err != nil {
					R.RJson(c, "ERROR_DB")
					return
				}
			case "article_comment":
				err = DB.MongoUpdate(Mgodb, "article", map[string]interface{}{"artid":originBussID}, map[string]interface{}{"$inc": map[string]interface{}{
					"comment_count": -1,
				}})
				if err != nil {
					R.RJson(c, "ERROR_DB")
					return
				}
			}
		case 4:
			err = DB.MongoUpdate(Mgodb, bussName, where, bson.M{"$set": map[string]interface{}{
				"stat":     4,

			}})
			if err != nil {
				R.RJson(c, "ERROR_DB")
				return
			}
			switch bussName {
			case "invitaion_comment":
				err = DB.MongoUpdate(Mgodb, "invitation", map[string]interface{}{"invid":originBussID}, map[string]interface{}{"$incr": map[string]interface{}{
					"comment_count": -1,
				}})
				if err != nil {
					R.RJson(c, "ERROR_DB")
					return
				}
			case "article_comment":
				err = DB.MongoUpdate(Mgodb, "article", map[string]interface{}{"artid":originBussID}, map[string]interface{}{"$incr": map[string]interface{}{
					"comment_count": -1,
				}})
				if err != nil {
					R.RJson(c, "ERROR_DB")
					return
				}
			}



		}
		//举报改成已处理
		err = DB.MongoUpdate(Mgodb, name, needSearchCharMap, bson.M{"$set": map[string]interface{}{
			"stat":     2,
			"report_stat":     resportStat,
			"update_time":time.Now().Unix(),
		}})
	}


	R.RData(c, sampleInfoArr)
}



/**
* @api GET /api/v1/manager_report/update  管理员举报审核
* @apiGroup 文章帖子
* @apiQuery userid  int	必填，用户ID
* @apiQuery token   string	必填，登录token
* @apiQuery name string 必填，请求的表名称 例“report”
* @apiQuery id   string  举报业务id
* @apiQuery stat string 必填 2通过 3不通过 4删除

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
*/
func ManagerReportUpdate(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}

	name:="manager_report"
	id := c.Query("id")
	if id == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}

	statStr:=c.Query("stat")
	stat,_:=strconv.Atoi(statStr)
	if stat == 0{
		R.RJson(c, "INVALID_PARAM")
		return
	}

	needSearchCharMap := make(map[string]interface{})
	db := Mgodb.C(name)
	needSearchCharMap["_id"] = bson.ObjectIdHex(id)
	fmt.Printf("needSearchCharMap=%#v\n", needSearchCharMap)
	var sampleInfoArr []map[string]interface{}
	err := db.Find(needSearchCharMap).All(&sampleInfoArr)

	//end:
	if err != nil {
		fmt.Printf("mongo find err=%v\n", err)
		R.RJson(c, "INTERNAL_ERROR")
		return
	}

	if len(sampleInfoArr)>0{
		where := make(map[string]interface{})

		bussName,_:=sampleInfoArr[0]["buss_name"].(string)
		where["_id"]=sampleInfoArr[0]["buss_id"]
		if bussName==""{
			fmt.Printf("report bussName is nil \n")
			R.RJson(c, "INTERNAL_ERROR")
			return
		}
		var infos []map[string]interface{}
		err := Mgodb.C(bussName).Find(where).All(&infos)

		//end:
		if err != nil {
			fmt.Printf("mongo find err=%v\n", err)
			R.RJson(c, "INTERNAL_ERROR")
			return
		}
		if len(infos)==0{
			fmt.Printf("mongo find is nil")
			R.RJson(c, "INTERNAL_ERROR")
			return
		}
		bussStat,_:=infos[0]["stat"].(int)
		originBussID, _ := sampleInfoArr[0]["buss_id"].(string)
		var resportStat int //1举报不通过 2 举报通过
		switch stat {//stat= 2 审核通过 3 审核不通过
		case 2:
			resportStat=2
			update:=map[string]interface{}{
				"$set": map[string]interface{}{
					"stat":     2,
					"pub_time": time.Now().Unix(),
					"update_time": time.Now().Unix(),
				},
			}

			err = DB.MongoUpdate(Mgodb, bussName, where, update)
			if err != nil {
				R.RJson(c, "ERROR_DB")
				return
			}
			if bussStat != 2 {
				switch bussName {
				case "invitaion_comment":
					err = DB.MongoUpdate(Mgodb, "invitation", map[string]interface{}{"invid":originBussID}, map[string]interface{}{"$inc": map[string]interface{}{
						"comment_count": 1,
					}})
					if err != nil {
						R.RJson(c, "ERROR_DB")
						return
					}
				case "article_comment":
					err = DB.MongoUpdate(Mgodb, "article", map[string]interface{}{"artid":originBussID}, map[string]interface{}{"$inc": map[string]interface{}{
						"comment_count": 1,
					}})
					if err != nil {
						R.RJson(c, "ERROR_DB")
						return
					}
				}
			}
		case 3:
			update:=map[string]interface{}{
				"$set": map[string]interface{}{
					"report_stat":     resportStat,
				},
			}
			resportStat=1
			err = DB.MongoUpdate(Mgodb, bussName, where, update)
			if err != nil {
				R.RJson(c, "ERROR_DB")
				return
			}
		case 4:
			update:=map[string]interface{}{
				"$set": map[string]interface{}{
					"stat":     4,

				},
			}

			err = DB.MongoUpdate(Mgodb, bussName, where, update)
			if err != nil {
				R.RJson(c, "ERROR_DB")
				return
			}



		}
		//举报改成已处理
		err = DB.MongoUpdate(Mgodb, name, needSearchCharMap, bson.M{"$set": map[string]interface{}{
			"stat":     2,
			"report_stat":     resportStat,
			"update_time":time.Now().Unix(),
		}})
	}


	R.RData(c, sampleInfoArr)
}

