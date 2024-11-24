package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dsn string = "root:lost9725ost@tcp(127.0.0.1:3306)/sms"

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
