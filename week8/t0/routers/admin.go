package routers

import (
	"github.com/gin-gonic/gin"
	"t0/middleware"
	"t0/web"
)

func initAdmin(r *gin.Engine) {
	v1 := r.Group("v1")
	v1.Use(middleware.Auth)
	{
		v1.GET("/admin", web.GetBook)
		v1.POST("/admin", web.AddBook)
		v1.PUT("/admin", web.UpdateBook)
		v1.DELETE("/admin", web.DeleteBook)
	}
}
