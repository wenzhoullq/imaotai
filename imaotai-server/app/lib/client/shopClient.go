package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"imaotai_helper/constant"
	"imaotai_helper/dto/response"
	"imaotai_helper/init/common"
	"imaotai_helper/lib"
)

type ShopClient struct {
	client *resty.Client
}

func NewShopClient(ops ...func(model *ShopClient)) *ShopClient {
	mc := &ShopClient{
		client: resty.New(),
	}
	for _, op := range ops {
		op(mc)
	}
	return mc
}

func (mc *ShopClient) GetShopListUrl() (string, error) {
	url := "https://static.moutai519.com.cn/mt-backend/xhr/front/mall/resource/get"
	resp, err := mc.client.R().
		Get(url)
	if err != nil {
		return "", err
	}
	shopListResp := &response.ShopListResp{}
	err = json.Unmarshal(resp.Body(), shopListResp)
	if err != nil {
		return "", err
	}
	if shopListResp.Code != constant.CODESUCCESS {
		return "", errors.New("获取Url失败," + shopListResp.Message)
	}
	return shopListResp.Data.MtshopsPc.Url, nil
}
func (mc *ShopClient) GetShopList(url string) ([]byte, error) {
	resp, err := mc.client.R().
		Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func (mc *ShopClient) GetShopsByProvince(province, itemCode string) ([]string, error) {
	url := fmt.Sprintf("https://static.moutai519.com.cn/mt-backend/xhr/front/mall/shop/list/slim/v3/%d/%s/%s/%d", common.SessionID, province, itemCode, lib.GetCurrentDayTime())
	resp, err := mc.client.R().
		Get(url)
	if err != nil {
		return nil, err
	}
	shopsResp := &response.ShopByProvinceResp{}
	err = json.Unmarshal(resp.Body(), shopsResp)
	if err != nil {
		return nil, err
	}
	if shopsResp.Code != constant.CODESUCCESS {
		return nil, errors.New("获取shopByProvince失败," + shopsResp.Message)
	}
	res := make([]string, 0)
	for _, shop := range shopsResp.Data.Shops {
		res = append(res, shop.ShopID)
	}
	return res, nil
}
