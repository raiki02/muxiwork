package main

// @Title 图书管理系统-api
// @license.name Apache 2.0
// version 1.0

import (
	"t0/cmd"
)

func main() {

	cmd.InitDB()
	cmd.InitGin()

}
