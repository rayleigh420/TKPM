package routes

import (
	"github.com/baguette/go-lib/controllers"
	"github.com/gin-gonic/gin"
)

func HistoryRoute(router *gin.Engine){
	router.GET("/history",controllers.GetHistory())
	router.GET("/history/user/:user_id",controllers.GetHistoryByUserId())
	router.GET("/returnlist",controllers.GetReturnedBooks())
	router.GET("/returnlist/search",controllers.GetHistoryById())
}