package main

import (
	"UMS/middleware"
	"UMS/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Auth)
	routers.InitRouter(r)
	r.Run()
}
