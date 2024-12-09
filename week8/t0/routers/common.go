package routers

import (
	"github.com/gin-gonic/gin"
	"t0/middleware"
	"t0/web"
)

func initCommon(r *gin.Engine) {
	v1 := r.Group("v1")
	{
		v1.GET("/login", middleware.Auth, web.Login)
		v1.POST("/register", web.Register)
	}
}
