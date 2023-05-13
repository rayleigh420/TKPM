package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/baguette/go-lib/controllers"
)

func BookRoutes(router *gin.Engine){
	router.GET("/admin/books",controllers.GetBooks())
	router.GET("/books",controllers.GetBooks2())
	router.GET("/books/type",controllers.GetBooksByType())
	router.GET("/books/newest",controllers.GetNewestBooks())
	router.GET("/books/popular",controllers.GetPopularBooks())
	router.GET("/books/:book_id",controllers.GetBookByID())
	router.GET("/books/:book_id/detail",controllers.GetBookDetail())
	router.GET("/books/search",controllers.GetBooksByName())
	router.POST("/books",controllers.CreateBook())
	router.PUT("/books/:book_id",controllers.UpdateBook())
	router.DELETE("/books/:book_id",controllers.DeleteBookById())
}