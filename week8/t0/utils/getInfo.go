package utils

import (
	"github.com/gin-gonic/gin"
	"t0/model"
)

func GetBook(ctx *gin.Context) (book model.Book) {
	ctx.ShouldBindJSON(&book)
	return book
}
