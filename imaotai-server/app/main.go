package main

import (
	"imaotai_helper/init/common"
	"imaotai_helper/init/config"
	"imaotai_helper/init/cron"
	"imaotai_helper/init/db"
	"imaotai_helper/init/log"
	"imaotai_helper/init/route"
	l "log"
	"net/http"
	"os"
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
