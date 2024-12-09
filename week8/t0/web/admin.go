package web

import "github.com/gin-gonic/gin"

// @Summary 更新书籍信息
// @Tags admin
// @Accept json
// @Produce json
// @Param book_info body model.BookInfo true "book_info"
// @Success 200 {string} string "{"msg": "admin update book success"}"
// @Router /v1/admin [put]
// @Failure 401 {string} string "{"msg": "admin auth failed"}"
// @Failure 400 {string} string "{"msg": "admin update book bad request"}"
func UpdateBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "admin update book success",
	})
}

// @Summary 新增书籍
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {string} string "{"msg": "admin add book success"}"
// @Router /v1/admin [post]
// @Param book_info body model.BookInfo true "book_info"
// @Failure 401 {string} string "{"msg": "admin auth failed"}"
// @Failure 400 {string} string "{"msg": "admin add book bad request"}"
func AddBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "admin add book success",
	})
}

// @Summary 删除书籍
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {string} string "{"msg": "admin delete book success"}"
// @Router /v1/admin [delete]
// @Param book_info body model.BookInfo true "book_info"
// @Failure 401 {string} string "{"msg": "admin auth failed"}"
// @Failure 400 {string} string "{"msg": "admin delete book bad request"}"
func DeleteBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "admin delete book success",
	})
}
