package client

import (
	"testing"
)

func TestGetItemList(t *testing.T) {
	client := NewItemClient()
	//初始化配置
	_, err := client.GetItemList()
	if err != nil {
		t.Error(err)
		panic(err)
	}

}
