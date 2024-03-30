package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
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
	for start := config.Config.TravelStart; start < config.Config.TravelEnd; start = start + config.Config.TravelStep {
		cronStr := fmt.Sprintf("0 0 %d * *", start)
		err = c.AddFunc(cronStr, TravelReward)
		if err != nil {
			return err
		}
	}
	c.Start()
	return nil
}

func UpdateShop() {
	sm := shop.NewShopModel(shop.SetLog())
	err := sm.UpdateShop()
	if err != nil {
		sm.Logln(logrus.ErrorLevel, err.Error())
	}
}
func UpdateItem() {
	im := item.NewItemModel(item.SetLog())
	err := im.UpdateItems()
	if err != nil {
		im.Logln(logrus.ErrorLevel, err.Error())
	}
}

// 过期的用户改变状态
func ExUser() {
	um := user.NewUserModel(user.SetLog())
	err := um.ExpUser()
	if err != nil {
		um.Logln(logrus.ErrorLevel, err.Error())
	}
}

func Reservation() {
	um := user.NewUserModel(user.SetLog())
	err := um.Reservation()
	if err != nil {
		um.Logln(logrus.ErrorLevel, err.Error())
	}
}

func AddRecord() {
	um := user.NewUserModel(user.SetLog())
	err := um.AddRecord()
	if err != nil {
		um.Logln(logrus.ErrorLevel, err.Error())
	}
}

func TravelReward() {
	um := user.NewUserModel(user.SetLog())
	err := um.TravelReward()
	if err != nil {
		um.Logln(logrus.ErrorLevel, err.Error())
	}
}
