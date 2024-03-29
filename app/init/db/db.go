package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"zuoxingtao/dto/item"
	"zuoxingtao/dto/record"
	"zuoxingtao/dto/shop"
	"zuoxingtao/dto/user"
	"zuoxingtao/init/config"
	"zuoxingtao/init/log"
)

var DB *gorm.DB

func InitDB() (err error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", config.Config.UserName, config.Config.Pw, config.Config.Host, config.Config.Port, config.Config.DbName, config.Config.TimeOut)
	DB, err = gorm.Open(config.Config.Driver, dns)
	if err != nil {
		return err
	}
	err = DB.DB().Ping()
	if err != nil {
		return err
	}
	DB.AutoMigrate(&item.Item{}, &user.User{}, &record.Record{}, &shop.Shop{})
	DB.SetLogger(log.Logger)
	DB.LogMode(true)
	return nil
}
