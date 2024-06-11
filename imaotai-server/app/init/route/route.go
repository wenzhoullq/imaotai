package route

import (
	"github.com/gin-gonic/gin"
	"imaotai_helper/api/admin"
	"imaotai_helper/api/user"
	"imaotai_helper/init/middleware"
)

func RouteInit() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Cors(), middleware.LoggerToFile())
	imaotaiApi := r.Group("imaotai")
	{
		imaotaiApi.POST("/register", admin.Register)
		imaotaiApi.POST("/login", admin.Login)
		imaotaiApi.POST("/getAdminInfo", admin.GetAdminInfo)
		//imaotai.POST("/forgetPassword", user.ForgetPassword)
		userApi := imaotaiApi.Group("user")
		userApi.Use(middleware.CheckJwt())
		{
			userApi.POST("/addFlowerUser", user.AddFlowerUser)
			userApi.POST("/getVerifyCode", user.GetVerifyCode)
			userApi.POST("/updateFollowerUserToken", user.UpdateFollowerUserToken)
			userApi.POST("/suspendFollowerUser", user.SuspendFollowerUser)
			userApi.POST("/startFollowerUser", user.StartFollowerUser)
			userApi.POST("/deleteFollowerUser", user.DeleteFollowerUser)
			userApi.POST("/updateFollowerUserAddress", user.UpdateFollowerUserAddress)
			userApi.POST("/getFlowerUserList", user.GetFlowerUserList)
		}
	}
	return r
}
