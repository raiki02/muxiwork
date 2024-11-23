package routers

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	initUser(r)

}
