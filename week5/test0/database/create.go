package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dsn string = "user_name:user_password@tcp(127.0.0.1:3306)/database_name"

var DB *sqlx.DB
var err error

func init() {
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect mysql failed, err: ", err)
		return
	}
	fmt.Println("connect database successfully ! ")
}
