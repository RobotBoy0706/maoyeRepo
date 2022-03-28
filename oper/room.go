package main

import (
	"encoding/json"
	"fmt"
	"miao/Utils/DB"
	"miao/Utils/R"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"
	
	"github.com/gin-gonic/gin"
)

type RoomReq struct {
	Room
	Regulation Regulation             `json:"regulation"`
	Roomrole   []Roomrole             `json:"roomrole"`
	Roleinfo   map[string]interface{} `json:"roleinfo"`
}

type Room struct {
	Id                string `json:"id" gorm:"index"`            // 房间号
	Name              string `json:"name" gorm:"index"`          //游戏名称
	Password          string `json:"password"`                   // 密码
	Userid            int64  `json:"userid"`                     // 房主
	Kpname            string `json:"kpname"`                     // 房主名称
	Era               string `json:"era"`                        // 时代
	Language          string `json:"language"`                   // 母语
	Site              string `json:"site"`                       // 地点
	Max               int    `json:"max" gorm:"index"`           // 最大人数
	Min               int    `json:"min" gorm:"index"`           //最小人数
	Background        string `json:"background"`                 //背景
	Description       string `json:"description"`                //特别说明
	Recommendskills   string `json:"recommendskills"`            //推荐技能
	Unrecommendskills string `json:"unrecommendskills"`          //不推荐技能
	Limitedskills     string `json:"limitedskills"`              //技能限定
	Ruleid            int64  `json:"ruleid"`                     //规则ID
	Regulationid      int64  `json:"regulationid"`               //房规ID
	Flag              bool   `json:"flag"`                       //是否满员自动开始
	Status            int    `json:"status"`                     //房间状态     0:刚创建, 1:正在进行, 2:暂停
	Sidekps           Int64s `json:"Sidekps" sql:"type:json"`    // 副房主userid
	Players           Int64s `json:"players" sql:"type:json"`    // 玩家userid
	Spectators        Int64s `json:"spectators" sql:"type:json"` // 旁观者 userid
	Onlineusers       int    `json:"onlineusers"`                // 在线人数，旁观者+在线的玩家数
	Acceptold         bool   `json:"acceptold"`                  //是否接受老卡
	Car               string `json:"car"`                        //车卡
	Gugu              int    `json:"gugu"`                       //咕咕值
	Friendly          int    `json:"friendly"`                   //友好值
	Friend            bool   `json:"friend"`                     //是否需要是好友
	Checkother        bool   `json:"checkother"`                 //能否查看其它玩家角色卡

	Tags      Strings `json:"tags" sql:"type:json"`      //[]string
	Gametimes Strings `json:"gametimes" sql:"type:json"` //[]string
	Moduleid  int64   `json:"moduleid"`                  //模组的资源ID
	Icon      string  `json:"icon"`                      // 房间图表的md5值
	Chatid    int64   `json:"chatid"`                    //房间聊天室id

	Createtime int64 `json:"createtime"`
	Updatetime int64 `json:"updatetime" gorm:"index"`
}

type Roomrole struct {
	Id     int64   `json:"id"`
	Rid    string  `json:"rid" gorm:"index"`
	Userid int64   `json:"userid"`
	Roleid Strings `json:"roleid" sql:"type:json"`
}

type Npccards struct {
	Id      int64   `json:"id"`
	Rid     string  `json:"rid" gorm:"index"`
	Npccard Strings `json:"Npccard" sql:"type:json"` //[]string
}

var (
	RoomTable     string
	RoomroleTable string
	RoomMapTable string

	RoomIdNum   int
	roomidmutex sync.Mutex

	NpccardsTable string

	NpcIdNum int
	npcMutex sync.Mutex

	flag bool
)

func createRoomId() (id string) {
	roomidmutex.Lock()
	RoomIdNum++

	id = fmt.Sprintf("F%08d", RoomIdNum)
	roomidmutex.Unlock()

	return
}

func createNpcId() (id string) {
	npcMutex.Lock()
	NpcIdNum++

	id = fmt.Sprintf("F%08d", NpcIdNum)
	npcMutex.Unlock()

	return
}

