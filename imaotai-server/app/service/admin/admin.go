package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"imaotai_helper/constant"
	"imaotai_helper/dao"
	"imaotai_helper/dto/admin"
	requset "imaotai_helper/dto/request"
	"imaotai_helper/dto/response"
	user2 "imaotai_helper/dto/user"
	"imaotai_helper/init/common"
	"imaotai_helper/init/log"
	"imaotai_helper/lib"
	"imaotai_helper/lib/client"
	"imaotai_helper/service/user"
	"net/mail"
)

type AdminModel struct {
	*logrus.Logger
	*dao.AdminDao
	*dao.UserDao
	*user.UserModel
	*client.BaiduMapClient
}

func SetLog() func(model *AdminModel) {
	return func(model *AdminModel) {
		model.Logger = log.Logger
	}
}
func NewIndexModel(ops ...func(model *AdminModel)) *AdminModel {
	im := &AdminModel{
		AdminDao:       dao.NewAdminDao(),
		Logger:         log.Logger,
		BaiduMapClient: client.NewBaiduMapClient(),
		UserModel:      user.NewUserModel(),
		UserDao:        dao.NewUserDao(),
	}
	for _, op := range ops {
		op(im)
	}
	return im
}

func (ad *AdminModel) checkRegisterReq(req *requset.RegisterReq) error {
	if len(req.Mobile) == 0 || len(req.PassWord) == 0 || len(req.Email) == 0 || len(req.Address) == 0 {
		return errors.New("参数为空")
	}
	if !lib.IsValidChineseMobileNumber(req.Mobile) {
		return errors.New("手机号码格式错误")
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return errors.New("邮箱格式错误")
	}
	return nil
}

func (ad *AdminModel) Register(ctx *gin.Context, req *requset.RegisterReq) (*lib.Response, error) {
	resp := lib.NewResponse()
	if err := ad.checkRegisterReq(req); err != nil {
		return resp, lib.NewParamError(err)
	}
	_, err := ad.AdminDao.GetAdminByMobile(req.Mobile)
	if err == nil {
		return resp, lib.NewParamError(errors.New("该手机号码已被注册"))
	}
	if !gorm.IsRecordNotFoundError(err) {
		return resp, lib.NewDBError(err)
	}
	//加密密码
	encryptPwd, err := lib.HashPassword(req.PassWord)
	if err != nil {
		return resp, lib.NewServiceError(err)
	}
	admin := &admin.Admin{
		UID:      common.Node.Generate().String(),
		Mobile:   req.Mobile,
		PassWord: encryptPwd,
		Email:    req.Email,
		Status:   constant.ADMIN_NORMAL,
		Role:     constant.NorMalAdmin,
	}
	if err := ad.AdminDao.AddAdmin(admin); err != nil {
		return resp, lib.NewDBError(err)
	}
	// 将自己添加至user表
	user := user2.NewUser(user2.SetMobile(admin.Mobile), user2.SetUID(admin.UID), user2.SetAdminUID(admin.UID))
	//生成地址
	err = ad.SetAddress(user, req.Address)
	if err != nil {
		return resp, lib.NewServiceError(err)
	}
	if err := ad.AddUser(user); err != nil {
		return resp, lib.NewDBError(err)
	}
	//生成token
	token, err := lib.GenerateJwt(req.Mobile, user.AdminUID)
	if err != nil {
		return resp, lib.NewServiceError(err)
	}
	ctx.Header("Access-Control-Expose-Headers", constant.HEADER_JWT)
	ctx.Header(constant.HEADER_JWT, token)
	return resp, nil
}

func (ad AdminModel) checkLoginReq(req *requset.LoginReq) error {
	if len(req.Mobile) == 0 || len(req.PassWord) == 0 {
		return errors.New("参数为空")
	}
	if !lib.IsValidChineseMobileNumber(req.Mobile) {
		return errors.New("手机号码格式错误")
	}
	return nil
}

func (ad *AdminModel) Login(ctx *gin.Context, req *requset.LoginReq) (*lib.Response, error) {
	resp := lib.NewResponse()
	if err := ad.checkLoginReq(req); err != nil {
		return resp, lib.NewParamError(err)
	}
	admin, err := ad.AdminDao.GetAdminByMobile(req.Mobile)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return resp, lib.NewParamError(errors.New("该手机号码未注册"))
		}
		return resp, lib.NewDBError(err)
	}
	//校验密码
	if !lib.CheckPasswordHash(req.PassWord, admin.PassWord) {
		return resp, lib.NewServiceError(errors.New("密码错误"))
	}
	//生成token
	token, err := lib.GenerateJwt(req.Mobile, admin.UID)
	if err != nil {
		return resp, lib.NewServiceError(err)
	}
	ctx.Header("Access-Control-Expose-Headers", constant.HEADER_JWT)
	ctx.Header(constant.HEADER_JWT, token)
	return resp, nil
}

func (ad *AdminModel) GetAdminInfo(ctx *gin.Context, req *requset.GetAdminInfoReq) (*lib.Response, error) {
	resp := lib.NewResponse()
	admin, err := ad.AdminDao.GetAdminByMobile(req.Mobile)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return resp, lib.NewParamError(errors.New("该手机号码未注册"))
		}
		return resp, lib.NewDBError(err)
	}
	respAdmin := &response.GetAdminInfoResp{
		Uid:    admin.UID,
		Mobile: lib.ReplaceFourthToSeventhWithStars(admin.Mobile),
		Role:   admin.Role,
	}
	resp.Data = respAdmin
	return resp, nil
}
