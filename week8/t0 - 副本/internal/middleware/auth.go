package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"t0/internal/errMsg"
	"t0/tools"
)

func Auth(c *gin.Context) {

	name := c.PostForm("username")
	jwt, err := c.Cookie("jwt")

	//检测是否有jwt
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": errMsg.ErrServerInternalError,
		})
		c.Abort()
		return
	}

	//检测jwt是否匹配
	n_jwt := tools.GenJwt(name)
	if jwt != n_jwt {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": errMsg.ErrUnauthorized,
		})
		c.Abort()
		return
	}
	c.Next()

}
