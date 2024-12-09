package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var err error

func init() {
	dsn := "root:lost9725ost@tcp(localhost:3306)/getname"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect mysql failed, err:", err)
		return
	}
	fmt.Println("mysql connectec")

}