func checkRoomId(id string) bool {
	ismatch, err := regexp.MatchString("F[0-9]{8}", id)
	if err != nil {
		return false
	}

	return ismatch
}

func CalcDayBorderTime(ts int64) (min, max int64) {
	fmt.Println(ts)
	t := time.Unix(ts, 0)

	t_min := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	t_max := time.Date(t.Year(), t.Month(), t.Day(), 24, 0, 0, 0, t.Location())

	min = t_min.Unix()
	max = t_max.Unix()

	fmt.Println(min, "-", max)

	return
}

func UpdateTimeQuery(ts string) string {
	ts_string := []string{}
	err := json.Unmarshal([]byte(ts), &ts_string)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	append_string := []string{}
	for _, t_string := range ts_string {
		var (
			t_min int64
			t_max int64
		)
		fmt.Sscanf(t_string, "%d-%d", &t_min, &t_max)

		min, max := CalcDayBorderTime(t_min)

		if (t_min != min) || (t_max != max) {
			total := fmt.Sprintf("%d-%d", min, max)
			append_string = append(append_string, total)
		}
	}

	ts_string = append(ts_string, append_string...)
	m, err := json.Marshal(ts_string)
	if err != nil {
		return ""
	}
	fmt.Println(string(m))
	return string(m)

}

