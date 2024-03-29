package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"zuoxingtao/constant"
	"zuoxingtao/dto/response"
	"zuoxingtao/lib"
)

type ItemClient struct {
	client *resty.Client
}

func NewItemClient(ops ...func(model *ItemClient)) *ItemClient {
	ic := &ItemClient{
		client: resty.New(),
	}
	for _, op := range ops {
		op(ic)
	}
	return ic
}

func (ic *ItemClient) GetItemList() ([]*response.Item, error) {
	url := fmt.Sprintf("https://static.moutai519.com.cn/mt-backend/xhr/front/mall/index/session/get/%d", lib.GetCurrentDayTime())
	resp, err := ic.client.R().
		Get(url)
	if err != nil {
		return nil, err
	}
	sessionResp := &response.SessionResp{}
	err = json.Unmarshal(resp.Body(), sessionResp)
	if err != nil {
		return nil, err
	}
	if sessionResp.Code != constant.CODESUCCESS {
		return nil, errors.New("获取Url失败," + sessionResp.Message)
	}

	return sessionResp.Data.ItemList, nil
}
