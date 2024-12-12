package cmd

import (
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"t0/internal/logs"
	"t0/internal/routers"
	"t0/internal/service"
	"t0/tools"
)

var (
	DB  *gorm.DB
	err error
)

func InitDB() {
	dsn := tools.GetDsn()

	err := service.Init(dsn)
	if err != nil {
		log.Fatal(err)
	}
}

func InitGin() {
	logs.GinLog()

	r := gin.Default()
	routers.InitRouters(r)
	r.Run(":8080")
}
