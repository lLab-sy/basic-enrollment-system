package controller

import (
	"Basic-Enrollment-System/model"
	"Basic-Enrollment-System/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InfoController interface {
	GetTeacherAndStudentByFaculty(*gin.Context)
}

type infoController struct {
	infoService service.InfoService
}

// NewInfoController -> returns new Info service
func NewInfoController(r service.InfoService) infoController {
	return infoController{
		infoService: r,
	}
}

func (r infoController) GetTeacherAndStudentByFaculty(c *gin.Context) {
	facultyName := c.Param("facultyName")
	res, err := r.infoService.GetTeacherAndStudentByFaculty(facultyName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: res},
	)
}
