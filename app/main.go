package main

import (
	l "log"
	"net/http"
	"os"
	"zuoxingtao/init/common"
	"zuoxingtao/init/config"
	"zuoxingtao/init/cron"
	"zuoxingtao/init/db"
	"zuoxingtao/init/log"
	"zuoxingtao/init/route"
)

func main() {
	var err error
	switch os.Getenv("ENV") {
	case "test":
		err = config.ConfigInit("../config/configTest.toml")
		break
	case "dev":
		err = config.ConfigInit("../config/configDev.toml")
		break
	default:
		l.Panicln("Env is wrong and env is " + os.Getenv("ENV"))
		return
	}
	if err != nil {
		panic(err)
	}
	err = log.InitLog()
	if err != nil {
		panic(err)
	}
	err = db.InitDB()
	if err != nil {
		panic(err)
	}
	srv := &http.Server{
		Addr:    config.Config.ServerAddr,
		Handler: route.RouteInit(),
	}
	err = cron.InitCronTask()
	if err != nil {
		panic(err)
	}
	err = common.CommonInit()
	if err != nil {
		panic(err)
	}
	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
