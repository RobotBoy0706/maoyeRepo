package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	MysqlAccount   string
	MysqlPasswd    string
	MysqlUrl       string
	MysqlLogDebug bool
	Port           string
	MongodbAccount string
	MongodbPasswd  string
	MongodbUrl     string
	Datadir        string
	Picturedir     string
	Articledir     string
	//Mailserver string
	//Mailport int
	//Account string
	//Password string
	MailList        []Email
	adminUserid     []int
	DefaultRoomIcon string
)

type Email struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

func ConfigInit() {
	viper.AddConfigPath("./cfg")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/")

	viper.SetConfigName(strings.Replace("config.json", ".json", "", 1))
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	MysqlAccount = viper.GetString("mysql.account")
	MysqlPasswd = viper.GetString("mysql.password")
	MysqlUrl = viper.GetString("mysql.url")
	MysqlLogDebug=viper.GetBool("mysql.debug")
	if MysqlUrl == "" {
		MysqlUrl = "127.0.0.1:3306"
	}

	MongodbAccount = viper.GetString("mongodb.account")
	MongodbPasswd = viper.GetString("mongodb.password")
	MongodbUrl = viper.GetString("mongodb.url")
	if MongodbUrl == "" {
		MongodbUrl = "127.0.0.1:3306"
	}

	Port = viper.GetString("port")
	if Port == "" {
		Port = ":54300"
	}
	Datadir = viper.GetString("datadir")
	if Datadir == "" {
		Datadir = "./"
	}
	Picturedir = viper.GetString("picturedir")
	if Picturedir == "" {
		Picturedir = "./"
	}
	Articledir = viper.GetString("articledir")
	if Articledir == "" {
		Articledir = "./"
	}

	//Mailserver = viper.GetString("email.server")
	//Mailport = viper.GetInt("email.port")
	//Account = viper.GetString("email.account")
	//Password = viper.GetString("email.password")
	ml := viper.Get("email")
	vtmp, err := json.Marshal(ml)
	if err != nil {
		fmt.Println("email config error: ", err)
		os.Exit(1)
	}
	err = json.Unmarshal(vtmp, &MailList)
	if err != nil {
		fmt.Println("email config error: ", err)
		os.Exit(1)
	}
	fmt.Println(MailList[0].Server)
	fmt.Println(MailList[0].Port)
	fmt.Println(MailList[0].Account)
	fmt.Println(MailList[0].Password)

	adminUserid = viper.GetIntSlice("admin")

	DefaultRoomIcon = viper.GetString("default.room-icon")

}

//func IsAdmin(userid int64) bool {
//	for _, uid := range adminUserid {
//		if userid == int64(uid) {
//			return true
//		}
//	}
//	return false
//}
