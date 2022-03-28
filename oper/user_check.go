package main

import (
	"fmt"
	"miao/Utils/DB"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func GetSession(c *gin.Context) string {
	// id即session的name
	id := c.Query("userid")

	// 取得相应的session value
	sess, err := c.Cookie(id)
	if err != nil {
		return ""
	}

	return sess
}

func CheckToken(id int64, token string) bool {
	t := Token{}

	err := DB.GetByFieldM(&t, "uid", id, TokenTable)
	if err != nil {
		fmt.Print("token not found", id)
		return false
	}

	//fmt.Println(t.Token, token)
	// Token是否正确
	if t.Token != token {
		//fmt.Print("token not match")
		return false
	}

	// 是否过期
	nowTick := int(time.Now().Unix())
	if nowTick > t.Expired {
		//fmt.Print("token expired")
		return false
	}

	return true
}

func CheckSession(id int64, session string) bool {
	t := Opersession{}

	err := DB.GetByFieldM(&t, "uid", id, OpersessionTable)
	if err != nil {
		fmt.Print("session not found", id)
		return false
	}

	// Session是否正确
	fmt.Println(t.Sig, session)

	if t.Sig != session {
		fmt.Print("session not match")
		return false
	}

	// 是否过期
	nowTick := int(time.Now().Unix())
	if nowTick > t.Expired {
		fmt.Print("session expired")
		return false
	}

	return true
}

func Check(c *gin.Context) bool {
	s := GetSession(c)
	id := c.Query("userid")

	uid, err := strconv.Atoi(id)
	if err != nil {
		return false
	}
	fmt.Printf("req session=%v\n\n\n",s)
	if s == "" || !CheckSession(int64(uid), s) {
		return false
	}
	return true
}

func GenerateToken() string {
	u := uuid.NewV1()
	sig := u.String()
	sig = strings.Replace(sig, "-", "", -1)
	return sig
}
