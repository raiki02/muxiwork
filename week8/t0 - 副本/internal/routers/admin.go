package routers

import (
	"github.com/gin-gonic/gin"
	"t0/internal/controller"
	"t0/internal/middleware"
)

func initAdmin(r *gin.Engine) {
	v1 := r.Group("v1")
	v1.Use(middleware.Auth)
	{
		v1.GET("/admin", controller.GetBook)
		v1.POST("/admin", controller.AddBook)
		v1.PUT("/admin", controller.UpdateBook)
		v1.DELETE("/admin", controller.DeleteBook)
	}
}
