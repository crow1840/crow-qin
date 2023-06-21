package main

import (
	"crow-qin/src/cache"
	"crow-qin/src/config"
	"crow-qin/src/router"
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

func init() {
	err := envconfig.Process("crow", &config.MyEnvs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", &config.MyEnvs)
	//mysql初始化
	var tx config.MyConnect
	tx.ConnectMysql()
	//redis初始化
	config.RedisConnect()
	cache.CheckRides()

	////初始化日志
	//config.LoggerInit()

}

// @title crow-qin
// @version 1.0
// @termsOfService http://127.0.0.1
// @contact.name 刘炜
// @contact.url https://blog.csdn.net/xingzuo_1840
// @contact.email 40010355@qq.com
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	//fmt.Println(config.IdRSAFilePath)
	router.ServerRouter()

}
