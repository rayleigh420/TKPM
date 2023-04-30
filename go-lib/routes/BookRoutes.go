package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/baguette/go-lib/controllers"
)

func BookRoutes(router *gin.Engine){
	router.GET("/books",controllers.GetBooks2())
	router.GET("/books/type",controllers.GetBooksByType())
	router.GET("/books/latest",controllers.GetLatestBooks())
	router.GET("/books/popular",controllers.GetPopularBooks())
	router.GET("/books/:book_id",controllers.GetBookByID())
	router.GET("/books/search",controllers.GetBooksByName())
	router.POST("/books",controllers.CreateBook())
	router.POST("/books/:book_id/rent",controllers.RentABook())
	router.PUT("/books/:book_id",controllers.UpdateBook())
	router.DELETE("/books/:book_id")
}