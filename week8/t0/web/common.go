package web

import "github.com/gin-gonic/gin"

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
	c.JSON(200, gin.H{
		"msg": "common login success",
	})
}
