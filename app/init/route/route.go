package route

import (
	"github.com/gin-gonic/gin"
	"zuoxingtao/api/user"
	"zuoxingtao/init/middleware"
)

func RouteInit() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Cors(), middleware.LoggerToFile())
	imaotai := r.Group("imaotai")
	{
		logIn := imaotai.Group("user")
		{
			logIn.POST("/getVerifyCode", user.GetVerifyCode)
			logIn.POST("/login", user.LogIn)
			logIn.POST("/updateAddress", user.UpdateAddress)
			logIn.POST("/updateToken", user.UpdateToken)
			logIn.POST("/getUserStatus", user.GetUserStatus)
		}
	}
	return r
}
