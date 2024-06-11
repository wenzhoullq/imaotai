package user

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"imaotai_helper/constant"
	"imaotai_helper/dao"
	requset "imaotai_helper/dto/request"
	"imaotai_helper/dto/response"
	u "imaotai_helper/dto/user"
	"imaotai_helper/init/common"
	"imaotai_helper/init/log"
	"imaotai_helper/lib"
	"imaotai_helper/lib/client"
	"math"
	"time"
)

type UserModel struct {
	*logrus.Logger
	*dao.AdminDao
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
		AdminDao:       dao.NewAdminDao(),
	}
	statusMap := map[int]string{
		constant.USER_INIT:     constant.USER_NOTEXIST,
		constant.USER_NORMAL:   constant.USER_PORCESSING,
		constant.USER_ABNORMAL: constant.USER_TOKENEX,
	}
	um.statusMap = statusMap
	for _, op := range ops {
		op(um)
	}
	return um
}

func (um *UserModel) checkAddFlowerUserParams(c *gin.Context, req *requset.AddFlowerUserRequest) error {
	_, ok := c.Get("admin_uid")
	if !ok {
		return errors.New("uid is null")
	}
	if len(req.Address) == 0 || len(req.Mobile) == 0 {
		err := errors.New(" address or mobile is null")
		return err
	}
	if !lib.CheckMobileNumber(req.Mobile) {
		err := errors.New("mobile form is false,mobile is " + req.Mobile)
		return err
	}
	return nil
}

func (um *UserModel) AddFlowerUser(c *gin.Context, req *requset.AddFlowerUserRequest) (*lib.Response, error) {
	resp := lib.NewResponse()
	if err := um.checkAddFlowerUserParams(c, req); err != nil {
		lib.SetResponse(resp, lib.SetErrMsg(err.Error()), lib.SetErrNo(constant.ParamErr))
		return resp, err
	}
	//检查UID
	uidValue, _ := c.Get("admin_uid")
	uid := uidValue.(string)
	if _, err := um.GetAdminByUid(uid); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return resp, lib.NewParamError(err)
		}
		return resp, lib.NewDBError(err)
	}
	//检查用户是否存在
	_, err := um.GetUserByUidAndMobile(uid, req.Mobile)
	if err == nil {
		return resp, lib.NewParamError(errors.New("user is exist"))
	}
	//数据库出错
	if !gorm.IsRecordNotFoundError(err) {
		return resp, lib.NewDBError(err)
	}
	user := u.NewUser(u.SetUID(uid), u.SetMobile(req.Mobile))
	//设置地址
	if err = um.SetAddress(user, req.Address); err != nil {
		return resp, lib.NewServiceError(err)
	}
	//添加用户
	if err = um.AddUser(user); err != nil {
		return resp, lib.NewDBError(err)
	}
	return resp, nil
}

func (um *UserModel) checkVerifyCodeParams(req *requset.GetVerifyCodeReq) error {
	if len(req.Uid) == 0 {
		err := errors.New("Uid is empty")
		return err
	}
	return nil
}

func (um *UserModel) GetVerifyCode(c *gin.Context, req *requset.GetVerifyCodeReq) (*lib.Response, error) {
	resp := lib.NewResponse()
	err := um.checkVerifyCodeParams(req)
	if err != nil {
		return resp, lib.NewParamError(err)
	}
	user, err := um.GetUserByUid(req.Uid)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return resp, lib.NewParamError(errors.New("mobile is not exist"))
		}
		return resp, lib.NewDBError(err)
	}
	err = um.PostVerify(user.Mobile, lib.GetDeviceID())
	if err != nil {
		return resp, lib.NewRpcError(err)
	}
	return resp, nil
}
func (um *UserModel) checkActivationUserParams(req *requset.ActivationUserRequest) error {
	if !lib.CheckVerifyCode(req.Code) {
		err := errors.New("code form is false,code is " + req.Code)
		return err
	}
	return nil
}

func (um *UserModel) UpdateFollowerUserToken(c *gin.Context, req *requset.ActivationUserRequest) (*lib.Response, error) {
	resp := lib.NewResponse()
	err := um.checkActivationUserParams(req)
	if err != nil {
		return resp, lib.NewParamError(err)
	}
	user, err := um.GetUserByUid(req.Uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, lib.NewParamError(errors.New("mobile is not exist"))
		}
		return resp, lib.NewDBError(err)
	}
	err = um.PostLogIn(user, req.Code)
	if err != nil {
		return resp, lib.NewRpcError(err)
	}
	user.Status = constant.USER_NORMAL
	err = um.UpdateUser(user)
	if err != nil {
		return resp, lib.NewDBError(err)
	}
	return resp, nil
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

