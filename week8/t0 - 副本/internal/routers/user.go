package routers

import (
	"github.com/gin-gonic/gin"
	"t0/internal/controller"
	"t0/internal/middleware"
)

func initUser(r *gin.Engine) {
	v1 := r.Group("v1")
	v1.Use(middleware.Auth)
	{
		v1.GET("/user", controller.GetBook)
	}
}
