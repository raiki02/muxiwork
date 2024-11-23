package routers

import (
	"UMS/web"

	"github.com/gin-gonic/gin"
)

func InitLogin(r *gin.Engine) {
	su := r.Group("/login")
	su.POST("/", web.Login)
}
