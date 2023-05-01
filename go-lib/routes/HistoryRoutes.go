package routes

import (
	"github.com/gin-gonic/gin"
)

func HistoryRoute(router *gin.Engine){
	router.GET("/history/:book_id")
}