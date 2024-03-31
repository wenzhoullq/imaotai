package client

import (
	"testing"
	"zuoxingtao/dao"
)

func TestGetCode(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserByMobile("")
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
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserByMobile("")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	client := NewUserClient()
	err = client.PostLogIn(user, "")
	if err != nil {
		t.Error(err)
		panic(err)
	}
}
func TestGetAppointmentResults(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserByMobile("")
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

func TestReceiveReward(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserByMobile("")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	client := NewUserClient()
	//初始化配置
	err = client.ReceiveReward(user)
	if err != nil {
		t.Error(err)
		panic(err)
	}
}

func TestGetUserIsolationPageData(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserByMobile("")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	client := NewUserClient()
	//初始化配置
	resp, err := client.GetUserIsolationPageData(user)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	t.Log(resp)
}

func TestGetExchangeRateInfo(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserByMobile("")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	client := NewUserClient()
	//初始化配置
	res, err := client.GetExchangeRateInfo(user)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	t.Log(res)
}
func TestGetGetXmTravelReward(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserByMobile("")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	client := NewUserClient()
	//初始化配置
	_, err = client.GetXmTravelReward(user)
	if err != nil {
		t.Error(err)
		panic(err)
	}
}

func TestGetEnergyAward(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserByMobile("")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	client := NewUserClient()
	//初始化配置
	err = client.GetEnergyAward(user)
	if err != nil {
		t.Error(err)
		panic(err)
	}
}

func TestShareReward(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserByMobile("")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	client := NewUserClient()
	//初始化配置
	err = client.ShareReward(user)
	if err != nil {
		t.Error(err)
		panic(err)
	}
}
