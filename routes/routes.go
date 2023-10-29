package routes

import (
	"Basic-Enrollment-System/controller"
	"Basic-Enrollment-System/repository"
	"Basic-Enrollment-System/service"
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
	httpRouter.AddTeacherRoutes(apiRouter, mongoClient)
	httpRouter.AddStudentRoutes(apiRouter, mongoClient)
	httpRouter.AddCourseRoutes(apiRouter, mongoClient)

	apiRouter.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// Start the server
	httpRouter.router.Run(":" + os.Getenv("PORT"))
}

func (r routes) AddTeacherRoutes(rg *gin.RouterGroup, mongoClient *mongo.Client) {
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

func (r routes) AddStudentRoutes(rg *gin.RouterGroup, mongoClient *mongo.Client) {
	studentRepository := repository.NewStudentRepository(mongoClient)
	studentService := service.NewStudentService(studentRepository)
	studentController := controller.NewStudentController(studentService)
	studentRouter := rg.Group("student")

	studentRouter.GET("/", studentController.GetAllStudentData)
	studentRouter.GET("/:id", studentController.GetOneStudentData)
	studentRouter.POST("/", studentController.CreateStudentData)
	studentRouter.PATCH("/:id", studentController.EditStudentData)
	studentRouter.DELETE("/:id", studentController.DeleteStudentData)
}

func (r routes) AddCourseRoutes(rg *gin.RouterGroup, mongoClient *mongo.Client) {
	courseRepository := repository.NewCourseRepository(mongoClient)
	courseService := service.NewCourseService(courseRepository)
	courseController := controller.NewCourseController(courseService)
	courseRouter := rg.Group("course")

	courseRouter.GET("/", courseController.GetAllCourseData)
	courseRouter.GET("/:id", courseController.GetOneCourseData)
	courseRouter.POST("/", courseController.CreateCourseData)
	courseRouter.PATCH("/:id", courseController.EditCourseData)
	courseRouter.DELETE("/:id", courseController.DeleteCourseData)
}
