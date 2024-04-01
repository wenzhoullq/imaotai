package user

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"math"
	"time"
	"zuoxingtao/constant"
	"zuoxingtao/dao"
	"zuoxingtao/dto/item"
	"zuoxingtao/dto/record"
	requset "zuoxingtao/dto/request"
	"zuoxingtao/dto/response"
	u "zuoxingtao/dto/user"
	"zuoxingtao/init/common"
	"zuoxingtao/init/log"
	"zuoxingtao/lib"
	"zuoxingtao/lib/client"
)

type UserModel struct {
	*logrus.Logger
	*dao.UserDao
	*dao.ItemDao
	*dao.ShopDao
	*dao.RecordDao
	*client.UserClient
	*client.BaiduMapClient
	*client.ShopClient
	statusMap map[int]string
}

func SetLog() func(model *UserModel) {
	return func(model *UserModel) {
		model.Logger = log.Logger
	}
}

func NewUserModel(ops ...func(model *UserModel)) *UserModel {
	um := &UserModel{
		UserClient:     client.NewUserClient(),
		BaiduMapClient: client.NewBaiduMapClient(),
		ShopClient:     client.NewShopClient(),
		UserDao:        dao.NewUserDao(),
		ItemDao:        dao.NewItemDao(),
		ShopDao:        dao.NewShopDao(),
		RecordDao:      dao.NewRecordDao(),
	}
	statusMap := map[int]string{
		constant.USER_INIT:     constant.USER_NOTEXIST,
		constant.USER_IDLE:     constant.USER_ADDRESSNULL,
		constant.USER_NORMAL:   constant.USER_PORCESSING,
		constant.USER_ABNORMAL: constant.USER_TOKENEX,
	}
	um.statusMap = statusMap
	for _, op := range ops {
		op(um)
	}
	return um
}

func (um *UserModel) checkVerifyCodeParams(req *requset.GetVerifyCodeReq) error {
	if !lib.CheckMobileNumber(req.Mobile) {
		err := errors.New("mobile form is false,mobile is " + req.Mobile)
		return err
	}
	return nil
}

func (um *UserModel) VerifyCode(c *gin.Context, req *requset.GetVerifyCodeReq) (*lib.Response, error) {
	resp := lib.NewResponse()
	err := um.checkVerifyCodeParams(req)
	if err != nil {
		lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.ParamErr))
		return resp, err
	}
	user := u.NewUser(u.SetMobile(req.Mobile))
	_, err = um.GetUserByMobile(user.Mobile)
	if err != nil {
		//如果没查到数据,则添加User
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = um.AddUser(user)
			if err != nil {
				lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.DBErr))
				return resp, err
			}
			um.Logln(logrus.InfoLevel, "add User Success, user mobile is "+req.Mobile)
		} else {
			lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.DBErr))
			return resp, err
		}
	}
	err = um.PostVerify(user)
	if err != nil {
		lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.ClientErr))
		return resp, err
	}
	um.Logln(logrus.InfoLevel, "get VerifyCode success , and mobile is "+req.Mobile)
	lib.SetResponse(resp, lib.SetErrNo(constant.Success))
	return resp, nil
}
func (um *UserModel) checkLogInParams(req *requset.LoginRequest) error {
	if !lib.CheckMobileNumber(req.Mobile) {
		err := errors.New("mobile form is false,mobile is " + req.Mobile)
		return err
	}
	if !lib.CheckVerifyCode(req.Code) {
		err := errors.New("code form is false,code is " + req.Code)
		return err
	}
	return nil
}

