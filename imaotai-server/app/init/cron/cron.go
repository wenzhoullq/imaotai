package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"imaotai_helper/init/common"
	"imaotai_helper/init/config"
	cr "imaotai_helper/service/cron"
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
	cm := cr.NewCronModel(cr.SetLog())
	err := cm.UpdateShop()
	if err != nil {
		cm.Logln(logrus.ErrorLevel, err.Error())
	}
}
func UpdateItem() {
	im := cr.NewCronModel(cr.SetLog())
	err := im.UpdateItems()
	if err != nil {
		im.Logln(logrus.ErrorLevel, err.Error())
	}
}

// 过期的用户
func ExUser() {
	cm := cr.NewCronModel(cr.SetLog())
	err := cm.ExpUser()
	if err != nil {
		cm.Logln(logrus.ErrorLevel, err.Error())
	}
}

func Reservation() {
	cm := cr.NewCronModel(cr.SetLog())
	err := cm.Reservation()
	if err != nil {
		cm.Logln(logrus.ErrorLevel, err.Error())
	}
}

func AddRecord() {
	cm := cr.NewCronModel(cr.SetLog())
	err := cm.AddRecord()
	if err != nil {
		cm.Logln(logrus.ErrorLevel, err.Error())
	}
}

func TravelReward() {
	cm := cr.NewCronModel(cr.SetLog())
	err := cm.TravelReward()
	if err != nil {
		cm.Logln(logrus.ErrorLevel, err.Error())
	}
}
