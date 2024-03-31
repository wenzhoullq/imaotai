package user

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
	um.ExpUser()
}
func TestAddRecord(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	um := NewUserModel(SetLog())
	um.AddRecord()
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