func (um *UserModel) LogIn(c *gin.Context, req *requset.LoginRequest) (*lib.Response, error) {
	resp := lib.NewResponse()
	err := um.checkLogInParams(req)
	if err != nil {
		lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.ParamErr))
		return resp, err
	}
	user, err := um.GetUserByMobile(req.Mobile)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			lib.SetResponse(resp, lib.SetErrMsg("请求验证码后再登录"), lib.SetErrNo(constant.DBErr))
		} else {
			lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.DBErr))
		}
		return resp, err
	}
	err = um.PostLogIn(user, req.Code)
	if err != nil {
		lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.DBErr))
		return resp, err
	}
	// 因为还没设置地址，因此是闲置状态
	user.Status = constant.USER_IDLE

	err = um.UpdateUser(user)
	if err != nil {
		lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.DBErr))
		return resp, err
	}
	um.Logln(logrus.InfoLevel, "user LogIn success , and mobile is "+req.Mobile)
	lib.SetResponse(resp, lib.SetErrNo(constant.Success))
	return resp, nil
}

func (um *UserModel) Reservation() error {
	users, err := um.GetUsersByStatus(constant.USER_NORMAL)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	items, err := um.GetItemByStatus(constant.ITEM_OPEN)
	if err != nil {
		return err
	}
	// 过滤掉不想要的酒
	items = um.FilterItem(items)
	for _, user := range users {
		for _, item := range items {
			time.Sleep(time.Second * 3)
			shopsID, err := um.GetShopsByProvince(user.ProvinceName, item.ItemCode)
			if err != nil {
				um.Logln(logrus.ErrorLevel, err.Error())
				continue
			}
			//选择距离最近的shop
			shopID, err := um.GetMinDistanceShop(user, shopsID)
			if err != nil {
				um.Logln(logrus.ErrorLevel, err.Error())
				continue
			}
			err = um.PostReserve(user, item.ItemCode, shopID)
			if err != nil {
				um.Logln(logrus.ErrorLevel, err.Error(), "userID:", user.UserID, " reserve shopID:", shopID, " sessionID:", common.SessionID, " itemCode", item.ItemCode, " fail")
				continue
			}
		}
	}
	return nil
}
func (um *UserModel) GetMinDistanceShop(user *u.User, shopIDs []string) (string, error) {
	var shopID string
	MinDis := math.MaxFloat64
	for _, v := range shopIDs {
		shop, err := um.GetShopByID(v)
		if err != nil {
			return "", err
		}
		dis := lib.CalDis(user.Lat, user.Lng, shop.Lat, shop.Lng)
		if dis < MinDis {
			MinDis = dis
			shopID = v
		}
	}
	return shopID, nil
}
func (um *UserModel) FilterItem(items []*item.Item) []*item.Item {
	filterItems := make([]*item.Item, 0)
	for _, v := range items {
		if _, ok := common.FilterSet[v.ItemCode]; !ok {
			continue
		}
		filterItems = append(filterItems, v)
	}
	return filterItems
}

func (um *UserModel) checkUpdateUserAddressParams(req *requset.UpdateAddressRequest) error {
	if !lib.CheckMobileNumber(req.Mobile) {
		err := errors.New("mobile form is false,mobile is " + req.Mobile)
		return err
	}
	if len(req.Address) == 0 {
		err := errors.New("address is empty")
		return err
	}
	return nil
}

func (um *UserModel) UpdateUserAddress(c *gin.Context, req *requset.UpdateAddressRequest) (*lib.Response, error) {
	resp := lib.NewResponse()
	err := um.checkUpdateUserAddressParams(req)
	if err != nil {
		lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.ParamErr))
		return resp, err
	}
	user, err := um.GetUserByMobile(req.Mobile)
	if err != nil {
		lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.DBErr))
		return resp, err
	}
	err = um.setAddress(user, req.Address)
	if err != nil {
		lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.DBErr))
		return resp, err
	}
	err = um.UpdateUser(user)
	if err != nil {
		lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.DBErr))
		return resp, err
	}
	um.Logln(logrus.InfoLevel, "user update success , and mobile is "+req.Mobile)
	lib.SetResponse(resp, lib.SetErrNo(constant.Success))
	return resp, nil
}

