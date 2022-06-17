package router

import (
	"github.com/gin-gonic/gin"
	"mygin.web/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/books", apis.ListBooksApi)
	router.POST("/addBook", apis.AddBookApi)
	router.PUT("/borrow/:id", apis.BorrowBookApi)
	router.PUT("/return/:id", apis.ReturnBookApi)
	return router
}
