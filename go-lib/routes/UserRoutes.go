package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/baguette/go-lib/controllers"
)

func UserRoute(router *gin.Engine){
	router.POST("/login",controllers.Login())
	router.POST("/signup",controllers.SignUp())
	router.PUT("/user/:user_id",controllers.Updateuser())
	router.GET("/token",controllers.CheckToken())
}