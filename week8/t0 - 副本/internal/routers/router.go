package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "t0/api/docs"
)

func InitRouters(r *gin.Engine) {

	initCommon(r)
	initUser(r)
	initAdmin(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
