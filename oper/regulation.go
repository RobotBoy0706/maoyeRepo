package main

import (
	"fmt"
	"miao/Utils/DB"
	"miao/Utils/ErrCode"
	"miao/Utils/R"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Regulation struct {
	Id            int64  `json:"id" gorm:"index"`
	Name          string `json:"name" gorm:"index"`
	Userid        int64  `json:"userid" gorm:"index"` //规则由谁创建
	Tenstrikemax  int    `json:"tenstrikemax"`        //大成功最大值 < 50
	Tenstrikemin  int    `json:"tenstrikemin"`        //大成功最小值 < 50
	Fiascomax     int    `json:"fiascomax"`           //大失败最大值 < 50
	Fiascomin     int    `json:"fiascomin"`           //大失败最小值 < 50
	Tenstrikemax2 int    `json:"tenstrikemax2"`       //大成功最大值 >50
	Tenstrikemin2 int    `json:"tenstrikemin2"`       //大成功最小值 >50
	Fiascomax2    int    `json:"fiascomax2"`          //大失败最大值 >50
	Fiascomin2    int    `json:"fiascomin2"`          //大失败最小值 >50

	Acceptold bool   `json:"acceptold,omitempty"` //是否接受老卡
	Car       string `json:"car,omitempty"`       //车卡
	Gugu      int    `json:"gugu,omitempty"`      //咕咕值
	Friendly  int    `json:"friendly,omitempty"`  //友好值
	Friend    bool   `json:"friend,omitempty"`    //是否需要是好友

	Createtime int64 `json:"createtime"`
	Updatetime int64 `json:"updatetime"`
}

var (
	RegulationTable string
)

func CreateOrPutRegulation(re *Regulation) (int, string) {
	if re == nil || re.Userid == 0 {
		return ErrCode.GetReturnCode("INVALID_PARAM")
	}

	if re.Name == "" {
		re.Name = "default"
	}
	//
	//if re.Tenstrikemax == 0 {//实际使用的大成功
	//    re.Tenstrikemax = 1
	//}
	//if re.Tenstrikemin == 0 {
	//    re.Tenstrikemin = 1
	//}
	//if re.Fiascomax == 0 {//实际使用的大失败
	//    re.Fiascomax = 100
	//}
	//if re.Fiascomin == 0 {
	//    re.Fiascomin = 95
	//}
	//if re.Tenstrikemax2 == 0 {
	//    re.Tenstrikemax2 = 1
	//}
	//if re.Tenstrikemin2 == 0 {
	//    re.Tenstrikemin2 = 1
	//}
	//if re.Fiascomax2 == 0 {
	//    re.Fiascomax2 = 100
	//}
	//if re.Fiascomin2 == 0 {
	//    re.Fiascomin2 = 100
	//}

	res := []Regulation{}
	DB.Db.Table(RegulationTable).Where("userid = ? AND name = ?", re.Userid, re.Name).Find(&res)
	if len(res) != 0 {
		re.Id = res[0].Id
		re.Createtime = res[0].Createtime
		re.Updatetime = time.Now().Unix()

		DB.UpdateByFieldM(&re, "id", re.Id, RegulationTable)
	} else {
		re.Createtime = time.Now().Unix()
		re.Updatetime = time.Now().Unix()

		err := DB.CreateM(&re, RegulationTable)
		if err != nil {
			return ErrCode.GetReturnCode("ERROR_DB")
		}
	}

	return ErrCode.GetReturnCode("OK")
}

/*
* @api POST /api/v1/regulation/regulation 新增一个规则(如果操作者名下已有名字相同的规则则覆盖之)
* @apiGroup regulation
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token

* @apiRequest json
* @apiParam name string 选填，规则名称
* @apiParam tenstrikemax	int 必填，大成功最大值(技能值<50)
* @apiParam tenstrikemin	int 必填，大成功最小值(技能值<50)
* @apiParam fiascomax int 必填，大失败最大值(技能值<50)
* @apiParam fiascomin int 必填，大失败最小值(技能值<50)
* @apiParam tenstrikemax2	int 必填，大成功最大值(技能值>50)
* @apiParam tenstrikemin2	int 必填，大成功最小值(技能值>50)
* @apiParam fiascomax2 int 必填，大失败最大值(技能值>50)
* @apiParam fiascomin2 int 必填，大失败最小值(技能值>50)
* @apiParam acceptold int 必填，是否接受老卡
* @apiParam car string 必填，车卡
* @apiParam gugu int 必填，咕咕值
* @apiParam friendly int 必填，友好值
* @apiParam friend bool 必填，是否需要是好友
* @apiExample json
* {
	"name":"r1",
	"tenstrikemax":5,   //大成功最大值(技能值<50)
    "tenstrikemin":0,    //大成功最小值(技能值<50)
    "fiascomax":90,      //大失败最大值(技能值<50)
    "fiascomin":60,      //大失败最小值(技能值<50)
    "tenstrikemax2":5,   //大成功最大值(技能值>50)
    "tenstrikemin2":0,    //大成功最小值(技能值>50)
    "fiascomax2":90,      //大失败最大值(技能值>50)
    "fiascomin2":60,      //大失败最小值(技能值>50)
	"acceptold":false,   //是否接受老卡
	"car":"...",         //车卡
	"gugu":34,        //咕咕值
	"friendly": 54,    //友好值
	"friend":true   //是否需要是好友
}

* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiParam data json 新建的规则的信息，除了上面所填，还包括后台生成的规则编号
* @apiExample json
* {
    "errcode":0,
    "errmsg":"ok",
    "data":{
        "id":1,
        "name":"r1",
        "userid":2,   //规则创建者
        "tenstrikemax":5,   //大成功最大值(技能值<50)
        "tenstrikemin":0,    //大成功最小值(技能值<50)
        "fiascomax":90,      //大失败最大值(技能值<50)
        "fiascomin":60,      //大失败最小值(技能值<50)
        "tenstrikemax2":5,   //大成功最大值(技能值>50)
        "tenstrikemin2":0,    //大成功最小值(技能值>50)
        "fiascomax2":90,      //大失败最大值(技能值>50)
        "fiascomin2":60,      //大失败最小值(技能值>50)
        "acceptold":false,   //是否接受老卡
        "car":"...",         //车卡
        "gugu":34,        //咕咕值
        "friendly": 54,    //友好值
        "friend":true   //是否需要是好友
    }
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
func PostRegulation(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}

	userid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}

	re := Regulation{}
	err = c.BindJSON(&re)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}

	re.Userid = int64(userid)
	ec, em := CreateOrPutRegulation(&re)
	if ec != 0 {
		R.Jr(c, http.StatusOK, gin.H{"errcode": ec, "errmsg": em})
		return
	}

	R.RData(c, re)
}

/*
* @api GET /api/v1/regulation/regulation 根据规则id或者名字获取规则(需要规则属于查询者的)
* @apiGroup regulation
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery id int 选填，规则ID
* @apiQuery name 必填，规则名称， id和name必填一个，两个都填则返回id和name都匹配的值

* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiParam data json 规则的信息
* @apiExample json
* {
    "errcode":0,
    "errmsg":"ok",
    "data":{
            "id":1,
            "name":"r1",
            "userid":2,   //规则创建者
            "tenstrikemax":5,   //大成功最大值(技能值<50)
            "tenstrikemin":0,    //大成功最小值(技能值<50)
            "fiascomax":90,      //大失败最大值(技能值<50)
            "fiascomin":60,      //大失败最小值(技能值<50)
            "tenstrikemax2":5,   //大成功最大值(技能值>50)
            "tenstrikemin2":0,    //大成功最小值(技能值>50)
            "fiascomax2":90,      //大失败最大值(技能值>50)
            "fiascomin2":60,      //大失败最小值(技能值>50)
            "acceptold":false,   //是否接受老卡
            "car":"...",         //车卡
            "gugu":34,        //咕咕值
            "friendly": 54,    //友好值
            "friend":true   //是否需要是好友
        }
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
func GetRegulation(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}

	userid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}

	id := c.Query("id")
	name := c.Query("name")

	res := []Regulation{}

	if id != "" && name != "" {
		i, err := strconv.Atoi(id)
		if err != nil {
			R.RJson(c, "INVALID_PARAM")
			return
		}

		DB.Db.Table(RegulationTable).Where("id = ? AND name = ? AND userid = ?", i, name, userid).Find(&res)
	} else if id != "" {
		i, err := strconv.Atoi(id)
		if err != nil {
			R.RJson(c, "INVALID_PARAM")
			return
		}

		DB.Db.Table(RegulationTable).Where("id = ? AND userid = ?", i, userid).Find(&res)
	} else if name != "" {
		DB.Db.Table(RegulationTable).Where("name = ? AND userid = ?", name, userid).Find(&res)
	} else {
		R.RJson(c, "INVALID_PARAM")
		return
	}

	if len(res) == 0 {
		R.RJson(c, "NOT_FOUND")
		return
	}

	R.RData(c, res[0])

}

/*
* @api GET /api/v1/regulation/regulations 查询查询者创建的所有规则
* @apiGroup regulation
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery page string 选填，获取多少页的信息,必须和limit一起填写
* @apiQuery limit string 选填，获取多少个信息,必须和page一起填写

* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiParam data json 规则的信息
* @apiExample json
* {
    "errcode":0,
    "errmsg":"ok",
    "data":[
            {
            "id":1,
            "name":"r1",
            "userid":2,   //规则创建者
            "tenstrikemax":5,   //大成功最大值(技能值<50)
            "tenstrikemin":0,    //大成功最小值(技能值<50)
            "fiascomax":90,      //大失败最大值(技能值<50)
            "fiascomin":60,      //大失败最小值(技能值<50)
            "tenstrikemax2":5,   //大成功最大值(技能值>50)
            "tenstrikemin2":0,    //大成功最小值(技能值>50)
            "fiascomax2":90,      //大失败最大值(技能值>50)
            "fiascomin2":60,      //大失败最小值(技能值>50)
            "acceptold":false,   //是否接受老卡
            "car":"...",         //车卡
            "gugu":34,        //咕咕值
            "friendly": 54,    //友好值
            "friend":true   //是否需要是好友
        }
    ]
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
func GetRegulations(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}

	userid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}

	page := c.Query("page")
	limit := c.Query("limit")

	rs := []Regulation{}

	if page != "" && limit != "" {
		p, err := strconv.Atoi(c.Query("page"))
		l, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			fmt.Println(err)
			R.RJson(c, "INVALID_PARAM")
			return
		}

		offset := (p - 1) * l

		DB.Db.Table(RegulationTable).Where("userid = ?", userid).Offset(offset).Limit(l).Find(&rs)
	} else {
		DB.GetByFieldM(&rs, "userid", userid, RegulationTable)
	}

	R.RData(c, rs)
}

func RegulationInit() {
	DB.CreateTableM(Regulation{})
	RegulationTable = DB.GetTableNameM(Regulation{})
	DB.Db.Table(RegulationTable).AutoMigrate(&Regulation{})
}
