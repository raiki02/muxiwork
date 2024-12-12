package main

// @Title 图书管理系统-api测试
// @license.name Apache 2.0
// version 1.0

import (
	"github.com/gin-gonic/gin"
	"t0/logs"
	"t0/routers"
)

func main() {
	logs.InitGin()
	
	r := gin.Default()

	routers.InitRouters(r)

	r.Run(":8080")

}
