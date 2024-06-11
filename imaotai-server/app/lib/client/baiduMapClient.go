package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"imaotai_helper/constant"
	"imaotai_helper/dto/response"
	"imaotai_helper/init/config"
)

type BaiduMapClient struct {
	client *resty.Client
	AK     string
}

func NewBaiduMapClient(ops ...func(model *BaiduMapClient)) *BaiduMapClient {
	bmc := &BaiduMapClient{
		AK:     config.Config.AK,
		client: resty.New(),
	}
	for _, op := range ops {
		op(bmc)
	}
	return bmc
}

func (bmc *BaiduMapClient) ParseIPToLngAndLat(ip string) (*response.BaiduParseIPResp, error) {
	url := "https://api.map.baidu.com/location/ip"
	QueryMap := map[string]string{
		"ip":   ip,
		"ak":   bmc.AK,
		"coor": "bd09ll",
	}
	resp, err := bmc.client.R().SetQueryParams(QueryMap).Get(url)
	if err != nil {
		return nil, err
	}
	bmr := &response.BaiduParseIPResp{}
	err = json.Unmarshal(resp.Body(), bmr)
	if err != nil {
		return nil, err
	}
	if bmr.Status != constant.BAIDUMAPSUCCESS {
		return nil, errors.New("ParseIPToLngAndLat fail")
	}
	return bmr, nil
}

func (bmc *BaiduMapClient) ParseAddressToLngAndLat(address string) (*response.BaiduParseAddressResp, error) {
	url := "https://api.map.baidu.com/geocoding/v3"
	QueryMap := map[string]string{
		"address": address,
		"ak":      bmc.AK,
		"output":  "json",
	}
	resp, err := bmc.client.R().SetQueryParams(QueryMap).Get(url)
	if err != nil {
		return nil, err
	}
	bpr := &response.BaiduParseAddressResp{}
	err = json.Unmarshal(resp.Body(), bpr)
	if err != nil {
		return nil, err
	}
	if bpr.Status != constant.BAIDUMAPSUCCESS {
		return nil, errors.New("ParseAddressToLngAndLat fail")
	}
	return bpr, nil
}

func (bmc *BaiduMapClient) ParseLngAndLatToAddress(lng, lat float64) (*response.BaiduParseLngAndLatsResp, error) {
	url := "https://api.map.baidu.com/reverse_geocoding/v3/"
	QueryMap := map[string]string{
		"location":  fmt.Sprintf("%f,%f", lat, lng),
		"ak":        bmc.AK,
		"output":    "json",
		"coordtype": "wgs84ll",
	}
	resp, err := bmc.client.R().SetQueryParams(QueryMap).Get(url)
	if err != nil {
		return nil, err
	}
	bpr := &response.BaiduParseLngAndLatsResp{}
	err = json.Unmarshal(resp.Body(), bpr)
	if err != nil {
		return nil, err
	}
	if bpr.Status != constant.BAIDUMAPSUCCESS {
		return nil, errors.New("ParseLngAndLatToAddress fail")
	}
	return bpr, nil
}
