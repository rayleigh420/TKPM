package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/baguette/go-lib/controllers"
)

func BookDetailRoute(router *gin.Engine){
	// router.GET("/books/book_detail/:book_id")
	router.POST("/books/book_detail/:book_id",controllers.CreateBooKDetail())
}