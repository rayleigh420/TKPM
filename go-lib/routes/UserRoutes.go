package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/baguette/go-lib/controllers"
)

func UserRoute(router *gin.Engine){
	router.POST("/login",controllers.Login())
	router.POST("/signup",controllers.SignUp())
	router.PUT("/user/:user_id",controllers.Updateuser())
	router.DELETE("/user/:user_id",controllers.DeleteUser())
	router.GET("/token",controllers.CheckToken())
	router.GET("/admin/users",controllers.GetUsers())
	router.GET("/inphieu/:book_rent_id",controllers.RequestInPhieu())
	router.GET("/inphieumuon/:book_hire_id",controllers.RequestInPhieuMuon())
	// router.GET("/inphieumuon",controllers.InPhieuMuon())
}