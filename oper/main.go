package main

import (
	"github.com/globalsign/mgo"
	"miao/Utils/Server"
	"fmt"
	"miao/Utils/DB"
	"os"
)

const (
	Version    = "V0.2.01"
	modulename = "trgps"
)

var (
	Mgodb *mgo.Database
)


func DbInit() {
	// DbUrl := "root:123@tcp(127.0.0.1:3306)/"
	var err error
	DbUrl := fmt.Sprintf("%s:%s@tcp(%s)/", MysqlAccount, MysqlPasswd, MysqlUrl)
	DB.InitM(DbUrl, modulename)

	fmt.Println(DbUrl)

	DB.Db.DB().SetMaxIdleConns(1000)
	DB.Db.DB().SetMaxOpenConns(2000)
	DB.Db.LogMode(MysqlLogDebug)

	Mgodb, err = DB.InitMongo(MongodbAccount, MongodbPasswd, MongodbUrl, modulename)
	// Mgodb, err = DB.InitMongo("", "", "127.0.0.1:27017", modulename)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}


}

func main() {
	ConfigInit()
	DbInit()
	UserInit()
	ResourceInit()
	BanInit()
	AdvertisementInit()
	r := []Server.Route{
		Server.Route{"POST", "/api/v1/user/login", PostLogin},
		Server.Route{"GET", "/api/v1/user/user", GetUser},
		Server.Route{"PUT", "/api/v1/user/user", PutUser},
		Server.Route{"POST", "/api/v1/advertisement/advertisement", PostAdvertisement},
		Server.Route{"PUT", "/api/v1/advertisement/advertisement", PutAdvertisement},
		Server.Route{"DELETE", "/api/v1/advertisement/advertisement", DeleteAdvertisement},
		Server.Route{"GET", "/api/v1/advertisement/advertisement", GetAdvertisement},
		Server.Route{"GET", "/api/v1/advertisement/advertisementPicture", GetAdvertisementPicture},

		Server.Route{"POST", "/api/v1/msg/msg", SystemMsg},
		Server.Route{"POST", "/api/v1/resource/resource", PostResource},
		Server.Route{"PUT", "/api/v1/resource/resource", PutResource},
		Server.Route{"GET", "/api/v1/resource/data", GetResourceData},
		Server.Route{"DELETE", "/api/v1/resource/resource", DeleteResource},
		
		Server.Route{"DELETE", "/api/v1/rolecard/info", DeleteRoleInfo},
		Server.Route{"POST", "/api/v1/rolecard/info", PostRoleInfo},
		Server.Route{"GET", "/api/v1/rolecard/count", GetCount},
		Server.Route{"GET", "/api/v1/rolecard/list", GetList},
		Server.Route{"PUT", "/api/v1/rolecard/update", PutUpdate},






		Server.Route{"GET", "/api/v1/room/roompass", CheckRoomPass},
		Server.Route{"PUT", "/api/v1/room/room", PutRoom},
		Server.Route{"DELETE", "/api/v1/room/room", DeleteRoom},
		Server.Route{"POST", "/api/v1/user_ban/ban_user", PostBanUser},
		
		Server.Route{"GET", "/api/v1/client_user/manager_list", ManagerUserList},
		Server.Route{"GET", "/api/v1/collection/detail", AricleDetail},
		Server.Route{"GET", "/api/v1/report/update", ReportUpdate},
		Server.Route{"GET", "/api/v1/manager_report/update", ManagerReportUpdate},
		Server.Route{"GET", "/api/v1/client_user/set_manager", SetManagerUser},
		Server.Route{"GET", "/api/v1/client_user/user", GetManagerUser},
		Server.Route{"GET", "/api/v1/client_user/remove_manager", RemoveManagerUser},
	}

	s := Server.Setting{}
	s.Logger = nil
	s.Port = Port
	// s.Domain = "127.0.0.1"
	s.Domain = "0.0.0.0"
	s.Name = "trgp"
	s.Version = Version
	s.Debug = true
	s.Routes = r
	Server.Start(s)
}