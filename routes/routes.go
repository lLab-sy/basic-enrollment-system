package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type routes struct {
	router *gin.Engine
}

func SetupRoutes(mongoClient *mongo.Client) {
	// Create a new Gin router
	httpRouter := routes{
		router: gin.Default(),
	}

	apiRouter := httpRouter.router.Group("/api")

	apiRouter.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// Start the server
	httpRouter.router.Run(":" + os.Getenv("PORT"))
}
