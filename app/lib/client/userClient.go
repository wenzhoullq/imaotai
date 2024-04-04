package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
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
		return errors.New(fmt.Sprintf("PostVerify fail,userID:%d,%s,code:%d", user.UserID, imaotaiResp.Message, imaotaiResp.Code))
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
		return errors.New(fmt.Sprintf("PostLogIn fail,userId:%d,%s,code:%d", user.UserID, mtLogInResp.Message, mtLogInResp.Code))
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
		return errors.New(fmt.Sprintf("PostReserve fail,userID:%d,%s,code:%d", user.UserID, reverseResp.Message, reverseResp.Code))
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
	recordResp := &response.RecordResp{}
	err = json.Unmarshal(resp.Body(), recordResp)
	if err != nil {
		return nil, err
	}
	if recordResp.Code != constant.CODESUCCESS {
		return nil, errors.New(fmt.Sprintf("GetAppointmentResults fail,UserID:%d,%s,code:%d", user.UserID, recordResp.Message, recordResp.Code))
	}
	return recordResp, nil
}

// 获得小茅运
func (uc *UserClient) ReceiveReward(user *user.User) error {
	url := "https://h5.moutai519.com.cn/game/xmTravel/receiveReward"
	header := map[string]string{
		"MT-token":       user.Token,
		"MT-APP-Version": common.MtVersion,
		"user-Agent":     "android;29;google;sailfish",
		"MT-Device-ID":   user.DeviceID,
		"MT-Lat":         fmt.Sprintf("%f", user.Lat),
		"MT-Lng":         fmt.Sprintf("%f", user.Lng),
	}
	cookies := make([]*http.Cookie, 2)
	cookies[0] = &http.Cookie{
		Name:  "MT-Token-Wap",
		Value: user.Cookie,
	}
	cookies[1] = &http.Cookie{
		Name:  "MT-Device-ID-Wap",
		Value: user.DeviceID,
	}
	resp, err := uc.client.R().
		SetHeaders(header).
		SetCookies(cookies).
		Post(url)
	if err != nil {
		return err
	}
	receiveResp := &response.ImaotaiResp{}
	err = json.Unmarshal(resp.Body(), receiveResp)
	if err != nil {
		return err
	}
	if receiveResp.Code != constant.CODESUCCESS {
		return errors.New(fmt.Sprintf("ReceiveReward fail,userID:%d,%s,code:%d", user.UserID, receiveResp.Message, receiveResp.Code))
	}
	return nil
}

// 分享获取耐力值
func (uc *UserClient) ShareReward(user *user.User) error {
	url := "https://h5.moutai519.com.cn/game/xmTravel/shareReward"
	header := map[string]string{
		"MT-token":       user.Token,
		"MT-APP-Version": common.MtVersion,
		"user-Agent":     "android;29;google;sailfish",
		"MT-Device-ID":   user.DeviceID,
		"MT-Lat":         fmt.Sprintf("%f", user.Lat),
		"MT-Lng":         fmt.Sprintf("%f", user.Lng),
	}
	cookies := make([]*http.Cookie, 2)
	cookies[0] = &http.Cookie{
		Name:  "MT-Token-Wap",
		Value: user.Cookie,
	}
	cookies[1] = &http.Cookie{
		Name:  "MT-Device-ID-Wap",
		Value: user.DeviceID,
	}
	resp, err := uc.client.R().
		SetHeaders(header).
		SetCookies(cookies).
		Post(url)
	if err != nil {
		return err
	}
	imaotaiResp := &response.ImaotaiResp{}
	err = json.Unmarshal(resp.Body(), imaotaiResp)
	if err != nil {
		return err
	}
	if imaotaiResp.Code != constant.CODESUCCESS {
		return errors.New(fmt.Sprintf("ShareReward fail,userID:%d,%s,code:%d", user.UserID, imaotaiResp.Message, imaotaiResp.Code))
	}
	return nil
}

// 查看可以获得的小茅运
func (uc *UserClient) GetUserIsolationPageData(user *user.User) (*response.PageDataResp, error) {
	url := "https://h5.moutai519.com.cn/game/isolationPage/getUserIsolationPageData"
	header := map[string]string{
		"MT-APP-Version": common.MtVersion,
		"user-Agent":     "android;29;google;sailfish",
		"MT-Device-ID":   user.DeviceID,
	}
	cookies := make([]*http.Cookie, 2)
	cookies[0] = &http.Cookie{
		Name:  "MT-Token-Wap",
		Value: user.Cookie,
	}
	cookies[1] = &http.Cookie{
		Name:  "MT-Device-ID-Wap",
		Value: user.DeviceID,
	}
	resp, err := uc.client.R().
		SetHeaders(header).
		SetCookies(cookies).
		Get(url)
	if err != nil {
		return nil, err
	}
	pageDataResp := &response.PageDataResp{}
	err = json.Unmarshal(resp.Body(), pageDataResp)
	if err != nil {
		return nil, err
	}
	if pageDataResp.Code != constant.CODESUCCESS {
		return nil, errors.New(fmt.Sprintf("GetUserIsolationPageData fail,userID:%d,%s,code:%d", user.UserID, pageDataResp.Message, pageDataResp.Code))
	}
	return pageDataResp, nil
}