/*
* @api POST /api/v1/room/room 新增一个房间
* @apiGroup room
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token

* @apiRequest json
* @apiParam name string 必填，房间（游戏）名称
* @apiParam kpname string 必填，房主（KP）名称
* @apiParam password string 选填，房间密码
* @apiParam era	string 选填，时代
* @apiParam language	string 选填，母语
* @apiParam site string 选填，地点
* @apiParam max int 必填，最多人数
* @apiParam min int 必填，最少人数
* @apiParam background string 必填，背景
* @apiParam description string 选填，特别说明
* @apiParam recommendskills string 选填，推荐技能
* @apiParam unrecommendskills string 选填，不推荐技能
* @apiParam limitedskills string 选填，技能限定
* @apiParam flag bool 选填，是否满员自动开始，不填默认为false
* @apiParam ruleid int 必填，规则id
* @apiParam gametimes []string 必填，游戏时间, 值统一为时间戳，例如["1558080000-1558083600", "1558170000-1558173600"]
* @apiParam tags []string 选填，标签
* @apiParam moduleid int 选填，模组的资源ID
* @apiParam acceptold int 必填，是否接受老卡
* @apiParam car string 必填，车卡
* @apiParam gugu int 必填，咕咕值
* @apiParam friendly int 必填，友好值
* @apiParam friend bool 必填，是否需要是好友
* @apiParam checkother bool 必填，能否查看其它玩家角色卡
* @apiParam regulation json 选填，房规, 如果执行者名下还没有对应名字的规则，则创建，否则，更新之, 有如下值：
                    name string 选填，规则名称（如果没填则是name=default）
                    tenstrikemax	int 选填，大成功值
                    fiascomax int 选填，大失败值

* @apiExample json
* {
    "name":"room1",
    "kpname":"墨竹",
    "era":"现代",
    "site":"中国",
    "language":"汉语",
    "max":2,
    "min":3,
    "background":"...",
    "description":"...",
    "recommendskills":"...",
    "unrecommendskills":"...",
    "limitedskills":"...",
    "flag":true,
    "acceptold":false,   //是否接受老卡
    "car":"...",         //车卡
    "gugu":34,        //咕咕值
    "friendly": 54,    //友好值
	"friend":true,   //是否需要是好友
	"checkother":true,//能否查看其它玩家角色卡
    "regulation":{
        "name":"r1",
        "tenstrikemax":5,   //大成功最大值
        "fiascomax":90      //大失败最大值
    },
    "ruleid":1,
    "gametimes":["1558080000-1558083600", "1558170000-1558173600"],
    "tags":["大西洋", "海盗"]
}

* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiParam data json 新建的房间的信息，除了上面所填，还包括后台生成的房间号
* @apiExample json
* {
    "errcode":0,
    "errmsg":"ok",
    "data":{
        "id":"F00000001",  // 房间号
        "name":"room1",
        "era":"现代",
        "site":"中国",
        "language":"汉语",
        "max":2,
        "min":3,
        "background":"...",
        "description":"...",
        "recommendskills":"...",
        "unrecommendskills":"...",
        "limitedskills":"...",
        "flag":true,
        "acceptold":false,   //是否接受老卡
        "car":"...",         //车卡
        "gugu":34,        //咕咕值
        "friendly": 54,    //友好值
		"friend":true,   //是否需要是好友
		"checkother":true,//能否查看其它玩家角色卡
        "regulation":{
            "name":"r1",
            "tenstrikemax":5,   //大成功最大值
        	"fiascomax":90      //大失败最大值
        },
        "ruleid":1,
        "gametimes":["1558080000-1558083600", "1558170000-1558173600"],
        "tags":["大西洋", "海盗"]
    }
}

func GetRoomNpc(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}

	rid := c.Query("rid")
	// role, _ := req["roleid"].(string)

	rs := []Room{}
	DB.GetByFieldM(&rs, "id", rid, RoomTable)
	if len(rs) == 0 {
		R.RJson(c, "NOT_FOUND")
		return
	}

	rr := Npccards{}
	DB.GetByFieldM(&rr, "rid", rid, NpccardsTable)
	// rr := Npccards{}
	// DB.Db.Table(NpccardsTable).Where("rid = ?", rid).First(&rr)
	// if rr.Id == 0 {
	// 	R.RJson(c, "NOT_FOUND")
	// 	return
	// } else {
	// 	R.RData(c, rr)
	// }

	R.RData(c, rr)

	return
}

/*
* @api PUT /api/v1/room/status 设置房间状态（用于开始、暂停、停止游戏）
* 只用房主才可以操作
* @apiGroup room
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token

* @apiRequest json
* @apiParam id	string 必填，房间id
* @apiParam status	int 必填，0：停止游戏；1：开始游戏；2：暂停游戏
* @apiExample json
* {
    "status":1
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
func PutRoomStatus(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}

	req := make(map[string]interface{})

	err := c.BindJSON(&req)
	if err != nil {
		fmt.Println(err)
		R.RJson(c, "INVALID_PARAM")
		return
	}

	rid, _ := req["id"].(string)
	s, _ := req["status"].(float64)

	status := int(s)

	rs := []Room{}
	DB.GetByFieldM(&rs, "id", rid, RoomTable)
	if len(rs) == 0 {
		R.RJson(c, "NOT_FOUND")
		return
	}

	uid := c.Query("userid")
	userid, _ := strconv.Atoi(uid)

	if rs[0].Userid != int64(userid) {
		R.RJson(c, "NO_PERMISSION")
		return
	}

	rs[0].Status = status
	DB.UpdateByFieldM(&rs[0], "id", rs[0].Id, RoomTable)

	R.RJson(c, "OK")
}

/*
* @api DELETE /api/v1/room/room 删除房间
* 只用房主及管理员才可以操作
* @apiGroup room
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery id string 必填，房间id

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
func DeleteRoom(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}

	rid := c.Query("id")
	if rid == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	//
	//rs := []Room{}
	//DB.GetByFieldM(&rs, "id", rid, RoomTable)
	//if len(rs) == 0 {
	//	R.RJson(c, "NOT_FOUND")
	//	return
	//}
	//
	//uid := c.Query("userid")
	//userid, _ := strconv.Atoi(uid)
	//
	//if rs[0].Userid != int64(userid) && !IsAdmin(int64(userid)) {
	//	R.RJson(c, "NO_PERMISSION")
	//	return
	//}

	err := DB.DeleteByFieldM("id", rid, RoomTable)
	if err != nil {
		R.RJson(c, "ERROR_DB")
		return
	}

	R.RJson(c, "OK")
}

/*
* @api PUT /api/v1/room/room 修改房间
* 只用管理员及房主才可以操作
* @apiGroup room
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token
* @apiQuery id string 必填，想要修改房间的ID

* @apiRequest json
* @apiParam name string 选填，房间（游戏）名称
* @apiParam kpname string 选填，房主（KP）名称
* @apiParam password string 选填，房间密码
* @apiParam icon string 选填，房间头像
* @apiParam era	string 选填，时代
* @apiParam language	string 选填，母语
* @apiParam site string 选填，地点
* @apiParam max int 选填，最多人数
* @apiParam min int 选填，最少人数
* @apiParam background string 选填，背景
* @apiParam description string 选填，特别说明
* @apiParam recommendskills string 选填，推荐技能
* @apiParam unrecommendskills string 选填，不推荐技能
* @apiParam limitedskills string 选填，技能限定
* @apiParam flag bool 选填，是否满员自动开始，不填默认为false
* @apiParam ruleid int 选填，规则id
* @apiParam gametimes []string 选填，游戏时间, 值统一为时间戳，例如["1558080000-1558083600", "1558170000-1558173600"]
* @apiParam tags []string 选填，标签
* @apiParam moduleid int 选填，模组的资源ID
* @apiParam acceptold int 选填，是否接受老卡
* @apiParam car string 选填，车卡
* @apiParam gugu int 选填，咕咕值
* @apiParam friendly int 选填，友好值
* @apiParam friend bool 选填，是否需要是好友
* @apiParam checkother bool 选填，能否查看其它玩家角色卡
* @apiParam regulation json 选填，房规, 如果执行者名下还没有对应名字的规则，则创建，否则，更新之, 有如下值：
                    name string 选填，规则名称（如果没填则是name=default）
                    tenstrikemax	int 选填，大成功值
                    fiascomax int 选填，大失败值

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
func PutRoom(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}

	rid := c.Query("id")
	if rid == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	userid:=999999
	rs := []Room{}
	DB.GetByFieldM(&rs, "id", rid, RoomTable)
	if len(rs) == 0 {
		R.RJson(c, "NOT_FOUND")
		return
	}
	//
	//uid := c.Query("userid")
	//userid, _ := strconv.Atoi(uid)
	//
	//if rs[0].Userid != int64(userid) && !IsAdmin(int64(userid)) {
	//	R.RJson(c, "NO_PERMISSION")
	//	return
	//}

	roomreq := make(map[string]interface{})
	err := c.Bind(&roomreq)
	if err != nil {
		fmt.Println(err)
		R.RJson(c, "INVALID_PARAM")
		return
	}

	needSetSlice := false
	tags := Strings{}
	gametimes := Strings{}
	tagsTmp, ok := roomreq["tags"]
	if ok {
		data, err := json.Marshal(tagsTmp)
		if err == nil {
			err = json.Unmarshal(data, &tags)
			if err == nil {
				needSetSlice = true
			}
		}
	}
	gametimesTmp, ok := roomreq["gametimes"]
	if ok {
		data, err := json.Marshal(gametimesTmp)
		if err == nil {
			err = json.Unmarshal(data, &gametimes)
			if err == nil {
				needSetSlice = true
			}
		}
	}
	delete(roomreq, "tags")
	delete(roomreq, "gametimes")

	if needSetSlice {
		roomreqStruct := RoomReq{}
		roomreqStruct.Tags = tags
		roomreqStruct.Gametimes = gametimes
		roomreqStruct.Room.Password = strByXOR(roomreqStruct.Room.Password)
		e := DB.Db.Table(RoomTable).Where("id = ?", rid).Updates(roomreqStruct)
		if e.Error != nil {
			R.RJson(c, "ERROR_DB")
			return
		}
	}

	delete(roomreq, "id")
	roomreq["updatetime"] = time.Now().Unix()

	reg, ok := roomreq["regulation"]
	if ok {
		regulationStruct, ok := reg.(map[string]interface{})
		if ok {
			regulation := Regulation{}

			regulation.Userid = int64(userid)

			name, ok := regulationStruct["name"].(string)
			if ok {
				regulation.Name = name
			}

			tenstrikemax, ok := regulationStruct["tenstrikemax"]
			if ok {
				tmp, ok := tenstrikemax.(float64)
				if ok {
					regulation.Tenstrikemax = int(tmp)
				} else {
					regulation.Tenstrikemax, _ = tenstrikemax.(int)
				}

			}
			fiascomax, ok := regulationStruct["fiascomax"]
			if ok {
				tmp, ok := fiascomax.(float64)
				if ok {
					regulation.Fiascomax = int(tmp)
				} else {
					regulation.Fiascomax, _ = fiascomax.(int)
				}
			}
			ec, em := CreateOrPutRegulation(&regulation)
			if ec != 0 {
				R.Jr(c, http.StatusOK, gin.H{"errcode": ec, "errmsg": em})
				return
			}

			roomreq["regulationid"] = regulation.Id

			delete(roomreq, "regulation")
		}
	}
	var pass string
	pass = roomreq["password"].(string)
	roomreq["password"] = strByXOR(pass)
	e := DB.Db.Table(RoomTable).Where("id = ?", rid).Updates(roomreq)
	if e.Error != nil {
		R.RJson(c, "ERROR_DB")
		return
	}

	R.RJson(c, "OK")
}

/**
* @api GET /api/v1/room/roompass 检测房间密码是否正确

* @apiGroup room
* @apiParam userid int 必填，执行请求的用户id
* @apiParam token string 必填，执行请求的用户token
* @apiParam id string 选填，房间号
* @apiParam name string 选填，房间名称（id和name必填一个，如果两个都填，以id为准）
* @apiParam password string 必填，要检查的对应房间的密码

* @apiSuccess 200 OK
* @apiExample json
* {
    "data": "true",
    "errcode": 0,
    "errmsg": "ok"
*}
* @apiError 200 OK

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
* 房间没找到
* {
    "errcode":30015,
    "errmsg":"not found"
* }
*/
func CheckRoomPass(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	//userid, err := strconv.Atoi(c.Query("userid"))
	//if err != nil {
	//	R.RJson(c, "INVALID_PARAM")
	//	return
	//}
	if /*IsAdmin(int64(userid))*/true {
		flag = true
		R.RData(c, flag)
	} else {
		isEmpty := false
		id := c.Query("id")
		name := c.Query("name")
		pass := c.Query("password")
		if pass == "" {
			isEmpty = true
		}
		if id == "" && name == "" {
			R.RJson(c, "INVALID_PARAM")
			return
		}
		if id != "" {
			ismatch := checkRoomId(id)
			if !ismatch {
				R.RJson(c, "INVALID_PARAM")
				return
			}
		}
		rs := []Room{}
		if id != "" {
			DB.GetByFieldM(&rs, "id", id, RoomTable)
			if len(rs) == 0 {
				R.RJson(c, "NOT_FOUND")
				return
			}
		} else {
			DB.GetByFieldM(&rs, "name", name, RoomTable)
			if len(rs) == 0 {
				R.RJson(c, "NOT_FOUND")
				return
			}
		}

		if pass == strByXOR(rs[0].Password) {
			flag = true
			R.RData(c, flag)
		} else if isEmpty && rs[0].Password == "" {
			flag = true
			R.RData(c, flag)
		} else {
			flag = false
			R.RData(c, flag)
		}
	}

}
func strByXOR(message string) string { //一个简单的加密解密函数
	keywords := "84161615849" //随便写的密钥
	messageLen := len(message)
	keywordsLen := len(keywords)
	result := ""
	for i := 0; i < messageLen; i++ {
		result += string(message[i] ^ keywords[i%keywordsLen])
	}
	return result
}
func RoomInit() {
	DB.CreateTableM(Room{})
	RoomTable = DB.GetTableNameM(Room{})
	DB.Db.Table(RoomTable).AutoMigrate(&Room{})

	DB.CreateTableM(Roomrole{})
	RoomroleTable = DB.GetTableNameM(Roomrole{})
	DB.Db.Table(RoomroleTable).AutoMigrate(&Roomrole{})
	

	// 初始化 RoomIdNum
	r := Room{}
	DB.GetLastM(&r, RoomTable)
	if r.Id != "" {
		fmt.Sscanf(r.Id, "F%08d", &RoomIdNum)
	} else {
		RoomIdNum = 0
	}
}

func NpcInit() {
	DB.CreateTableM(Npccards{})
	NpccardsTable = DB.GetTableNameM(Npccards{})
	DB.Db.Table(NpccardsTable).AutoMigrate(&Npccards{})

	r := Npccards{}
	DB.GetLastM(&r, NpccardsTable)
}
