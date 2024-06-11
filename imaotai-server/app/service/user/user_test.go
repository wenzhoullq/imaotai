package user

import (
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
	_, err = um.GetAppointmentRecord(user)
	if err != nil {
		t.Error(err)
		panic(err)
	}
}
func TestUserEnergyAward(t *testing.T) {
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
	err = um.GetEnergyAward(user)
	if err != nil {
		t.Error(err)
		panic(err)
	}
}
