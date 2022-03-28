package main

import (
	_ "encoding/json"
	_ "fmt"
	"miao/Utils/DB"
	_ "miao/Utils/DB"
	"miao/Utils/R"
	_ "miao/Utils/R"
	_ "net/http"
	_ "regexp"
	_ "strconv"
	_ "sync"
	"time"
	_ "time"
	
	"github.com/gin-gonic/gin"
	_ "github.com/globalsign/mgo/bson"
)

var (
	Ban_usertable string
)

type BanReq struct {
	UId  int64 `json:"uid"`
	Time int64 `json:"times"`
}
type Ban_user struct {
	UId       int64  `json:"uid" gorm:"index"`
	Suspend   bool   `json:"suspend"`
	Starttime string `json:"starttime"`
	Endtime   string `json:"endtime"`
}

/**
* @api POST /api/v1/user_ban/ban_user 新增封号用户
* @apiGroup user_ban
* @apiQuery userid int 必填，执行请求的用户id（只有管理员用户才能封禁,管理员id在config.json中admin处设置）
* @apiQuery token string 必填，执行请求的用户token
* @apiRequest multipart/form-data

* @apiQuery times int 必填，要封禁的时间，以天为单位
* @apiQuery uid string 必填，要封禁用户的id

* @apiExample json
*{
    "data": {
        "uid": 2,
        "suspend": true,
        "starttime": "2021-04-26 12:31:14",
        "endtime": "2021-05-11 12:31:14"
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

func PostBanUser(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	

	banreq := BanReq{}
	err := c.BindJSON(&banreq)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	uid := banreq.UId
	times := banreq.Time
	//uid, err :=strconv.Atoi(c.Request.FormValue("uid"))
	//if err != nil {
	//	R.RJson(c, "INVALID_PARAM")
	//	return
	//}
	//
	//times, err := strconv.Atoi(c.Request.FormValue("times"))
	//if err != nil {
	//	R.RJson(c, "INVALID_PARAM")
	//	return
	//}

	bts := []Ban_user{}
	DB.GetByFieldM(&bts, "uid", uid, Ban_usertable)
	if len(bts) != 0 {
		DB.Db.Table(Ban_usertable).Where("`uid`=?", uid).Delete(nil)
	}

	ban_user := Ban_user{}
	ts := make([]Token,0)

	DB.GetByFieldM(&ts, "uid", int64(uid), TokenTable)
	for _,t:=range ts{
		t.Uid = int64(uid)
		t.Token = "1"
		t.Expired = int(time.Now().Unix()) + (3600 * 24 * TokenExpire)
		if t.Id != 0 {
			DB.SaveM(&t, TokenTable)
		} else {
			DB.CreateM(&t, TokenTable)
		}
	}
	

	ban_user.UId = int64(uid)
	//ban_user.Email= email
	ban_user.Suspend = true
	nowtime := time.Now()
	ban_user.Starttime = time.Unix(nowtime.Unix(), 0).Format("2006-01-02 15:04:05")
	ban_user.Endtime = time.Unix(nowtime.AddDate(0, 0, int(times)).Unix(), 0).Format("2006-01-02 15:04:05")

	err = DB.CreateM(&ban_user, Ban_usertable)
	if err != nil {
		DB.Db.Table(Ban_usertable).Where("`uid`=?", ban_user.UId).Delete(nil)
		R.RJson(c, "ERROR_DB")
		return
	}
	R.RData(c, ban_user)
}

func DEBan() { //检测是否有账号解封
	bts := []Ban_user{}
	DB.GetByFieldM(&bts, "suspend", true, Ban_usertable)
	if len(bts) == 0 {
		//R.RData(c, "OKe")
		return
	}
	var i int
	for i = 0; i < len(bts); i++ {
		nowtime := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
		//endtime, _ :=time.ParseInLocation("2006-01-02 15:04:05", bt.Endtime, time.Local)
		t1, err1 := time.Parse("2006-01-02 15:04:05", nowtime)
		t2, err2 := time.Parse("2006-01-02 15:04:05", bts[i].Endtime)
		if err1 == nil && err2 == nil && t1.After(t2) {
			bts[i].Suspend = false
			//DB.UpdateByFieldM(&bts[i],"uid", bts[i].UId, Ban_usertable)
			err := DB.DeleteByFieldM("uid", bts[i].UId, Ban_usertable)
			if err != nil {
				return
			}
			err = DB.SaveM(&bts[i], Ban_usertable)
			if err != nil {
				return
			}
		}
	}
}

func GetBan(c *gin.Context) { //暂时弃用
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	email := c.Query("email")
	if email == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}

	bts := []Ban_user{}
	DB.GetByFieldM(&bts, "email", email, Ban_usertable)
	if len(bts) == 0 {
		R.RJson(c, "NOT_FOUND")
		return
	}

	bt := bts[0]
	//直接返回封号信息，检测是否已到封号结束在另一个接口
	//R.RData(c, bt)

	//获取与检测封号信息都在本接口完成
	nowtime := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	//endtime, _ :=time.ParseInLocation("2006-01-02 15:04:05", bt.Endtime, time.Local)
	t1, err1 := time.Parse("2006-01-02 15:04:05", nowtime)
	t2, err2 := time.Parse("2006-01-02 15:04:05", bt.Endtime)

	if err1 == nil && err2 == nil && t1.After(t2) {
		bt.Suspend = false
		err := DB.DeleteByFieldM("email", email, Ban_usertable)
		if err != nil {
			R.RJson(c, "ERROR_DB")
			return
		}
		err = DB.SaveM(&bt, Ban_usertable)
		if err != nil {
			R.RJson(c, "ERROR_DB")
			return
		}
	}
	R.RData(c, bt)

}

func BanInit() {
	DB.CreateTableM(Ban_user{})
	Ban_usertable = DB.GetTableNameM(Ban_user{})
	r := Ban_user{}
	DB.GetLastM(&r, Ban_usertable)
}
