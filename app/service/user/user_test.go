package user

import (
	"github.com/sirupsen/logrus"
	"testing"
	"zuoxingtao/init/common"
	"zuoxingtao/init/config"
	"zuoxingtao/init/db"
	"zuoxingtao/init/log"
)

func TestReservation(t *testing.T) {
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
	err = db.InitDB()
	if err != nil {
		t.Error(err)
		panic(err)
	}
	log.Logger = logrus.New()
	um := NewUserModel(SetLog())
	um.Reservation()
}
func TestExUser(t *testing.T) {
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
	err = db.InitDB()
	if err != nil {
		t.Error(err)
		panic(err)
	}
	log.Logger = logrus.New()
	um := NewUserModel(SetLog())
	um.ExpUser()
}
func TestAddRecord(t *testing.T) {
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
	err = db.InitDB()
	if err != nil {
		t.Error(err)
		panic(err)
	}
	log.Logger = logrus.New()
	um := NewUserModel(SetLog())
	um.AddRecord()
}