func (um *UserModel) checkUpdateUserAddressParams(req *requset.UpdateAddressRequest) error {
	if len(req.Address) == 0 {
		err := errors.New("address is empty")
		return err
	}
	return nil
}

func (um *UserModel) UpdateFollowerUserAddress(c *gin.Context, req *requset.UpdateAddressRequest) (*lib.Response, error) {
	resp := lib.NewResponse()
	err := um.checkUpdateUserAddressParams(req)
	if err != nil {
		return resp, lib.NewParamError(err)
	}
	user, err := um.GetUserByUid(req.Uid)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return resp, lib.NewParamError(err)
		}
		return resp, lib.NewDBError(err)
	}
	err = um.SetAddress(user, req.Address)
	if err != nil {
		return resp, lib.NewRpcError(err)
	}
	err = um.UpdateUser(user)
	if err != nil {
		return resp, lib.NewDBError(err)
	}
	return resp, nil
}

func (um *UserModel) SetAddress(user *u.User, address string) error {
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
	return resp, nil
}

func (um *UserModel) checkGetFlowerUserListParams(c *gin.Context, req *requset.GetFlowerUserListReq) error {
	_, ok := c.Get("admin_uid")
	if !ok {
		return errors.New("uid is null")
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 25
	}
	return nil
}
func (um *UserModel) GetFlowerUserList(c *gin.Context, req *requset.GetFlowerUserListReq) (*lib.Response, error) {
	resp := lib.NewResponse()
	if err := um.checkGetFlowerUserListParams(c, req); err != nil {
		return resp, err
	}
	uidValue, _ := c.Get("admin_uid")
	adminUID := uidValue.(string)
	userList, err := um.GetFlowerUserListByAdminUid(adminUID, req.Page, req.PageSize)
	if err != nil {
		return resp, lib.NewDBError(err)
	}
	userListResp := make([]*response.GetUserListResp, 0)

	for _, v := range userList {
		user := &response.GetUserListResp{
			Uid:      v.UID,
			Mobile:   lib.ReplaceFourthToSeventhWithStars(v.Mobile),
			Status:   common.StatusMap[v.Status],
			UserName: v.UserName,
			Address:  fmt.Sprintf("%s%s%s", v.ProvinceName, v.CityName, v.DistrictName),
		}
		if len(v.Token) > 0 {
			token, err := lib.ParseImaoTaiToken(v.Token)
			if err != nil {
				um.Logln(logrus.ErrorLevel, "parse token to claims error,and err is ", err)
				continue
			}
			user.ExpTime = time.Unix(token.Exp, 0).Format("2006-01-02 15:04:05")
		}

		userListResp = append(userListResp, user)
	}
	resp.Data = userListResp
	return resp, nil
}

func (um *UserModel) DeleteFollowerUser(c *gin.Context, req *requset.DeleteFollowerUserRequest) (*lib.Response, error) {
	resp := lib.NewResponse()
	err := um.DeleteUser(req.Uid)
	if err != nil {
		return resp, lib.NewDBError(err)
	}
	return resp, nil
}

func (um *UserModel) SuspendFollowerUser(c *gin.Context, req *requset.SuspendFollowerUserRequest) (*lib.Response, error) {
	resp := lib.NewResponse()
	user, err := um.GetUserByUid(req.Uid)
	if err != nil {
		return resp, lib.NewDBError(err)
	}
	token := user.Token
	if len(token) == 0 {
		return resp, lib.NewParamError(errors.New("token为空,请激活用户"))
	}
	user.Status = constant.User_SUSPENDED
	err = um.UpdateUserStatus(user)
	if err != nil {
		return resp, lib.NewDBError(err)
	}
	return resp, nil
}

func (um *UserModel) StartFollowerUser(c *gin.Context, req *requset.StartFollowerUserRequest) (*lib.Response, error) {
	resp := lib.NewResponse()
	user, err := um.GetUserByUid(req.Uid)
	if err != nil {
		return resp, lib.NewDBError(err)
	}
	token := user.Token
	if len(token) == 0 {
		return resp, lib.NewParamError(errors.New("token为空,请激活用户"))
	}
	//更新user的status
	user.Status = constant.USER_NORMAL
	err = um.UpdateUserStatus(user)
	if err != nil {
		return resp, lib.NewDBError(err)
	}
	return resp, nil
}
