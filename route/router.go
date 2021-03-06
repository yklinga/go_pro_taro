package route

import (
	"github.com/gin-gonic/gin"
	"go_pro_taro/api"
	"go_pro_taro/middleware"
)

func RouterCollect(r * gin.Engine) * gin.Engine  {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", api.Register)
	r.POST("/api/auth/login", api.Login)
	r.GET("/api/auth/userinfo", middleware.AuthMiddleware(), api.Userinfo)

	return r
}