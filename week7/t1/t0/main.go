package main

import (
	_ "sms/database"
	"sms/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.InitRouters(r)

	r.Run(":9999")
}
