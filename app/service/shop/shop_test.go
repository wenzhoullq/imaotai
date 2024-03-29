package shop

import (
	"github.com/sirupsen/logrus"
	"testing"
	"zuoxingtao/init/common"
	"zuoxingtao/init/config"
	"zuoxingtao/init/db"
	"zuoxingtao/init/log"
)

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
	err := common.CommonInit()
	if err != nil {
		t.Error(err)
		panic(err)
	}
	err = config.ConfigInit("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	err = log.InitLog()
	if err != nil {
		t.Error(err)
		panic(err)
	}
	err = db.InitDB()
	if err != nil {
		t.Error(err)
		panic(err)
	}
	log.Logger = logrus.New()
	sm := NewShopModel(SetLog())
	sm.UpdateShop()
}
