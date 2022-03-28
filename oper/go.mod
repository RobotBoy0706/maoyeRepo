module dg/maoyetrpg-back/oper

go 1.16

require (
	github.com/gin-contrib/pprof v1.3.0 // indirect
	github.com/gin-gonic/gin v1.7.4
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.9.0
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	miao/Utils/DB v0.0.0-00010101000000-000000000000
	miao/Utils/ErrCode v0.0.0-00010101000000-000000000000
	miao/Utils/R v0.0.0-00010101000000-000000000000
	miao/Utils/Server v0.0.0-00010101000000-000000000000
	miao/Utils/common v0.0.0-00010101000000-000000000000
)

replace miao/Utils/DB => ../miao/Utils/DB

replace miao/Utils/Server => ../miao/Utils/Server

replace miao/Utils/R => ../miao/Utils/R

//replace miao/Utils/zutil => ../miao/Utils/zutil

replace miao/Utils/ErrCode => ../miao/Utils/ErrCode

//replace miao/Utils/rand => ../miao/Utils/rand

replace miao/Utils/common => ../miao/Utils/common
