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

// GetAllStudents godoc
// @Summary Get student list
// @Description Get paginated student list with optional keyword filter
// @Tags students
// @Produce json
// @Param page query int false "page number" default(1)
// @Param page_size query int false "page size" default(10)
// @Param keyword query string false "search keyword"
// @Success 200 {object} StudentListResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /students [get]
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

// GetStudentByID godoc
// @Summary Get student by ID
// @Description Get detailed student information by student ID
// @Tags students
// @Produce json
// @Param id path int true "student id"
// @Success 200 {object} StudentResponseDoc
// @Failure 400 {object} ErrorResponseDoc
// @Failure 404 {object} ErrorResponseDoc
// @Router /students/{id} [get]
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

// CreateStudent godoc
// @Summary Create student
// @Description Create a new student
// @Tags students
// @Accept json
// @Produce json
// @Param student body CreateStudentRequestDoc true "student payload"
// @Success 200 {object} StudentResponseDoc
// @Failure 400 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /students [post]
func (h *Handler) CreateStudent(c *gin.Context) {
	var req CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errno.CodeInvalidStudentBody, err.Error())
		return
	}

	student := Student{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	res, err := h.service.CreateStudent(student)
	if err != nil {
		response.Error(c, errno.CodeCreateStudentFailed, err.Error())
		return
	}

	response.OK(c, toStudentDTO(res))
}

// UpdateStudent godoc
// @Summary Update student
// @Description Update student information by student ID
// @Tags students
// @Accept json
// @Produce json
// @Param id path int true "student id"
// @Param student body UpdateStudentRequestDoc true "student payload"
// @Success 200 {object} StudentResponseDoc
// @Failure 400 {object} ErrorResponseDoc
// @Failure 404 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /students/{id} [put]
func (h *Handler) UpdateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, errno.CodeInvalidStudentID, "invalid student id")
		return
	}

	var req UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errno.CodeInvalidStudentBody, err.Error())
		return
	}

	student := Student{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	res, err := h.service.UpdateStudent(id, student)
	if err != nil {
		response.Error(c, errno.CodeUpdateStudentFailed, err.Error())
		return
	}

	response.OK(c, toStudentDTO(res))
}

// DeleteStudent godoc
// @Summary Delete student
// @Description Delete a student by student ID
// @Tags students
// @Produce json
// @Param id path int true "student id"
// @Success 200 {object} MessageResponseDoc
// @Failure 400 {object} ErrorResponseDoc
// @Failure 404 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /students/{id} [delete]
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

// GetGradesByStudentID godoc
// @Summary Get grades by student ID
// @Description Get all grades for a specific student
// @Tags grades
// @Produce json
// @Param id path int true "student id"
// @Success 200 {object} GradeListResponseDoc
// @Failure 400 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /students/{id}/grades [get]
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

// AddGrade godoc
// @Summary Add grade
// @Description Add a grade to a specific student
// @Tags grades
// @Accept json
// @Produce json
// @Param id path int true "student id"
// @Param grade body CreateGradeRequestDoc true "grade payload"
// @Success 200 {object} MessageResponseDoc
// @Failure 400 {object} ErrorResponseDoc
// @Failure 404 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /students/{id}/grades [post]
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

// GetGradeByID godoc
// @Summary Get grade by ID
// @Description Get detailed grade information by grade ID
// @Tags grades
// @Produce json
// @Param id path int true "grade id"
// @Success 200 {object} GradeResponseDoc
// @Failure 400 {object} ErrorResponseDoc
// @Failure 404 {object} ErrorResponseDoc
// @Router /grades/{id} [get]
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

// UpdateGrade godoc
// @Summary Update grade
// @Description Update grade information by grade ID
// @Tags grades
// @Accept json
// @Produce json
// @Param id path int true "grade id"
// @Param grade body UpdateGradeRequestDoc true "grade payload"
// @Success 200 {object} GradeResponseDoc
// @Failure 400 {object} ErrorResponseDoc
// @Failure 404 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /grades/{id} [put]
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

// DeleteGrade godoc
// @Summary Delete grade
// @Description Delete a grade by grade ID
// @Tags grades
// @Produce json
// @Param id path int true "grade id"
// @Success 200 {object} MessageResponseDoc
// @Failure 400 {object} ErrorResponseDoc
// @Failure 404 {object} ErrorResponseDoc
// @Failure 500 {object} ErrorResponseDoc
// @Router /grades/{id} [delete]
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
