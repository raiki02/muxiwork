package web

import (
	"fmt"
	"sms/model"

	"github.com/gin-gonic/gin"
)

var User model.User

func Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"user": User,
	})
}

func Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "update user",
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}

func Register(c *gin.Context) {
	cookie, err := c.Cookie("cookie")
	if err != nil {
		cookie = "notset"
		c.SetCookie("cookie", "this_is_cookie", 3600, "/", "localhost", false, true)
	}
	c.ShouldBind(&User)
	User.ID = c.PostForm("id")
	User.Username = c.PostForm("username")
	User.Password = c.PostForm("password")
	c.JSON(200, gin.H{
		"message":  "register",
		"cookie":   cookie,
		"id":       User.ID,
		"username": User.Username,
		"password": User.Password,
	})

	fmt.Println(User)
}
