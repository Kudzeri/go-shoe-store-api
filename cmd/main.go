package main

import (
	"github.com/Kudzeri/go-shoe-store-api/config"
	"github.com/Kudzeri/go-shoe-store-api/controllers"
	"github.com/Kudzeri/go-shoe-store-api/middleware"
	"github.com/Kudzeri/go-shoe-store-api/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	config.ConnectDB()

	router := routes.SetupRouter()
	router.Static("/swagger", "./swagger-ui")
	router.StaticFile("/docs/swagger.yaml", "./docs/swagger.yaml")

	// Public route
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/login", controllers.Login)

	protected := router.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/protected", func(c *gin.Context) {
			username := c.MustGet("username").(string)
			c.JSON(200, gin.H{
				"message": "Hello " + username,
			})
		})

	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
