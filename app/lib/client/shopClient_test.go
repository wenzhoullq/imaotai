package client

import (
	"fmt"
	"testing"
)

func TestGetShopList(t *testing.T) {
	client := NewShopClient()
	//初始化配置
	url, err := client.GetShopListUrl()
	if err != nil {
		t.Error(err)
		panic(err)
	}
	body, err := client.GetShopList(url)
	fmt.Println(string(body))
}

func TestGetShopsByProvince(t *testing.T) {
	err := initTest("../../../config/configTest.toml")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	client := NewShopClient()
	//初始化配置
	shopsId, err := client.GetShopsByProvince("", "")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	fmt.Sprint(shopsId)
}
