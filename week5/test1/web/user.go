package web

import "github.com/gin-gonic/gin"

type User struct {
	UserName    string `json: "username"`
	NickName    string `json: "nickname"`
	Id          int    `json: "id"`
	PhoneNumber string `json: "phonenumber"`
	Gender      int    `json: "gender"`
	Age         int    `json: "age"`
}

func GetUser(c *gin.Context) {

}

func AddUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
