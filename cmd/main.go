package main

import (
	"github.com/Kudzeri/go-shoe-store-api/config"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	router := gin.Default()
	config.ConnectDB()
	routes.SetupRoutes(router)

	router.Run(":" + port)
}