func (um *UserModel) setAddress(user *u.User, address string) error {
	resp1, err := um.ParseAddressToLngAndLat(address)
	if err != nil {
		return err
	}
	resp2, err := um.ParseLngAndLatToAddress(resp1.Result.Location.Lng, resp1.Result.Location.Lat)
	if err != nil {
		return err
	}
	user.Lat = resp1.Result.Location.Lat
	user.Lng = resp1.Result.Location.Lng
	user.ProvinceName = resp2.Result.AddressComponent.Province
	user.CityName = resp2.Result.AddressComponent.City
	user.DistrictName = resp2.Result.AddressComponent.District
	user.Status = constant.USER_NORMAL
	return nil
}

func (um *UserModel) ExpUser() error {
	users, err := um.GetUsersByStatus(constant.USER_NORMAL)
	if err != nil {
		return err
	}
	expNum := 0
	for _, u := range users {
		token, err := lib.ParseToken(u.Token)
		//解析失败直接改变状态
		if err != nil {
			um.Logln(logrus.ErrorLevel, err.Error())
			u.Status = constant.USER_ABNORMAL
			err = um.UpdateUserStatus(u)
			if err != nil {
				um.Logln(logrus.ErrorLevel, err.Error())
			}
			expNum++
			continue
		}
		if lib.Overdue(token.Exp) {
			u.Status = constant.USER_ABNORMAL
			err = um.UpdateUserStatus(u)
			if err != nil {
				um.Logln(logrus.ErrorLevel, err.Error())
				continue
			}
			expNum++
			um.Logln(logrus.InfoLevel, "user :", u.UserID, u.UserName, "has overdue")
		}
	}
	um.Logln(logrus.InfoLevel, "overDue user num:", expNum)
	return nil
}

func (um *UserModel) checkUpdateTokenParams(req *requset.UpdateTokenRequest) error {
	if !lib.CheckMobileNumber(req.Mobile) {
		err := errors.New("mobile form is false,mobile is " + req.Mobile)
		return err
	}
	if !lib.CheckVerifyCode(req.Code) {
		err := errors.New("code form is false,code is " + req.Code)
		return err
	}
	return nil
}

func (um *UserModel) UpdateToken(c *gin.Context, req *requset.UpdateTokenRequest) (*lib.Response, error) {
	resp := lib.NewResponse()
	err := um.checkUpdateTokenParams(req)
	if err != nil {
		lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.ParamErr))
		return resp, err
	}
	user, err := um.GetUserByMobile(req.Mobile)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			lib.SetResponse(resp, lib.SetErrMsg("请求验证码后再登录"), lib.SetErrNo(constant.DBErr))
		} else {
			lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.DBErr))
		}
		return resp, err
	}
	err = um.PostLogIn(user, req.Code)
	if err != nil {
		lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.DBErr))
		return resp, err
	}
	// updateToken设置过地址了
	user.Status = constant.USER_NORMAL

	err = um.UpdateUser(user)
	if err != nil {
		lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.DBErr))
		return resp, err
	}
	um.Logln(logrus.InfoLevel, "user LogIn success , and mobile is "+req.Mobile)
	lib.SetResponse(resp, lib.SetErrNo(constant.Success))
	return resp, nil
}

func (um *UserModel) parseRecord(user *u.User, resp *response.RecordResp) []*record.Record {
	records := make([]*record.Record, 0)
	for _, v := range resp.Data.ReservationItemVOS {
		if v.Status != constant.AWARD || v.SessionType != constant.SESSION_TYPE_NORMAL_ORDER {
			continue
		}
		record := &record.Record{
			UserID:   user.UserID,
			UserName: user.UserName,
			ItemName: v.ItemName,
			ItemID:   v.ItemID,
		}
		records = append(records, record)
	}
	return records
}

