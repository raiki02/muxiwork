package web

import (
	"github.com/gin-gonic/gin"
	"t0/dao"
	"t0/errors"
	"t0/model"
	"t0/utils"
)

// @Summary 获取书籍信息
// @Tags common
// @Accept json
// @Produce json
// @Success 200 {string} string "{"msg": "get book success"}"
// @Router /v1/user [get]
// @failure 401 {string} string "{"msg": " auth failed"}"
// @failure 400 {string} string "{"msg": " get book bad request"}"
func GetBook(c *gin.Context) {

	c.JSON(200, gin.H{
		"msg": "common get book success",
	})
}

// @Summary 注册
// @Tags common
// @Accept json
// @Produce json
// @Success 200 {string} string "{"msg": "common register success"}"
// @Router /v1/register [post]
// @param user_info body model.UserInfo true "user_info"
// @failure 400 {string} string "{"msg": "common register bad request"}"
func Register(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)

	db_user, _ := dao.FindUser(user.UserName)

	if user.UserName == db_user.UserName {
		c.JSON(400, gin.H{
			"msg": errors.ErrAlreadyExists,
		})
		return
	}
	
	if user.UserName == "" || user.Password == "" {
		c.JSON(400, gin.H{
			"msg": errors.ErrBlankInfo,
		})
		return
	}

	dao.Register(user.UserName, user.Password)

	name := c.PostForm("username")
	ss := utils.GenJwt(name)
	c.SetCookie("token", ss, 3600, "/", "localhost", false, true)

	c.JSON(200, gin.H{
		"msg": "common register success",
	})
}

// @Summary 登录
// @Tags common
// @Accept json
// @Produce json
// @Success 200 {string} string "{"msg": "common login success"}"
// @Router /v1/login [post]
// @param user_info body model.UserInfo true "user_info"
// @failure 400 {string} string "{"msg": "common login bad request"}"
func Login(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)

	_, err := dao.FindUser(user.UserName)
	if err != nil {
		c.JSON(401, gin.H{
			"msg": errors.ErrUnauthorized,
		})
		return
	}

	name := user.UserName
	ss := utils.GenJwt(name)
	c.SetCookie("token", ss, 3600, "/", "localhost", false, true)

	c.JSON(200, gin.H{
		"msg": "common login success",
	})
}
