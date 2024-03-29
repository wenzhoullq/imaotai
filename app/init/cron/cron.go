package cron

import (
	"github.com/robfig/cron"
	"zuoxingtao/init/common"
	"zuoxingtao/init/config"
	"zuoxingtao/service/item"
	"zuoxingtao/service/shop"
	"zuoxingtao/service/user"
)

func InitCronTask() error {
	c := cron.New()
	err := c.AddFunc(config.Config.UpdateShop, UpdateShop)
	if err != nil {
		return err
	}
	err = c.AddFunc(config.Config.UpdateItem, UpdateItem)
	//err = c.AddFunc("*/1 * * * *", UpdateItem)
	if err != nil {
		return err
	}
	err = c.AddFunc(config.Config.ExUser, ExUser)
	if err != nil {
		return err
	}
	err = c.AddFunc(config.Config.UpdateCommon, common.UpdateCommon)
	if err != nil {
		return err
	}
	err = c.AddFunc(config.Config.Reservation, Reservation)
	if err != nil {
		return err
	}
	err = c.AddFunc(config.Config.AddRecord, AddRecord)
	if err != nil {
		return err
	}
	c.Start()
	return nil
}

func UpdateShop() {
	sm := shop.NewShopModel(shop.SetLog())
	sm.UpdateShop()
}
func UpdateItem() {
	im := item.NewItemModel(item.SetLog())
	im.UpdateItems()
}

// 过期的用户改变状态
func ExUser() {
	um := user.NewUserModel(user.SetLog())
	um.ExpUser()
}

func Reservation() {
	um := user.NewUserModel(user.SetLog())
	um.Reservation()
}

func AddRecord() {
	um := user.NewUserModel(user.SetLog())
	um.AddRecord()
}
