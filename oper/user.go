package main

import (
	"encoding/json"
	"fmt"
	"miao/Utils/DB"
	"miao/Utils/R"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	CaptchaExpire = 5  // 5 minutes
	TokenExpire   = 30 // 30 days
	PswExpire     = 60 // 60 seconds
)

var (
	OperuserTable         string
	UserTable             string
	TokenTable            string
	OpersessionTable      string
	OperCryptoTable       string
	userRelationshipTable string
)

type Operuser struct {
	Id     int64  `json:"userid" gorm:"index"`
	Name   string `json:"name" gorm:"index"`
	Cnname string `json:"cnname"`
	Openid string `json:"openid"`
	Email  string `json:"email" gorm:"index"`
	Phone  string `json:"phone"`
	Passwd string `json:"passwd"`
	IM     string `json:"im"`
	QQ     string `json:"qq"`
	//Rolecard Strings `json:"rolecard" sql:"type:json"`
	Sex      string `json:"sex"`
	Touxiang string `json:"touxiang"`
	Birthday int64  `json:"birthday"`
	Sign     string `json:"sign"`
	DiceID   int64  `json:"dice_id"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}
type UserRelationships struct {
	Id           int64  `json:"id" gorm:"index"`
	UserId       int64  `json:"userid" gorm:"index"`
	FriendId     int64  `json:"friendid" gorm:"index"`
	Cid          int64  `json:"cid" gorm:"column:cid"`
	UserName     string `json:"username" gorm:"column:user_name"`
	FriendName   string `json:"friendname" gorm:"column:friend_name"`
	UserRemark   string `json:"userremark" gorm:"column:user_remark"`
	FriendRemark string `json:"friendremark" gorm:"column:friend_remark"`
	UserEmail    string `json:"useremail" gorm:"column:user_email"`
	FriendEmail  string `json:"friendemail" gorm:"column:friend_email"`
	UserImage    string `json:"userimage" gorm:"column:user_image"`
	FriendImage  string `json:"friendimage" gorm:"column:friend_image"`
	Created      string `json:"created"`
	Updated      string `json:"updated"`
}
type Opersession struct {
	Id      int64  `json:"id" gorm:"index"`
	Uid     int64  `json:"uid" gorm:"index"`
	Sig     string `json:"sig"`
	Expired int
}

type Token struct {
	Id      int64  `json:"id" gorm:"index"`
	Uid     int64  `json:"uid" gorm:"index"`
	Token   string `json:"token"`
	Expired int
}

type Opercrypto struct {
	Id       int64  `json:"id" gorm:"index"`
	Uid      int64  `json:"uid" gorm:"index"`
	Password string `json:"password"`
}

type UserResp struct {
	Id   int64    `json:"userid,omitempty"`
	Name string   `json:"name,omitempty"`
	Sig  string   `json:"sig,omitempty"`
	Info Operuser `json:"info,omitempty"`
}

type userReq struct {
	Operuser
}

type passwdReq struct {
	Name      string `json:"name"`
	Pswmd5    string `json:"pswmd5"`
	Timestamp int    `json:"timestamp"`
}

type User struct {
	Id         int64   `json:"userid" gorm:"index"`
	Name       string  `json:"name" gorm:"index"`
	Cnname     string  `json:"cnname"`
	Openid     string  `json:"openid"`
	Email      string  `json:"email" gorm:"index"`
	Phone      string  `json:"phone"`
	Passwd     string  `json:"passwd"`
	IM         string  `json:"im"`
	QQ         string  `json:"qq"`
	Rolecard   Strings `json:"rolecard" sql:"type:json"`
	Sex        string  `json:"sex"`
	Touxiang   string  `json:"touxiang"`
	Birthday   int64   `json:"birthday"`
	Sign       string  `json:"sign"`
	DiceID     int64   `json:"dice_id"`
	PersonAuth int     `json:"person_auth"` //0未认证
	Manager    int     `json:"manager"`     //0 非管理员  1管理员
	Created    string  `json:"created"`
	Updated    string  `json:"updated"`
}

/*
* @api POST /api/v1/user/login 登录
* @apiGroup user
* @apiRequest json
* @apiParam name string 必填，执行请求的用户email
* @apiParam pswmd5 string 必填，执行请求的用户password的md5值
* @apiParam timestamp int 必填，执行请求时的时间戳，该时间戳到服务器处理时，不能超过60s
* @apiExample json
* {
    "email":"1234@test.com",
    "pswmd5":"12345678901234567890123456789012",
    "timestamp":156800000
* }

* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiParam data json 用户id、name、登录sig和token
* @apiExample json
* {
    "errcode":0,
    "errmsg":"ok",
    "data":{
        "name":"bob",
        "token":"aaaa",
        "sig":"bbbb",
        "userid":1,
        "info":{
            "userid":1,
            "name":"bob",
            "cnname":"bo",
            "email":"test@test.com",
            "phone":88888888,
            "im":12345,
            "qq":54321,
            "passwd":"*****",
            "sex":"male",             //性别，male：男、female：女、其他为保密
            "touxiang":"12345678901234567890123456789012",    // 头像MD5值
            "birthday":15002939021,   //生日时间戳
            "sign":"asdfasdfasdfadsf"    //签名
        }
    }
}

* @apiError 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* 服务器内部错误
* {
    "errcode":30101,
    "errmsg":"unknown user"
* }
* @apiExample json
* 密码错误
* {
    "errcode":30103,
    "errmsg":"wrong password"
* }
* @apiExample json
* 参数错误
* {
    "errcode":30002,
    "errmsg":"invalid parameter"
* }
* @apiExample json
* 时间戳超时
* {
    "errcode":30004,
    "errmsg":"http request timeout"
* }
*/
func PostLogin(c *gin.Context) {

	info := passwdReq{}
	err := c.Bind(&info)
	if err != nil || info.Pswmd5 == "" || info.Timestamp == 0 {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	users := []Operuser{}
	DB.GetByFieldM(&users, "name", info.Name, OperuserTable)
	if len(users) == 0 {
		R.RJson(c, "UNKNOWN_USER")
		return
	}
	user := users[0]
	id_int64 := user.Id

	//DEBan() //检测有无解封账号
	//检测封号
	//bts := []Ban_user{}
	//DB.GetByFieldM(&bts, "uid", user.Id, Ban_usertable)
	//if len(bts) != 0 {
	//	bt := bts[0]
	//	if bt.Suspend == true {
	//		R.Jr(c, http.StatusOK, gin.H{"errcode": 32000, "errmsg": "USER_IS_BANED", "data": bt})
	//		return
	//	}
	//}

	userid := fmt.Sprintf("%d", id_int64)

	cys := []Opercrypto{}
	DB.GetByFieldM(&cys, "uid", id_int64, OperCryptoTable)
	if len(cys) == 0 {
		R.RJson(c, "WRONG_PASSWORD")
		return
	}

	cy := cys[0]

	fmt.Println(info.Pswmd5)

	err = bcrypt.CompareHashAndPassword([]byte(cy.Password), []byte(info.Pswmd5))
	if err != nil {
		R.RJson(c, "WRONG_PASSWORD")
		return
	}

	s := Opersession{}

	DB.GetByFieldM(&s, "uid", id_int64, OpersessionTable)
	s.Uid = id_int64
	s.Sig = GenerateToken()
	s.Expired = int(time.Now().Unix()) + (3600 * 24 * TokenExpire)
	if s.Id != 0 {
		DB.SaveM(&s, OpersessionTable)
	} else {
		DB.CreateM(&s, OpersessionTable)
	}

	uid_cookie := &http.Cookie{
		Name:  userid,
		Value: s.Sig,
		Path:  "/",
	}

	http.SetCookie(c.Writer, uid_cookie)

	R.RData(c, UserResp{
		Id:   id_int64,
		Sig:  s.Sig,
		Name: user.Name,
		Info: user,
	})
}

/*
 * @api GET /api/v1/user/user 获取用户信息
* @apiGroup user
* @apiRequest query
 * @apiGroup user
 * @apiQuery userid int 必填，用户id
 * @apiQuery token string 必填，用户token
 * @apiQuery id string 选填，要查询的用户id，id和email必填一项
 * @apiQuery email string 选填，要查询的用户邮箱，id和email必填一项

 * @apiSuccess 200 OK
 * @apiParam errcode int 错误代码
 * @apiParam errmsg string 错误信息
 * @apiParam data json 用户信息
 * @apiExample json
 * {
    "errcode":0,
    "errmsg":"ok",
    "data":{
        "userid":1,
        "name":"bob",
        "cnname":"bo",
        "email":"test@test.com",
        "phone":88888888,
        "im":12345,
        "qq":54321,
        "passwd":"*****",      // 已经设置密码则为*****，没设置密码则为空
        "rolecard":["coc7_2002", "coc7_4123"],   //（弃用，用户所拥有的角色卡在rolecard list接口中查看）
        "sex":"male",             //性别，male：男、female：女、其他为保密
        "touxiang":"12345678901234567890123456789012",    // 头像MD5值
        "birthday":15002939021,   //生日时间戳
        "sign":"asdfasdfasdfadsf"    //签名
    }
}
 * @apiError 200 OK
 * @apiParam errcode int 错误代码
 * @apiParam errmsg string 错误信息
* @apiExample json
* 校验token和session失败
* {
   "errcode":30305,
   "errmsg":"auth failed"
* }
 * @apiExample json
 * 无效参数
 * {
    "errcode":30002,
    "errmsg":"invalid parameter"
 * }
 * @apiExample json
 * 未知的用户
 * {
    "errcode":30101,
    "errmsg":"unknown user"
 * }
 * @apiExample json
 * 没有权限
 * {
    "errcode":30005,
    "errmsg":"no permission"
 * }
*/
func GetUser(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}

	uid := c.Query("userid")
	if uid == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}

	user := Operuser{}

	if uid != "" {
		id_int, err := strconv.Atoi(uid)
		if err != nil {
			R.RJson(c, "INVALID_PARAM")
			return
		}
		DB.GetM(&user, int64(id_int), OperuserTable)
		if user.Id == 0 {
			R.RJson(c, "UNKNOWN_USER")
			return
		}
	}

	R.RData(c, user)
}

/*
 * @api GET /api/v1/user/users 获取用户列表及信息
 * @apiGroup user
* @apiRequest query
 * @apiQuery userid int 必填，请求用户列表的用户的id
 * @apiQuery token string 必填，用户token
 * @apiQuery ids array 选填，指明要获取的用户id列表
 * @apiSuccess 200 OK
 * @apiParam errcode int 错误代码
 * @apiParam errmsg string 错误信息
 * @apiParam data array 用户信息列表，列表中的每一个元素是json对象
 * @apiExample json
 * {
    "errcode":0,
    "errmsg":"ok",
    "data":[
            {
                "userid":1,
                "name":"bob",
                "cnname":"",
                "email":"test@test.com",
                "phone":88888888,
                "im":12345,
                "qq":54321,
                "sex":"male",             //性别，male：男、female：女、其他为保密
                "touxiang":"12345678901234567890123456789012",    // 头像MD5值
                "birthday":15002939021,   //生日时间戳
                "sign":"asdfasdfasdfadsf"    //签名
            },
            {
                "userid":2,
                "email":"test@test.com"
            }
    ]
 * }
 * @apiError 200 OK
 * @apiParam errcode int 错误代码
 * @apiParam errmsg string 错误信息
* @apiExample json
* 校验token和session失败
* {
   "errcode":30305,
   "errmsg":"auth failed"
* }
 * @apiExample json
 * 无效参数（ids参数存在，但是值不正确时）
 * {
    "errcode":30002,
    "errmsg":"invalid parameter"
 * }
 * @apiExample json
 * 服务器内部错误
 * {
    "errcode":30002,
    "errmsg":"system internal error"
 * }
*/
func GetUsers(c *gin.Context) {
	if !Check(c) {
		R.RJson(c, "AUTH_FAILED")
		return
	}

	ids := c.Query("ids")

	pList := []Operuser{}

	if ids != "" {
		var ids_int []int64

		err := json.Unmarshal([]byte(ids), &ids_int)
		if err != nil {
			R.RJson(c, "INVALID_PARAM")
			return
		}

		for _, id := range ids_int {
			var one Operuser
			DB.GetM(&one, id, OperuserTable)
			if one.Id != 0 {
				pList = append(pList, one)
			}
		}

	} else {
		DB.GetAllM(&pList, OperuserTable)
	}

	if len(pList) == 0 {
		R.RData(c, pList)
		return
	}

	R.RData(c, pList)
}

/*
 * @api GET /api/v1/user/users 定量分页导出用户列表
 * @apiGroup user
* @apiRequest query
 * @apiQuery userid int 必填，请求用户列表的管理人员id
 * @apiQuery token string 必填，管理人员token
 * @apiQuery pageSize int 必填，每页的用户数量
 * @apiQuery currentPage int 必填，当前所在页数

 * @apiSuccess 200 OK
 * @apiParam errcode int 错误代码
 * @apiParam errmsg string 错误信息
 * @apiExample json
 * {
    "errcode":0,
    "errmsg":"ok",
    "data":[
            {
                "userid":1,
                "name":"bob",
                "cnname":"",
                "email":"test@test.com",
                "phone":88888888,
                "im":12345,
                "qq":54321,
                "sex":"male",             //性别，male：男、female：女、其他为保密
                "touxiang":"12345678901234567890123456789012",    // 头像MD5值
                "birthday":15002939021,   //生日时间戳
                "sign":"asdfasdfasdfadsf"    //签名
            },
            {
                "userid":2,
                "email":"test@test.com"
            }
    ]
 * }
 * @apiError 200 OK
 * @apiParam errcode int 错误代码
 * @apiParam errmsg string 错误信息
* @apiExample json
* 校验token和session失败
* {
   "errcode":30305,
   "errmsg":"auth failed"
* }
 * @apiExample json
 * 服务器内部错误
 * {
    "errcode":30002,
    "errmsg":"system internal error"
 * }
*/
func pagingGetUsers(c *gin.Context) {
	if !Check(c) {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	//|| !IsAdmin(int64(Msg{}.Senderid))

	pageSize := c.Query("pageSize")
	currentPage := c.Query("currentPage")

	list := []Operuser{}

	if pageSize == "" || currentPage == "" {
		R.RJson(c, "NECESSARY_MARAMETER_LACK")
		return
	}

	var pageSizeInt int
	err := json.Unmarshal([]byte(pageSize), &pageSizeInt)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}
	var currentPageInt int
	err = json.Unmarshal([]byte(currentPage), &currentPageInt)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}

	DB.Db.Scopes(DB.Paginate(currentPageInt, pageSizeInt)).Find(&list)

	if len(list) == 0 {
		R.RData(c, list)
		return
	}

	R.RData(c, list)
}


func UserInit() {
	DB.CreateTableM(User{})
	UserTable = DB.GetTableNameM(User{})
	DB.Db.Table(UserTable).AutoMigrate(&User{})

	DB.CreateTableM(Operuser{})
	OperuserTable = DB.GetTableNameM(Operuser{})
	DB.Db.Table(OperuserTable).AutoMigrate(&Operuser{})

	//DB.CreateTableM(Token{})
	//TokenTable = DB.GetTableNameM(Token{})
	//DB.Db.Table(TokenTable).AutoMigrate(&Token{})

	DB.CreateTableM(Operuser{})
	OperuserTable = DB.GetTableNameM(Operuser{})
	DB.Db.Table(OperuserTable).AutoMigrate(&Operuser{})

	DB.CreateTableM(Opercrypto{})
	OperCryptoTable = DB.GetTableNameM(Opercrypto{})
	DB.Db.Table(OperCryptoTable).AutoMigrate(&Opercrypto{})

	DB.CreateTableM(Opersession{})
	OpersessionTable = DB.GetTableNameM(Opersession{})
	DB.Db.Table(OpersessionTable).AutoMigrate(&Opersession{})

	//DB.CreateTableM(UserRelationships{})
	//userRelationshipTable = DB.GetTableNameM(UserRelationships{})
	//DB.Db.Table(userRelationshipTable).AutoMigrate(&UserRelationships{})
}

/*
* @api PUT /api/v1/user/user 修改一个user信息(包括修改密码)
* @apiGroup user
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token

* @apiRequest json
* @apiParam userid int 必填，要修改的user id
* @apiParam name string 选填，要修改的name
* @apiParam cnname string 选填，要修改的cnname
* @apiParam phone string 选填，要修改的123455
* @apiParam passwd string 选填，要修改的Passwd
* @apiParam im string 选填，要修改的im
* @apiParam qq string 选填，要修改的qq
* @apiParam touxiang string 选填，用户头像md5值
* @apiParam birthday int64 选填，生日时间戳
* @apiParam sex string 选填，用户性别 male：男；female：女；其他非空字符串：保密
* @apiParam sign string 选填，用户签名
* @apiExample json
* {
    "id":1,
    "name":"wubo",
    "cnname":"wubo",
    "phone":"123455",
    "Passwd":"1245677",
    "im":"1245677",
    "qq":"1245677",
    "rolecard":["coc7_2342"],
    "userid":1,
    "sex":"male",             //性别，male：男、female：女、其他为保密
    "touxiang":"12345678901234567890123456789012",    // 头像MD5值
    "birthday":15002939021,   //生日时间戳
    "sign":"asdfasdfasdfadsf",    //签名
    "flag":true
* }

* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* {
    "errcode":0,
    "errmsg":"ok"
* }

* @apiError 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* 鉴权失败
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
* @apiExample json
* 未知的用户
* {
    "errcode":30101,
    "errmsg":"unknown user"
* }
* @apiExample json
* 密码加密错误
* {
    "errcode":30012,
    "errmsg":"bcrypt error"
* }
*/
func PutUser(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}

	p := userReq{}
	err := c.BindJSON(&p)
	if err != nil || (p.Id == 0) {
		fmt.Println(err, p.Id, p.Name)
		R.RJson(c, "INVALID_PARAM")
		return
	}

	new := Operuser{}
	DB.GetM(&new, p.Id, OperuserTable)
	if new.Id == 0 {
		R.RJson(c, "UNKNOWN_USER")
		return
	}

	new.Updated = time.Now().Format("2006-01-02 15:04:05")
	new.Name = p.Name
	if p.Name != "" {
		new.Name = p.Name
	}
	if p.Cnname != "" {
		new.Cnname = p.Cnname
	}
	if p.IM != "" {
		new.IM = p.IM
	}
	if p.QQ != "" {
		new.QQ = p.QQ
	}
	if p.Phone != "" {
		new.Phone = p.Phone
	}
	if p.Sex != "" {
		new.Sex = p.Sex
	}
	if p.Birthday != 0 {
		new.Birthday = p.Birthday
	}
	if p.Sign != "" {
		new.Sign = p.Sign
	}
	if p.Touxiang != "" {
		new.Touxiang = p.Touxiang
	}
	if p.Passwd != "" {
		cy := Opercrypto{}

		DB.GetByFieldM(&cy, "uid", p.Id, OperCryptoTable)

		hash, err := bcrypt.GenerateFromPassword([]byte(p.Passwd), bcrypt.DefaultCost)
		if err != nil {
			R.RJson(c, "BCRYPT_ERROR")
			return
		}

		cy.Uid = p.Id
		cy.Password = string(hash)

		err = DB.SaveM(&cy, OperCryptoTable)
		if err != nil {
			R.RJson(c, "ERROR_DB")
			return
		}

		new.Passwd = "******"
	}

	err = DB.SaveM(&new, OperuserTable)
	if err != nil {
		R.RJson(c, "ERROR_DB")
		return
	}

	R.RJson(c, "OK")
}


/**
* @api GET /api/v1/client_user/set_manager 将前台用户设置成管理员
* @apiGroup 文章帖子
* @apiParam userid int 必填，执行请求的用户id
* @apiParam token string 必填，执行请求的用户token
* @apiParam manager_uid string 前台管理员用户id

* @apiSuccess 200 OK
* @apiExample json
* {
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
func SetManagerUser(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	managerUID, _ := strconv.ParseInt(c.Query("manager_uid"), 10, 64)
	if managerUID == 0 {
		fmt.Printf("manager_uid is not int managerUID=%v\n", managerUID)
		R.RJson(c, "INVALID_PARAM")
		return
	}
	nowFormat:=time.Now().Format("2006-01-02 15:04:05")
	DB.Db.Table(UserTable).Where("id=?", managerUID).Updates(map[string]interface{}{
		"manager": 1,
		"updated":nowFormat,
	})
	R.RJson(c, "OK")
	return
}

/**
* @api GET /api/v1/client_user/remove_manager 移除后台管理员用户列表
* @apiGroup 文章帖子
* @apiParam userid int 必填，执行请求的用户id
* @apiParam token string 必填，执行请求的用户token
* @apiParam manager_uid string 前台管理员用户id
* @apiSuccess 200 OK
* @apiExample json
* {
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
func RemoveManagerUser(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	managerUID, _ := strconv.ParseInt(c.Query("manager_uid"), 10, 64)
	if managerUID == 0 {
		fmt.Printf("manager_uid is not int managerUID=%v\n", managerUID)
		R.RJson(c, "INVALID_PARAM")
		return
	}
	DB.Db.Table(UserTable).Where("id=?", managerUID).Updates(map[string]interface{}{
		"manager": 0,
	})
	R.RJson(c, "OK")
	return
}


/**
* @api GET /api/v1/client_user/manager_list 后台查看前台管理员用户列表
* @apiRequest query
* @apiGroup 文章帖子
* @apiParam userid int 必填，执行请求的用户id
* @apiParam token string 必填，执行请求的用户token
* @apiParam page_index string 选填，页码
* @apiParam page_size string 选填，每页数量

* @apiSuccess 200 OK
* @apiExample json
* {
    "data": {
        "list":[
         {"name":"xx","id":1,"touxiang":"xx"},
         {"name":"xx","id":1,"touxiang":"xx"},

         ],
         "total":2
    },
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
func ManagerUserList(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	pageIndex, _ := strconv.Atoi(c.Query("page_index"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	users := make([]User, 0)
	q := DB.Db.Table(UserTable).Where("manager=1")
	q.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Order("Updated desc").Find(&users)
	var count int64
	q.Count(&count)

	R.RJsonData(c, "OK", map[string]interface{}{
		"list":  users,
		"total": count,
	})
	return
}


/**
* @api GET GET /api/v1/client_user/user 获取前台管理员用户信息
* @apiGroup 文章帖子
* @apiQuery userid int 必填，用户id
* @apiQuery token string 必填，用户token
* @apiQuery manager_uid string 选填，要查询的用户id，id和email必填一项
* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiParam data json 用户信息
* @apiExample json
* {
    "errcode":0,
    "errmsg":"ok",
    "data":{
        "userid":1,
        "name":"bob",
        "cnname":"bo",
        "email":"test@test.com",
        "phone":88888888,
        "im":12345,
        "qq":54321,
        "passwd":"*****",      // 已经设置密码则为*****，没设置密码则为空
        "rolecard":["coc7_2002", "coc7_4123"],   //（弃用，用户所拥有的角色卡在rolecard list接口中查看）
        "sex":"male",             //性别，male：男、female：女、其他为保密
        "touxiang":"12345678901234567890123456789012",    // 头像MD5值
        "birthday":15002939021,   //生日时间戳
        "sign":"asdfasdfasdfadsf"    //签名
    }
}
* @apiError 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiExample json
* 校验token和session失败
* {
   "errcode":30305,
   "errmsg":"auth failed"
* }
 * @apiExample json
 * 无效参数
 * {
    "errcode":30002,
    "errmsg":"invalid parameter"
 * }
 * @apiExample json
 * 未知的用户
 * {
    "errcode":30101,
    "errmsg":"unknown user"
 * }
 * @apiExample json
 * 没有权限
 * {
    "errcode":30005,
    "errmsg":"no permission"
 * }
*/
func GetManagerUser(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}

	uid := c.Query("manager_uid")
	if uid == "" {
		R.RJson(c, "INVALID_PARAM")
		return
	}

	user := User{}

	if uid != "" {
		id_int, err := strconv.Atoi(uid)
		if err != nil {
			R.RJson(c, "INVALID_PARAM")
			return
		}
		DB.GetM(&user, int64(id_int), UserTable)
		if user.Id == 0 {
			R.RJson(c, "UNKNOWN_USER")
			return
		}
	}

	R.RData(c, user)
}
