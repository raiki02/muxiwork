package routers

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	InitUser(r)
	InitLogin(r)
	InitSignup(r)
}
