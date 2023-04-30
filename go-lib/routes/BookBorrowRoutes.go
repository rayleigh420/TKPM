package routes

import (
	"github.com/baguette/go-lib/controllers"
	"github.com/gin-gonic/gin"
)

func BookBorrowedRoute(router *gin.Engine){
	router.POST("/books/:book_rent_id/hire",controllers.HireABook())
}