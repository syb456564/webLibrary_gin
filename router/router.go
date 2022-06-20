package router

import (
	"github.com/gin-gonic/gin"
	"mygin.web/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/books", apis.ListBooksApi)
	router.GET("/book", apis.GetBookApi)
	router.POST("/book", apis.AddBookApi)
	router.PUT("/book/:id/:status", apis.UpdateBookStatusApi)
	//router.PUT("/book/:id/:status", apis.ReturnBookApi)
	return router
}
