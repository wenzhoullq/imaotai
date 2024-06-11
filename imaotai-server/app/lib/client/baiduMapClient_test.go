package client

import (
	"fmt"
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

func TestParseIp(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	client := NewBaiduMapClient()
	//初始化配置
	resp, err := client.ParseIPToLngAndLat("")
	fmt.Printf("%#v", resp)
	if err != nil {
		t.Error(err)
	}
}
func TestParseAddress(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	client := NewBaiduMapClient()
	//初始化配置
	resp, err := client.ParseAddressToLngAndLat("")
	fmt.Printf("%#v", resp)
	if err != nil {
		t.Error(err)
	}
	resp2, err := client.ParseLngAndLatToAddress(resp.Result.Location.Lng, resp.Result.Location.Lat)
	fmt.Printf("%#v", resp2)
	if err != nil {
		t.Error(err)
	}
}
