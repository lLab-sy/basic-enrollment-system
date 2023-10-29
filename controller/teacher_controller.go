package controller

import (
	"Basic-Enrollment-System/model"
	"Basic-Enrollment-System/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TeacherController : represent the Teacher's controller contract
type TeacherController interface {
	GetOneTeacherData(*gin.Context)
	GetAllTeacherData(*gin.Context)
	CreateTeacherData(*gin.Context)
	EditTeacherData(*gin.Context)
	DeleteTeacherData(*gin.Context)
}

type teacherController struct {
	teacherService service.TeacherService
}

// NewTeacherController -> returns new Teacher controller
func NewTeacherController(s service.TeacherService) teacherController {
	return teacherController{
		teacherService: s,
	}
}

func (r teacherController) GetAllTeacherData(c *gin.Context) {
	fmt.Println("Hi")
	res, err := r.teacherService.GetAllTeacherData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: res},
	)
}

func (r teacherController) GetOneTeacherData(c *gin.Context) {
	id := c.Param("id")

	res, err := r.teacherService.GetOneTeacherData(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: res},
	)
}

func (r teacherController) CreateTeacherData(c *gin.Context) {
	var newData model.Teacher
	if err := c.BindJSON(&newData); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}

	res, err := r.teacherService.CreateTeacherData(newData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, model.Response{Status: http.StatusCreated, Message: "success", Data: res})
}

func (r teacherController) EditTeacherData(c *gin.Context) {
	id := c.Param("id")
	var patient model.Teacher
	if err := c.BindJSON(&patient); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}

	res, err := r.teacherService.EditTeacherData(id, &patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: res},
	)
}

func (r teacherController) DeleteTeacherData(c *gin.Context) {
	id := c.Param("id")

	res, err := r.teacherService.DeleteTeacherData(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: res},
	)
}
