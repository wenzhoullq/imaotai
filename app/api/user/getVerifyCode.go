package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	requset "zuoxingtao/dto/request"
	lib "zuoxingtao/lib"
	"zuoxingtao/service/user"
)

func GetVerifyCode(c *gin.Context) {
	var req requset.GetVerifyCodeReq
	if err := lib.RequestUnmarshal(c, &req); err != nil {
		lib.SetContextErrorResponse(c, err)
		return
	}
	um := user.NewUserModel(user.SetLog())
	resp, err := um.VerifyCode(c, &req)
	if err != nil {
		//打印日志
		um.Logln(logrus.ErrorLevel, err.Error())
	}
	lib.SetContextResponse(c, resp)
}
