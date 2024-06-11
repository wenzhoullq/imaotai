package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	requset "imaotai_helper/dto/request"
	"imaotai_helper/lib"
	"imaotai_helper/service/admin"
)

func Register(c *gin.Context) {
	var req requset.RegisterReq
	if err := lib.RequestUnmarshal(c, &req); err != nil {
		lib.SetReqErrorResponse(c, err)
		return
	}
	im := admin.NewIndexModel(admin.SetLog())
	resp, err := im.Register(c, &req)
	im.Logln(logrus.InfoLevel, lib.ToString(req))
	if err != nil {
		//打印日志
		im.Logln(logrus.ErrorLevel, err.Error())
		lib.SetContextErrorResponse(c, resp, err)
		return
	}
	im.Logln(logrus.InfoLevel, lib.ToString(resp))
	lib.SetContextSuccessResponse(c, resp)
}
