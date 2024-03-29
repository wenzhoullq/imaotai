package client

import (
	"testing"
	"zuoxingtao/dao"
	"zuoxingtao/init/common"
	"zuoxingtao/init/config"
	"zuoxingtao/init/db"
)

func TestGetCode(t *testing.T) {
	err := config.ConfigInit("../../config/config.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	err = db.InitDB()
	if err != nil {
		t.Error(err)
		panic(err)
	}
	err = common.CommonInit()
	if err != nil {
		t.Error(err)
		panic(err)
	}
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserByMobile("13868449322")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	client := NewUserClient()
	//初始化配置
	err = client.PostVerify(user)
	if err != nil {
		t.Error(err)
		panic(err)
	}

}

func TestLogIn(t *testing.T) {
	err := config.ConfigInit("../../config/config.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	err = db.InitDB()
	if err != nil {
		t.Error(err)
		panic(err)
	}
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserByMobile("13868449322")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	client := NewUserClient()
	err = client.PostLogIn(user, "611699")
	if err != nil {
		t.Error(err)
		panic(err)
	}
}
func TestGetAppointmentResults(t *testing.T) {
	err := config.ConfigInit("../../config/config.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	err = db.InitDB()
	if err != nil {
		t.Error(err)
		panic(err)
	}
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserByMobile("13868449322")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	err = common.CommonInit()
	if err != nil {
		t.Error(err)
		panic(err)
	}
	client := NewUserClient()
	//初始化配置
	_, err = client.GetAppointmentRecord(user)
	if err != nil {
		t.Error(err)
		panic(err)
	}

}
