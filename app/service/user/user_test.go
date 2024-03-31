package user

import (
	"github.com/sirupsen/logrus"
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

func TestReservation(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	um := NewUserModel(SetLog())
	um.Reservation()
}
func TestGetAppointmentRecord(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	um := NewUserModel(SetLog())
	user, err := um.GetUserByMobile("")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	um.GetAppointmentRecord(user)
}
func TestExUser(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	um := NewUserModel(SetLog())
	err = um.ExpUser()
	if err != nil {
		um.Logln(logrus.ErrorLevel, err.Error())
	}
}
func TestAddRecord(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	um := NewUserModel(SetLog())
	err = um.AddRecord()
	if err != nil {
		um.Logln(logrus.ErrorLevel, err.Error())
	}
}

func TestTravelReward(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	um := NewUserModel(SetLog())
	err = um.TravelReward()
	if err != nil {
		t.Error(err)
		panic(err)
	}
}
