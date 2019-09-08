package router

import (
	"github.com/gin-gonic/gin"
	"main/controller/api"
	"main/middleware"
)

func InitUserRouter(Router *gin.Engine) {
	UserRouter := Router.Group("user").Use(middleware.JWTAuth())
	{
		UserRouter.POST("changePassWord", api.ChangePassWord)
		UserRouter.POST("uploadHeaderImg", api.UploadHeaderImg)
		UserRouter.POST("getInfoList", api.GetInfoList)
	}
}
