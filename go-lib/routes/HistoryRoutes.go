package routes

import (
	"github.com/baguette/go-lib/controllers"
	"github.com/gin-gonic/gin"
)

func HistoryRoute(router *gin.Engine){
	router.GET("/history",controllers.GetHistory())
	router.GET("/history/:user_id",controllers.GetHistoryByUserId())
}