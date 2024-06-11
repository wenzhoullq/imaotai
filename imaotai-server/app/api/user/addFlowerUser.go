package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	requset "imaotai_helper/dto/request"
	"imaotai_helper/lib"
	"imaotai_helper/service/user"
)

func AddFlowerUser(c *gin.Context) {
	var req requset.AddFlowerUserRequest
	if err := lib.RequestUnmarshal(c, &req); err != nil {
		lib.SetReqErrorResponse(c, err)
		return
	}
	um := user.NewUserModel(user.SetLog())
	resp, err := um.AddFlowerUser(c, &req)
	um.Logln(logrus.InfoLevel, lib.ToString(req))
	if err != nil {
		//打印日志
		um.Logln(logrus.ErrorLevel, err.Error())
		lib.SetContextErrorResponse(c, resp, err)
		return
	}
	um.Logln(logrus.InfoLevel, lib.ToString(resp))
	lib.SetContextSuccessResponse(c, resp)
}
