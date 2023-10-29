package controller

import (
	"Basic-Enrollment-System/model"
	"Basic-Enrollment-System/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CourseController : represent the Course's controller contract
type CourseController interface {
	GetOneCourseData(*gin.Context)
	GetAllCourseData(*gin.Context)
	CreateCourseData(*gin.Context)
	EditCourseData(*gin.Context)
	DeleteCourseData(*gin.Context)
}

type courseController struct {
	courseService service.CourseService
}

// NewCourseController -> returns new Course controller
func NewCourseController(s service.CourseService) courseController {
	return courseController{
		courseService: s,
	}
}

func (r courseController) GetAllCourseData(c *gin.Context) {
	fmt.Println("Hi")
	res, err := r.courseService.GetAllCourseData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: res},
	)
}

func (r courseController) GetOneCourseData(c *gin.Context) {
	id := c.Param("id")

	res, err := r.courseService.GetOneCourseData(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: res},
	)
}

func (r courseController) CreateCourseData(c *gin.Context) {
	var newData model.Course
	if err := c.BindJSON(&newData); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}

	res, err := r.courseService.CreateCourseData(newData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, model.Response{Status: http.StatusCreated, Message: "success", Data: res})
}

func (r courseController) EditCourseData(c *gin.Context) {
	id := c.Param("id")
	var patient model.Course
	if err := c.BindJSON(&patient); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}

	res, err := r.courseService.EditCourseData(id, &patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: res},
	)
}

func (r courseController) DeleteCourseData(c *gin.Context) {
	id := c.Param("id")

	res, err := r.courseService.DeleteCourseData(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: res},
	)
}
