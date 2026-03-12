package portal

import (
	"distributed/dto"
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
		service: NewService(),
	}
}

func (h *Handler) RedirectToStudents(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/students")
}

func (h *Handler) RenderStudents(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")

	studentsResp, err := h.service.GetStudents(page, pageSize, keyword)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// 如果你当前 students.html 模板原本就是直接 range 一个切片
	// 那这里传 studentsResp.List 最稳
	c.HTML(http.StatusOK, "students.html", studentsResp.List)
}

func (h *Handler) RenderStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusNotFound)
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
		c.Status(http.StatusNotFound)
		return
	}

	score, err := strconv.ParseFloat(c.PostForm("Score"), 32)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	g := dto.Grade{
		Title: c.PostForm("Title"),
		Type:  c.PostForm("Type"),
		Score: float32(score),
	}

	if err := h.service.AddGrade(id, g); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/students/%v", id))
}
