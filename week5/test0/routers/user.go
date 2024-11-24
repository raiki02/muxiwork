package routers

import (
	"sms/middleware"
	"sms/web"

	"github.com/gin-gonic/gin"
)

func initUser(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/user", middleware.Auth(), web.GET)
		v1.POST("/user", middleware.Auth(), web.ADD)
		v1.PUT("/user", middleware.Auth(), web.UPDATE)
		v1.DELETE("/user", middleware.Auth(), web.DELETE)
	}
	login := v1.Group("/login")
	{
		login.POST("/", web.LOGIN)
	}
	register := v1.Group("/register")
	{
		register.POST("/", web.REGISTER)
	}
	logout := v1.Group("/logout")
	{
		logout.GET("/", middleware.Auth(), web.LOGOUT)
	}
}
