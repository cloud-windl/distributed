package grades

import (
	"distributed/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service StudentService
}

func NewHandler() *Handler {
	return &Handler{
		service: NewService(),
	}
}

func (h *Handler) GetAllStudents(c *gin.Context) {
	response.OK(c, h.service.GetAllStudents())
}

func (h *Handler) GetStudentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, 4001, "invalid id")
		return
	}
	student, err := h.service.GetStudentByID(id)
	if err != nil {
		response.Error(c, 4004, err.Error())
		return
	}

	response.OK(c, student)
}

func (h *Handler) AddGrade(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, 4001, "invalid id")
		return
	}
	var grade Grade
	if err = c.ShouldBindJSON(&grade); err != nil {
		response.Error(c, 4002, err.Error())
		return
	}

	err = h.service.AddGrade(id, grade)
	if err != nil {
		response.Error(c, 4003, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}
