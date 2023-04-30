package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/baguette/go-lib/controllers"
)

func BookRentRoutes(router *gin.Engine){
	router.GET("/RentList",controllers.GetRentList())
}