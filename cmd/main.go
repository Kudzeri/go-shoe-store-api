package main

import (
	"github.com/Kudzeri/go-shoe-store-api/config"
	"github.com/Kudzeri/go-shoe-store-api/routes"
	"log"
	"os"
)

func main() {
	config.ConnectDB()

	router := routes.SetupRouter()
	router.Static("/swagger", "../swagger-ui")
	router.StaticFile("/docs/swagger.yaml", "../docs/swagger.yaml")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
