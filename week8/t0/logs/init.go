package logs

import (
	"github.com/gin-gonic/gin"
	"os"
)

var f *os.File

func InitGin() {
	gin.DisableConsoleColor()
	f, _ = os.OpenFile("logs/gin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	gin.DefaultWriter = f
}
