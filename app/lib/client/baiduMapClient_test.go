package client

import (
	"fmt"
	"testing"
	"zuoxingtao/init/config"
)

func TestParseIp(t *testing.T) {
	config.ConfigInit("../../../config/config.toml")
	client := NewBaiduMapClient()
	//初始化配置
	resp, err := client.ParseIPToLngAndLat("")
	fmt.Printf("%#v", resp)
	if err != nil {
		t.Error(err)
	}
}
func TestParseAddress(t *testing.T) {
	config.ConfigInit("../../../config/config.toml")
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
