package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
	"zuoxingtao/constant"
	"zuoxingtao/dto/response"
	"zuoxingtao/dto/user"
	"zuoxingtao/init/common"
	"zuoxingtao/lib"
)

type UserClient struct {
	client *resty.Client
}

func NewUserClient(ops ...func(model *UserClient)) *UserClient {
	uc := &UserClient{
		client: resty.New(),
	}
	for _, op := range ops {
		op(uc)
	}
	return uc
}

func (uc *UserClient) PostVerify(user *user.User) error {
	url := "https://app.moutai519.com.cn/xhr/front/user/register/vcode"
	body := map[string]string{
		"mobile":    user.Mobile,
		"md5":       lib.Signature(user.Mobile),
		"timestamp": fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond)),
	}

	header := map[string]string{
		"MT-token":        "",
		"MT-APP-Version":  common.MtVersion,
		"user-Agent":      "android;29;google;sailfish",
		"Accept":          "*/*",
		"MT-Request-ID":   lib.GetUuID(),
		"MT-Device-ID":    user.DeviceID,
		"MT-Network-Type": "WIFI",
		"MT-Bundle-ID":    "com.moutai.mall",
		"MT-USER-TAG":     "0",
		"MT-RS":           "1080*1794",
		"Content-Type":    "application/json; charset=UTF-8",
		"Host":            "app.moutai519.com.cn",
		"MT-user-Tag":     "0",
		"MT-Team-ID":      "",
		"Content-Length":  "93",
		"Accept-Encoding": "gzip, deflate, br",
		"Connection":      "keep-alive",
		"Accept-Language": "en-CN;q=1, zh-Hans-CN;q=0.9",
	}
	resp, err := uc.client.R().
		SetBody(body).
		SetHeaders(header).
		Post(url)
	imaotaiResp := &response.ImaotaiResp{}
	err = json.Unmarshal(resp.Body(), imaotaiResp)
	if err != nil {
		return err
	}
	if imaotaiResp.Code != constant.CODESUCCESS {
		return errors.New("获取验证码失败")
	}

	return nil
}

func (uc *UserClient) PostLogIn(user *user.User, code string) error {
	url := "https://app.moutai519.com.cn/xhr/front/user/register/login"
	body := map[string]string{
		"mobile":         user.Mobile,
		"vCode":          code,
		"ydToken":        "",
		"ydLogId":        "",
		"md5":            user.Md5,
		"timestamp":      fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond)),
		"MT-APP-Version": common.MtVersion,
	}
	header := map[string]string{
		"MT-Lat":          fmt.Sprintf("%f", user.Lat),
		"MT-Lng":          fmt.Sprintf("%f", user.Lng),
		"MT-User-Tag":     "0",
		"MT-token":        "",
		"MT-APP-Version":  common.MtVersion,
		"user-Agent":      "android;29;google;sailfish",
		"Accept":          "*/*",
		"MT-Request-ID":   lib.GetUuID(),
		"MT-Device-ID":    user.DeviceID,
		"MT-Network-Type": "WIFI",
		"MT-Bundle-ID":    "com.moutai.mall",
		"MT-USER-TAG":     "0",
		"MT-RS":           "1080*1794",
		"Content-Type":    "application/json; charset=UTF-8",
		"Host":            "app.moutai519.com.cn",
		"MT-user-Tag":     "0",
		"MT-Team-ID":      "",
		"Content-Length":  "93",
		"Accept-Encoding": "gzip, deflate, br",
		"Connection":      "keep-alive",
		"Accept-Language": "en-CN;q=1, zh-Hans-CN;q=0.9",
	}
	resp, err := uc.client.R().
		SetBody(body).
		SetHeaders(header).
		Post(url)
	mtLogInResp := &response.MTLogInResp{}
	err = json.Unmarshal(resp.Body(), mtLogInResp)
	if err != nil {
		return err
	}
	if mtLogInResp.Code != constant.CODESUCCESS {
		return errors.New("登录失败," + mtLogInResp.Message)
	}
	user.Token = mtLogInResp.Data.Token
	user.UserID = mtLogInResp.Data.UserID
	user.UserName = mtLogInResp.Data.UserName
	user.Cookie = mtLogInResp.Data.Cookie
	return nil
}

func (uc *UserClient) PostReserve(user *user.User, itemCode string, shopID string) error {
	url := "https://app.moutai519.com.cn/xhr/front/mall/reservation/add"
	itemArray := make([]map[string]interface{}, 0)
	info := map[string]interface{}{
		"count":  1,
		"itemId": itemCode,
	}
	itemArray = append(itemArray, info)
	body := map[string]interface{}{
		"itemInfoList": itemArray,
		"shopId":       shopID,
		"userId":       fmt.Sprintf("%d", user.UserID),
		"sessionId":    fmt.Sprintf("%d", common.SessionID),
	}
	actParam, err := json.Marshal(body)
	if err != nil {
		return err
	}
	body["actParam"], err = lib.AesEncrypt(string(actParam))
	if err != nil {
		return err
	}
	header := map[string]string{
		"MT-Lat":          fmt.Sprintf("%f", user.Lat),
		"MT-Lng":          fmt.Sprintf("%f", user.Lng),
		"MT-User-Tag":     "0",
		"userId":          fmt.Sprintf("%d", user.UserID),
		"MT-token":        user.Token,
		"MT-APP-Version":  common.MtVersion,
		"user-Agent":      "android;29;google;sailfish",
		"Accept":          "*/*",
		"MT-Request-ID":   lib.GetUuID(),
		"MT-Device-ID":    user.DeviceID,
		"MT-Network-Type": "WIFI",
		"MT-Bundle-ID":    "com.moutai.mall",
		"MT-USER-TAG":     "0",
		"MT-RS":           "1080*1794",
		"Content-Type":    "application/json; charset=UTF-8",
		"Host":            "app.moutai519.com.cn",
		"MT-Info":         "028e7f96f6369cafe1d105579c5b9377",
		"MT-user-Tag":     "0",
		"MT-Team-ID":      "",
		"Content-Length":  "93",
		"Accept-Encoding": "gzip, deflate, br",
		"Connection":      "keep-alive",
		"Accept-Language": "en-CN;q=1, zh-Hans-CN;q=0.9",
	}
	resp, err := uc.client.R().
		SetBody(body).
		SetHeaders(header).
		Post(url)
	if err != nil {
		return err
	}
	//fmt.Println(string(resp.Body()))
	reverseResp := &response.ReserveResp{}
	err = json.Unmarshal(resp.Body(), reverseResp)
	if err != nil {
		return err
	}
	if reverseResp.Code != constant.CODESUCCESS {
		return errors.New("申购失败," + reverseResp.Message)
	}
	return nil
}

func (uc *UserClient) GetAppointmentRecord(user *user.User) (*response.RecordResp, error) {
	url := "https://app.moutai519.com.cn/xhr/front/mall/reservation/list/pageOne/query"
	header := map[string]string{
		"MT-token":       user.Token,
		"MT-APP-Version": common.MtVersion,
		"user-Agent":     "android;29;google;sailfish",
		"MT-Device-ID":   user.DeviceID,
	}
	resp, err := uc.client.R().
		SetHeaders(header).
		Get(url)
	if err != nil {
		return nil, err
	}
	lib.WriteToTxt(resp.Body())
	recordResp := &response.RecordResp{}
	err = json.Unmarshal(resp.Body(), recordResp)
	if err != nil {
		return nil, err
	}
	if recordResp.Code != constant.CODESUCCESS {
		return nil, errors.New("GetAppointmentResults fail")
	}
	return recordResp, nil
}
