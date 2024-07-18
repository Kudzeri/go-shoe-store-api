package routes

import (
	"github.com/Kudzeri/go-shoe-store-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	productRoutes := router.Group("/products")
	{
		productRoutes.GET("/", controllers.GetProducts)
		productRoutes.GET("/:id", controllers.GetProduct)
		productRoutes.POST("/", controllers.CreateProduct)
		productRoutes.PUT("/:id", controllers.UpdateProduct)
		productRoutes.DELETE("/:id", controllers.DeleteProduct)
	}

	orderRoutes := router.Group("/orders")
	{
		orderRoutes.GET("/", controllers.GetOrders)
		orderRoutes.GET("/:id", controllers.GetOrder)
		orderRoutes.POST("/", controllers.CreateOrder)
		orderRoutes.PUT("/:id", controllers.UpdateOrder)
		orderRoutes.DELETE("/:id", controllers.DeleteOrder)
	}

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", controllers.RegisterUser)
		userRoutes.GET("/:id", controllers.GetUser)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	return router
}
