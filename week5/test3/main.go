package main

import (
	"sms/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.InitRouter(r)
	r.Run()
}
