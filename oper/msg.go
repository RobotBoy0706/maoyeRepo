package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"miao/Utils/R"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Msg struct {
	Id          int64  `json:"id" gorm:"index"`
	Type        string `json:"type" gorm:"index"` // "system", "personal"
	Level       int    `json:"level"`             // 1,2,3
	Resourcemd5 string `json:"resourcemd5"`
	Data        string `json:"data"`
	Receiverid  int64  `json:"receiverid"` //接收者id，如果是系统消息等公告，则填0
	Senderid    int64  `json:"senderid"`   //发送者id
	Sender      string `json:"sender"`     //发送者名称
	Status      int    `json:"status"`     //消息状态，0：未读，1：已读
	Createtime  int64  `json:"createtime"` //消息创建时间
}


/*
* @api POST /api/v1/msg/msg 新增消息
新增系统消息会通过cmd=msg的websocket消息发送给在线用户，具体消息在extend字段中，不在线用户需要在登陆时，前端通过接口获取消息列表
新增私信会通过cmd=msg的websocket消息发送给具体在线用户，具体消息在extend字段中，如果用户不在线，需要在登陆时，前端通过接口获取消息列表
* @apiGroup msg
* @apiQuery userid int 必填，执行请求的用户id
* @apiQuery token string 必填，执行请求的用户token

* @apiRequest json
* @apiParam data	string 必填，消息内容
* @apiParam type	string 选填，消息类别，system：系统消息(只用管理员可以发送)；personal：私信。默认是系统消息；friendRequest：好友申请
* @apiParam level	int 选填，消息程度， 正常（普通）：1; 警告（严重）：2; 错误（紧急）：3。默认是1
* @apiParam resourcemd5 string 选填，该条消息包含图片，这个填写图片在resource接口中的md5值
* @apiParam receiverid	int64 选填，表示消息要发给谁；如果是系统消息，则可以不填，或者填0
* @apiExample json
* {
    "data":"hello",
    "type":"system",
    "level":1
}

* @apiSuccess 200 OK
* @apiParam errcode int 错误代码
* @apiParam errmsg string 错误信息
* @apiParam data int 新建的msg id
* @apiExample json
* {
    "errcode":0,
    "errmsg":"ok",
	"data":1
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
func SystemMsg(c *gin.Context) {
	if Check(c) == false {
		R.RJson(c, "AUTH_FAILED")
		return
	}
	//todo 拼接参数
	//todo 请求发送系统消息
	msg := Msg{}
	err := c.BindJSON(&msg)
	if err != nil {
		R.RJson(c, "INVALID_PARAM")
		return
	}

	go func() {

		url := "127.0.0.1:54218/inner/v1/system/msg?userid=9999999"
		method := "POST"
		msgBytes,_:=json.Marshal(msg)
		payload := strings.NewReader(string(msgBytes))

		client := &http.Client {
		}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))
	}()
	return
}