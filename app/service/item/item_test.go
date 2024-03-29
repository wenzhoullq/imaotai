package item

import (
	"testing"
	"zuoxingtao/init/common"
	"zuoxingtao/init/config"
	"zuoxingtao/init/db"
	"zuoxingtao/init/log"
)

func TestUpdateItems(t *testing.T) {
	err := common.CommonInit()
	if err != nil {
		t.Error(err)
		panic(err)
	}
	err = config.ConfigInit("../../config/config.toml")
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
	im := NewItemModel(SetLog())
	//初始化配置
	im.UpdateItems()

}
