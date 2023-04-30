package routes

import (
	"github.com/gin-gonic/gin"
)

func LibraryRoute(router *gin.Engine){
	router.GET("/library")
	router.GET("/library/:library_id")
	router.PUT("/library/:library_id")
}