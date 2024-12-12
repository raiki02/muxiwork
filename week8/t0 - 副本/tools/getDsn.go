package tools

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func GetDsn() (dsn string) {
	viper.SetConfigFile("./conf/conf.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	//获得用户数据库基本信息
	port := viper.GetString("mysql.port")
	host := viper.GetString("mysql.host")
	password := viper.GetString("mysql.password")
	dbName := viper.GetString("mysql.dbName")
	user := viper.GetString("mysql.user")

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=local", user, password, host, port, dbName)
	return dsn
}
