package routes

import (
	"Basic-Enrollment-System/controller"
	"Basic-Enrollment-System/repository"
	"Basic-Enrollment-System/service"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type routes struct {
	router *gin.Engine
}

func SetupRoutes(mongoClient *mongo.Client) {
	// Create a new Gin router
	log.Println("check")
	httpRouter := routes{
		router: gin.Default(),
	}

	apiRouter := httpRouter.router.Group("/api")
	httpRouter.AddTeacherRoutes(apiRouter, mongoClient)

	apiRouter.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// Start the server
	httpRouter.router.Run(":" + os.Getenv("PORT"))
}

func (r routes) AddTeacherRoutes(rg *gin.RouterGroup, mongoClient *mongo.Client) {
	fmt.Println("not error na")
	teacherRepository := repository.NewTeacherRepository(mongoClient)
	teacherService := service.NewTeacherService(teacherRepository)
	teacherController := controller.NewTeacherController(teacherService)
	teacherRouter := rg.Group("teacher")

	teacherRouter.GET("/", teacherController.GetAllTeacherData)
	teacherRouter.GET("/:id", teacherController.GetOneTeacherData)
	teacherRouter.POST("/", teacherController.CreateTeacherData)
	teacherRouter.PATCH("/:id", teacherController.EditTeacherData)
	teacherRouter.DELETE("/:id", teacherController.DeleteTeacherData)
}
