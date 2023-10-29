package controller

import (
	"Basic-Enrollment-System/model"
	"Basic-Enrollment-System/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StudentController : represent the Student's controller contract
type StudentController interface {
	GetOneStudentData(*gin.Context)
	GetAllStudentData(*gin.Context)
	CreateStudentData(*gin.Context)
	EditStudentData(*gin.Context)
	DeleteStudentData(*gin.Context)
}

type studentController struct {
	studentService service.StudentService
}

// NewStudentController -> returns new Student controller
func NewStudentController(s service.StudentService) studentController {
	return studentController{
		studentService: s,
	}
}

func (r studentController) GetAllStudentData(c *gin.Context) {
	fmt.Println("Hi")
	res, err := r.studentService.GetAllStudentData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: res},
	)
}

func (r studentController) GetOneStudentData(c *gin.Context) {
	id := c.Param("id")

	res, err := r.studentService.GetOneStudentData(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: res},
	)
}

func (r studentController) CreateStudentData(c *gin.Context) {
	var newData model.Student
	if err := c.BindJSON(&newData); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}

	res, err := r.studentService.CreateStudentData(newData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, model.Response{Status: http.StatusCreated, Message: "success", Data: res})
}

func (r studentController) EditStudentData(c *gin.Context) {
	id := c.Param("id")
	var patient model.Student
	if err := c.BindJSON(&patient); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}

	res, err := r.studentService.EditStudentData(id, &patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: res},
	)
}

func (r studentController) DeleteStudentData(c *gin.Context) {
	id := c.Param("id")

	res, err := r.studentService.DeleteStudentData(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: res},
	)
}
