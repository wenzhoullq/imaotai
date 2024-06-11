package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"imaotai_helper/constant"
	"imaotai_helper/dto/response"
	"imaotai_helper/lib"
	"regexp"
	"strings"
)

type CommonClient struct {
	client *resty.Client
}

func NewCommonClient(ops ...func(model *CommonClient)) *CommonClient {
	client := &CommonClient{
		client: resty.New(),
	}
	for _, op := range ops {
		op(client)
	}
	return client
}

func (cc *CommonClient) GetMTVersion() (string, error) {
	url := "https://apps.apple.com/cn/app/i%E8%8C%85%E5%8F%B0/id1600482450"
	resp, err := cc.client.R().
		Get(url)
	if err != nil {
		return "", err
	}
	pattern := regexp.MustCompile("new__latest__version\">(.*?)</p>")
	matches := pattern.FindStringSubmatch(string(resp.Body()))
	if len(matches) > 1 {
		mtVersion := matches[1]
		mtVersion = strings.ReplaceAll(mtVersion, "版本 ", "")
		return mtVersion, nil
	}

	return "", errors.New("获取版本号失败")
}

func (cc *CommonClient) GetSessionID() (int, error) {
	url := fmt.Sprintf("https://static.moutai519.com.cn/mt-backend/xhr/front/mall/index/session/get/%d", lib.GetCurrentDayTime())
	resp, err := cc.client.R().
		Get(url)
	if err != nil {
		return 0, err
	}
	sessionResp := &response.SessionResp{}
	err = json.Unmarshal(resp.Body(), sessionResp)
	if err != nil {
		return 0, err
	}
	if sessionResp.Code != constant.CODESUCCESS {
		return 0, errors.New("获取Url失败," + sessionResp.Message)
	}

	return sessionResp.Data.SessionID, nil
}
