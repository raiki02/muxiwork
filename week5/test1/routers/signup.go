package routers

import (
	"UMS/web"

	"github.com/gin-gonic/gin"
)

func InitSignup(r *gin.Engine) {
	su := r.Group("/signup")
	su.POST("/", web.Signup)
}
