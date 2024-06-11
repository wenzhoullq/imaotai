package cron

import (
	"github.com/sirupsen/logrus"
	"imaotai_helper/init/common"
	"imaotai_helper/init/config"
	"imaotai_helper/init/db"
	"imaotai_helper/init/log"
	"testing"
)

func initTest(confAddress string) error {
	err := config.ConfigInit(confAddress)
	if err != nil {
		return err
	}
	err = common.CommonInit()
	if err != nil {
		return err
	}
	err = log.InitLog()
	if err != nil {
		return err
	}
	err = db.InitDB()
	if err != nil {
		return err
	}
	return nil
}
func TestUpdateItems(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	cm := NewCronModel(SetLog())
	//初始化配置
	cm.UpdateItems()

}

func TestGetShopList(t *testing.T) {
	cm := NewCronModel(SetLog())
	//初始化配置
	url, err := cm.GetShopListUrl()
	if err != nil {
		t.Error(err)
		panic(err)
	}
	json, err := cm.GetShopList(url)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	_, err = cm.ParseShopJson(json)
	if err != nil {
		t.Error(err)
		panic(err)
	}
}

func TestUpdateShop(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	cm := NewCronModel(SetLog())
	cm.UpdateShop()
}
func TestReservation(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	cm := NewCronModel(SetLog())
	cm.Reservation()
}
func TestExUser(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	cm := NewCronModel(SetLog())
	err = cm.ExpUser()
	if err != nil {
		cm.Logln(logrus.ErrorLevel, err.Error())
	}
}
func TestAddRecord(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	cm := NewCronModel(SetLog())
	err = cm.AddRecord()
	if err != nil {
		cm.Logln(logrus.ErrorLevel, err.Error())
	}
}

func TestTravelReward(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	cm := NewCronModel(SetLog())
	err = cm.TravelReward()
	if err != nil {
		t.Error(err)
		panic(err)
	}
}
