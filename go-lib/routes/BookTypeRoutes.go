package routes

import (
	"github.com/baguette/go-lib/controllers"
	"github.com/gin-gonic/gin"
)

func TypeRoute(router *gin.Engine){
	router.GET("/types",controllers.GetTypes())
	router.POST("/types",controllers.CreateType())
	router.PUT("/types/:type_id",controllers.UpdateType())
	router.DELETE("/types/:type_id",controllers.DeleteType())
}