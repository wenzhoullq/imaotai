package shop

import (
	"testing"
	"zuoxingtao/init/common"
	"zuoxingtao/init/config"
	"zuoxingtao/init/db"
	"zuoxingtao/init/log"
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

func TestGetShopList(t *testing.T) {
	sm := NewShopModel()
	//初始化配置
	url, err := sm.GetShopListUrl()
	if err != nil {
		t.Error(err)
		panic(err)
	}
	json, err := sm.GetShopList(url)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	_, err = sm.ParseShopJson(json)
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
	sm := NewShopModel(SetLog())
	sm.UpdateShop()
}
