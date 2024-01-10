package main

import (
	"reward/pkg/config"
	"reward/pkg/middleware"
	"reward/pkg/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	//Initialize Gin Router
	router := gin.Default()

	//Initialize DB connection
	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	//Adding middleware to GinRouter
	router.Use(middleware.AuthMiddleware(db))
	//Declaring API routes
	router.GET("api/wallets/", repository.GetWallet)

	// Start server
	router.Run(":8080")
}