// 获取本月剩余奖励耐力值
func (uc *UserClient) GetExchangeRateInfo(user *user.User) (float64, error) {
	url := "https://h5.moutai519.com.cn/game/synthesize/exchangeRateInfo"
	header := map[string]string{
		"MT-APP-Version": common.MtVersion,
		"user-Agent":     "android;29;google;sailfish",
		"MT-Device-ID":   user.DeviceID,
	}
	cookies := make([]*http.Cookie, 2)
	cookies[0] = &http.Cookie{
		Name:  "MT-Token-Wap",
		Value: user.Cookie,
	}
	cookies[1] = &http.Cookie{
		Name:  "MT-Device-ID-Wap",
		Value: user.DeviceID,
	}
	resp, err := uc.client.R().
		SetHeaders(header).
		SetCookies(cookies).
		Get(url)
	if err != nil {
		return 0, err
	}
	exchangeRateInfoResp := &response.ExchangeRateInfoResp{}
	err = json.Unmarshal(resp.Body(), exchangeRateInfoResp)
	if err != nil {
		return 0, err
	}
	if exchangeRateInfoResp.Code != constant.CODESUCCESS {
		return 0, errors.New(fmt.Sprintf("GetExchangeRateInfo fail,userID:%d,%s,code:%d", user.UserID, exchangeRateInfoResp.Message, exchangeRateInfoResp.Code))
	}
	return exchangeRateInfoResp.Data.CurrentPeriodCanConvertXmyNum, nil
}

// 获取本月剩余小茅运
func (uc *UserClient) GetXmTravelReward(user *user.User) (float64, error) {
	url := "https://h5.moutai519.com.cn/game/xmTravel/getXmTravelReward"
	header := map[string]string{
		"MT-APP-Version": common.MtVersion,
		"user-Agent":     "android;29;google;sailfish",
		"MT-Device-ID":   user.DeviceID,
	}
	cookies := make([]*http.Cookie, 2)
	cookies[0] = &http.Cookie{
		Name:  "MT-Token-Wap",
		Value: user.Cookie,
	}
	cookies[1] = &http.Cookie{
		Name:  "MT-Device-ID-Wap",
		Value: user.DeviceID,
	}
	resp, err := uc.client.R().
		SetHeaders(header).
		SetCookies(cookies).
		Get(url)
	if err != nil {
		return 0, err
	}
	getXmTravelRewardResp := &response.GetXmTravelRewardResp{}
	err = json.Unmarshal(resp.Body(), getXmTravelRewardResp)
	if err != nil {
		return 0, err
	}
	if getXmTravelRewardResp.Code != constant.CODESUCCESS {
		return 0, errors.New(fmt.Sprintf("GetXmTravelReward fail,userID:%d,%s,code:%d", user.UserID, getXmTravelRewardResp.Message, getXmTravelRewardResp.Code))
	}
	return getXmTravelRewardResp.Data.TravelRewardXmy, nil
}

// 获得体力奖励
func (uc *UserClient) GetEnergyAward(user *user.User) error {
	url := "https://h5.moutai519.com.cn/game/isolationPage/getUserEnergyAward"
	header := map[string]string{
		"MT-APP-Version": common.MtVersion,
		"user-Agent":     "android;29;google;sailfish",
		"MT-Device-ID":   user.DeviceID,
		"MT-Lat":         fmt.Sprintf("%f", user.Lat),
		"MT-Lng":         fmt.Sprintf("%f", user.Lng),
	}
	cookies := make([]*http.Cookie, 2)
	cookies[0] = &http.Cookie{
		Name:  "MT-Token-Wap",
		Value: user.Cookie,
	}
	cookies[1] = &http.Cookie{
		Name:  "MT-Device-ID-Wap",
		Value: user.DeviceID,
	}
	resp, err := uc.client.R().
		SetHeaders(header).
		SetCookies(cookies).
		Post(url)
	lib.WriteToTxt(resp.Body())
	if err != nil {
		return err
	}
	imaotaiResp := &response.ImaotaiResp{}
	err = json.Unmarshal(resp.Body(), imaotaiResp)
	if err != nil {
		return err
	}
	// 原代码有错,用Msg判断
	if imaotaiResp.Message != "" {
		return errors.New(fmt.Sprintf("GetEnergyAward fail,userID:%d,%s,code:%d", user.UserID, imaotaiResp.Message, imaotaiResp.Code))
	}
	return nil
}

// 小茅运旅行活动
func (uc *UserClient) StartTravel(user *user.User) error {
	url := "https://h5.moutai519.com.cn/game/xmTravel/startTravel"
	header := map[string]string{
		"MT-APP-Version": common.MtVersion,
		"user-Agent":     "android;29;google;sailfish",
		"MT-Device-ID":   user.DeviceID,
	}
	cookies := make([]*http.Cookie, 2)
	cookies[0] = &http.Cookie{
		Name:  "MT-Token-Wap",
		Value: user.Cookie,
	}
	cookies[1] = &http.Cookie{
		Name:  "MT-Device-ID-Wap",
		Value: user.DeviceID,
	}
	resp, err := uc.client.R().
		SetHeaders(header).
		SetCookies(cookies).
		Post(url)
	if err != nil {
		return err
	}
	imaotaiResp := &response.ImaotaiResp{}
	err = json.Unmarshal(resp.Body(), imaotaiResp)
	if err != nil {
		return err
	}
	if imaotaiResp.Code != constant.CODESUCCESS {
		return errors.New(fmt.Sprintf("StartTravel fail,UserID:%d,%s,code:%d", user.UserID, imaotaiResp.Message, imaotaiResp.Code))
	}
	return nil
}
