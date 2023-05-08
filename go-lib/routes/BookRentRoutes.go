package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/baguette/go-lib/controllers"
)

func BookRentRoutes(router *gin.Engine){
	router.GET("/rentlist",controllers.GetRentList())
	router.GET("/rentlist/:book_rent_id",controllers.GetRentListById())
	router.POST("/rent",controllers.RentABook())
	router.POST("/abort/:book_rent_id",controllers.RentABook())
}