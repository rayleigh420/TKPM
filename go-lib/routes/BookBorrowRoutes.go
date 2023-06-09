package routes

import (
	"github.com/baguette/go-lib/controllers"
	"github.com/gin-gonic/gin"
)

func BookBorrowedRoute(router *gin.Engine){
	router.GET("/borrowlist",controllers.GetBorrowList())
	// router.GET("/borrowlist/:book_id",controllers.GetBorrowListOfABook())
	router.GET("/borrowlist/search",controllers.GetBookBorrowById())
	router.POST("/hire",controllers.HireABook())
	router.POST("/direct_hire",controllers.DirectlyBorrowHandler())
	router.POST("/return/:book_hire_id",controllers.ReturnABook())
}