package routers

import (
	"sms/web"

	"github.com/gin-gonic/gin"
)

func initUser(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/", web.Get)
		user.PUT("/", web.Update)
	}

	login := user.Group("/login")
	{
		login.POST("/", web.Login)
	}

	register := user.Group("/register")
	{
		register.POST("/", web.Register)
	}

}
