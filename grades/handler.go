package grades

import (
	"distributed/cmd/errno"
	"distributed/pkg/response"
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
		response.Error(c, errno.CodeInvalidStudentID, "invalid id")
		return
	}
	student, err := h.service.GetStudentByID(id)
	if err != nil {
		response.Error(c, errno.CodeStudentNotFound, err.Error())
		return
	}

	response.OK(c, student)
}

func (h *Handler) AddGrade(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, errno.CodeInvalidStudentID, "invalid id")
		return
	}
	var grade Grade
	if err = c.ShouldBindJSON(&grade); err != nil {
		response.Error(c, errno.CodeInvalidGradeBoyd, err.Error())
		return
	}

	err = h.service.AddGrade(id, grade)
	if err != nil {
		response.Error(c, errno.CodeAddGradeFailed, err.Error())
		return
	}

	response.OK(c, gin.H{"message": "grade added successfully"})
}
