package routers

import (
	"github.com/gin-gonic/gin"
	"t0/internal/controller"
	"t0/internal/middleware"
)

func initCommon(r *gin.Engine) {
	v1 := r.Group("v1")
	{
		v1.GET("/login", middleware.Auth, controller.Login)
		v1.POST("/register", controller.Register)
	}
}
