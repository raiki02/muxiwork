package routers

import (
	"github.com/gin-gonic/gin"
	"t0/api"
)

func InitRouters(r *gin.Engine) {
	api.ConnectAPI(r)
	initCommon(r)
	initUser(r)
	initAdmin(r)
}
