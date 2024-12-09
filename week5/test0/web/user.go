package web

import (
	"fmt"
	"net/http"
	"sms/database"
	"sms/user"

	"github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {
	var temp_user user.User
	c.ShouldBindJSON(&temp_user)
	c.JSON(200, gin.H{
		"msg":      "welcome to user page",
		"user_id":  temp_user.User_id,
		"username": temp_user.Username,
		"password": temp_user.Password,
	})
}

func ADD(c *gin.Context) {
	c.JSON(http.StatusMovedPermanently, gin.H{
		"msg": "please use /v1/register to register",
	})

	c.Redirect(302, "/v1/register")
}

func UPDATE(c *gin.Context) {
	var temp_user user.Update_user
	c.ShouldBindJSON(&temp_user)
	c.JSON(200, gin.H{

		"msg":          "user" + temp_user.User_id + ": you want to update password",
		"username":     temp_user.Username,
		"old_password": temp_user.Old_Password,
		"new_password": temp_user.New_Password,
	})
	database.Update_password(temp_user.Username, temp_user.New_Password, temp_user.Old_Password)
}

func DELETE(c *gin.Context) {
	var temp_user user.User
	c.ShouldBindJSON(&temp_user)
	c.JSON(200, gin.H{
		"msg": "delete user success",
	})
	database.Delete_user(temp_user.User_id)
}

func LOGIN(c *gin.Context) {
	var temp_user user.User
	c.ShouldBindJSON(&temp_user)
	cookie, err := c.Cookie("my_cookie")
	if err != nil {
		cookie = "1145141919810"
		c.SetCookie("my_cookie", cookie, 3600, "/", "localhost", false, true)
	}
	database.Query(temp_user.User_id)

	c.JSON(200, gin.H{
		"msg":      "login success",
		"user_id":  temp_user.User_id,
		"username": temp_user.Username,
		"password": temp_user.Password,
	})
}

func REGISTER(c *gin.Context) {
	var temp_user user.User
	err := c.ShouldBindJSON(&temp_user)
	fmt.Println("??? err = ", err)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "register failed",
		})
		return
	}

	cookie, err := c.Cookie("my_cookie")
	if err != nil {
		cookie = "1145141919810"
		c.SetCookie("my_cookie", cookie, 3600, "/", "localhost", false, true)
	}
	c.JSON(200, gin.H{
		"msg":      "register success",
		"user_id":  temp_user.User_id,
		"username": temp_user.Username,
		"password": temp_user.Password,
	})
	database.Insert(temp_user.User_id, temp_user.Username, temp_user.Password)
}

func LOGOUT(c *gin.Context) {
	cookie, err := c.Cookie("my_cookie")
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "logout failed, check your cookie first",
		})
		return
	}
	c.SetCookie("my_cookie", cookie, -1, "/", "localhost", false, true)

	c.JSON(200, gin.H{
		"msg": "logout success",
	})
}
