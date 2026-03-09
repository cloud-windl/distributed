package portal

import (
	"distributed/grades"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler() *Handler {
	return &Handler{
		service: NewService()}
}

func (h *Handler) RedirectToStudents(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/students")
}

func (h *Handler) RenderStudents(c *gin.Context) {
	students, err := h.service.GetStudents()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "students.html", students)
}

func (h *Handler) RenderStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	student, err := h.service.GetStudentByID(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "student.html", student)
}

func (h *Handler) AddGrade(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	score, err := strconv.ParseFloat(c.PostForm("Score"), 32)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	g := grades.Grade{
		Title: c.PostForm("Title"),
		Type:  grades.GradeType(c.PostForm("Type")),
		Score: float32(score),
	}
	if err := h.service.AddGrade(id, g); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/students/%v", id))
}
