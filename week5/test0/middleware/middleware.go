package middleware

import "github.com/gin-gonic/gin"

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("my_cookie")
		if err != nil {
			c.JSON(401, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		if cookie != "1145141919810" {
			c.JSON(401, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
