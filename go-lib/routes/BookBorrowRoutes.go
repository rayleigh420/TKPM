package routes

import (
	"github.com/baguette/go-lib/controllers"
	"github.com/gin-gonic/gin"
)

func BookBorrowedRoute(router *gin.Engine){
	router.GET("/borrowlist",controllers.GetBorrowList())
	router.POST("/hire/:book_rent_id",controllers.HireABook())
}