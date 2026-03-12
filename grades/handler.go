package grades

import (
	"strconv"

	"distributed/dto"
	"distributed/pkg/errno"
	"distributed/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service StudentService
}

func NewHandler(service StudentService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllStudents(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")

	students, total, err := h.service.GetAllStudents(page, pageSize, keyword)
	if err != nil {
		response.Error(c, errno.CodeQueryStudentsFailed, err.Error())
		return
	}

	response.OK(c, dto.StudentListResponse{
		List:     toStudentDTOList(students),
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

func (h *Handler) GetStudentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, errno.CodeInvalidStudentID, "invalid student id")
		return
	}

	student, err := h.service.GetStudentByID(id)
	if err != nil {
		response.Error(c, errno.CodeStudentNotFound, err.Error())
		return
	}

	response.OK(c, toStudentDTO(student))
}

func (h *Handler) CreateStudent(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		response.Error(c, errno.CodeInvalidStudentBody, err.Error())
		return
	}

	res, err := h.service.CreateStudent(student)
	if err != nil {
		response.Error(c, errno.CodeCreateStudentFailed, err.Error())
		return
	}

	response.OK(c, toStudentDTO(res))
}

func (h *Handler) UpdateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, errno.CodeInvalidStudentID, "invalid student id")
		return
	}

	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		response.Error(c, errno.CodeInvalidStudentBody, err.Error())
		return
	}

	res, err := h.service.UpdateStudent(id, student)
	if err != nil {
		response.Error(c, errno.CodeUpdateStudentFailed, err.Error())
		return
	}

	response.OK(c, toStudentDTO(res))
}

func (h *Handler) DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, errno.CodeInvalidStudentID, "invalid student id")
		return
	}

	if err := h.service.DeleteStudent(id); err != nil {
		response.Error(c, errno.CodeDeleteStudentFailed, err.Error())
		return
	}

	response.OK(c, gin.H{"message": "student deleted successfully"})
}

func (h *Handler) GetGradesByStudentID(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, errno.CodeInvalidStudentID, "invalid student id")
		return
	}

	grades, err := h.service.GetGradesByStudentID(studentID)
	if err != nil {
		response.Error(c, errno.CodeQueryGradesFailed, err.Error())
		return
	}

	response.OK(c, toGradeDTOList(grades))
}

func (h *Handler) AddGrade(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, errno.CodeInvalidStudentID, "invalid student id")
		return
	}

	var grade Grade
	if err := c.ShouldBindJSON(&grade); err != nil {
		response.Error(c, errno.CodeInvalidGradeBody, err.Error())
		return
	}

	if err := h.service.AddGrade(studentID, grade); err != nil {
		response.Error(c, errno.CodeAddGradeFailed, err.Error())
		return
	}

	response.OK(c, gin.H{"message": "grade added successfully"})
}

func (h *Handler) GetGradeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, errno.CodeInvalidGradeID, "invalid grade id")
		return
	}

	grade, err := h.service.GetGradeByID(id)
	if err != nil {
		response.Error(c, errno.CodeGradeNotFound, err.Error())
		return
	}

	response.OK(c, toGradeDTO(grade))
}

func (h *Handler) UpdateGrade(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, errno.CodeInvalidGradeID, "invalid grade id")
		return
	}

	var grade Grade
	if err := c.ShouldBindJSON(&grade); err != nil {
		response.Error(c, errno.CodeInvalidGradeBody, err.Error())
		return
	}

	res, err := h.service.UpdateGrade(id, grade)
	if err != nil {
		response.Error(c, errno.CodeUpdateGradeFailed, err.Error())
		return
	}

	response.OK(c, toGradeDTO(res))
}

func (h *Handler) DeleteGrade(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, errno.CodeInvalidGradeID, "invalid grade id")
		return
	}

	if err := h.service.DeleteGrade(id); err != nil {
		response.Error(c, errno.CodeDeleteGradeFailed, err.Error())
		return
	}

	response.OK(c, gin.H{"message": "grade deleted successfully"})
}