func (um *UserModel) AddRecord() error {
	users, err := um.GetUsersByStatus(constant.USER_NORMAL)
	if err != nil {
		return err
	}
	for _, u := range users {
		resp, err := um.GetAppointmentRecord(u)
		if err != nil {
			um.Logln(logrus.ErrorLevel, err.Error())
			continue
		}
		records := um.parseRecord(u, resp)
		for _, vv := range records {
			fmt.Sprintf("%#v", vv)
		}
		err = um.AddRecords(records)
		if err != nil {
			um.Logln(logrus.ErrorLevel, err.Error())
			continue
		}
	}
	return nil
}
func (um *UserModel) checkUserStatusParams(req *requset.GetUserStatusReq) error {
	if !lib.CheckMobileNumber(req.Mobile) {
		err := errors.New("mobile form is false,mobile is " + req.Mobile)
		return err
	}
	return nil
}
func (um *UserModel) GetUserStatus(c *gin.Context, req *requset.GetUserStatusReq) (*lib.Response, error) {
	resp := lib.NewResponse()
	statusResp := response.GetUserStatusResp{}
	if err := um.checkUserStatusParams(req); err != nil {
		return resp, err
	}
	user, err := um.GetUserByMobile(req.Mobile)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user = u.NewUser()
			user.Status = constant.USER_INIT
		} else {
			return resp, err
		}
	}
	statusResp.Status = um.statusMap[user.Status]
	lib.SetResponse(resp, lib.SetErrNo(constant.Success), lib.SetResults(statusResp))
	um.Logln(logrus.InfoLevel, " userID :", user.UserID, "userStatus:", user.Status)
	return resp, nil
}

func (um *UserModel) TravelReward() error {
	users, err := um.GetUsersByStatus(constant.USER_NORMAL)
	if err != nil {
		return err
	}
	for _, u := range users {
		pageData, err := um.GetUserIsolationPageData(u)
		if err != nil {
			um.Logln(logrus.ErrorLevel, err)
			continue
		}
		//现存体力
		curEnergy := pageData.Data.Energy
		// 获得体力
		if pageData.Data.EnergyReward.Value > 0 {
			err = um.GetEnergyAward(u)
			if err != nil {
				um.Logln(logrus.ErrorLevel, err)
				continue
			}
			curEnergy = curEnergy + pageData.Data.EnergyReward.Value
		}
		//正在旅行中
		if pageData.Data.XmTravel.Status == constant.TRAVEL_STATUS_PROCESSING {
			um.Logln(logrus.InfoLevel, "userID:", u.UserID, " is traveling")
			continue
		}
		// 如果旅行结束了,获取小茅运和首次分享获得体力值(该体力值不计入curEnergy,避免逻辑复杂化)
		if pageData.Data.XmTravel.Status == constant.TRAVEL_STATUS_FINISH {
			err = um.ReceiveReward(u)
			if err != nil {
				um.Logln(logrus.ErrorLevel, err)
				continue
			}
			err = um.ShareReward(u)
			if err != nil {
				um.Logln(logrus.ErrorLevel, err)
				continue
			}
		}
		travelRewardXmy, err := um.GetXmTravelReward(u)
		if err != nil {
			um.Logln(logrus.ErrorLevel, err)
			continue
		}
		exchangeRateInfo, err := um.GetExchangeRateInfo(u)
		if err != nil {
			um.Logln(logrus.ErrorLevel, err)
			continue
		}
		// 本月小茅运还有余额;今日次数还有;体力值大于一次旅行的消耗量;
		if exchangeRateInfo >= travelRewardXmy && curEnergy >= constant.TRAVEL_CONSUME && pageData.Data.XmTravel.RemainChance > 0 {
			err := um.StartTravel(u)
			if err != nil {
				um.Logln(logrus.ErrorLevel, err)
				continue
			}
		}
		um.Logln(logrus.InfoLevel, fmt.Sprintf("TravelReward success,userID:%d", u.UserID))
	}
	um.Logln(logrus.InfoLevel, "get reward Success")
	return nil
}
