package main

import (
	"log"
	"os"
	// "os"

	"github.com/baguette/go-lib/middleware"
	"github.com/baguette/go-lib/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading.env file")
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.CORSMiddleware())
	routes.UserRoute(router)
	routes.BookRoutes(router)
	routes.BookDetailRoute(router)	
	// routes.BookBorrowedRoute(router)
	routes.BookRentRoutes(router)
	// routes.HistoryRoute(router)

	router.Run(":"+os.Getenv("PORT"))
}