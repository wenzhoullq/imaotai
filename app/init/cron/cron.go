package cron

import (
	"github.com/robfig/cron"
	"zuoxingtao/init/common"
	"zuoxingtao/service/item"
	"zuoxingtao/service/shop"
	"zuoxingtao/service/user"
)

func InitCronTask() error {
	c := cron.New()
	err := c.AddFunc("0 0 8 * *", UpdateShop)
	if err != nil {
		return err
	}
	err = c.AddFunc("0 10 8 * *", UpdateItem)
	//err = c.AddFunc("*/1 * * * *", UpdateItem)
	if err != nil {
		return err
	}
	err = c.AddFunc("0 20 8 * *", ExUser)
	if err != nil {
		return err
	}
	err = c.AddFunc("0 30 8 * *", common.CommonUpdate)
	if err != nil {
		return err
	}
	err = c.AddFunc("0 1 9 * *", Reservation)
	if err != nil {
		return err
	}
	err = c.AddFunc("0 0 18 * *", AddRecord)
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
