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

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	studentsResp, err := h.service.GetStudents(page, pageSize, keyword)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	data := StudentsPageData{
		List:     studentsResp.List,
		Total:    studentsResp.Total,
		Page:     studentsResp.Page,
		PageSize: studentsResp.PageSize,
		Keyword:  keyword,
	}

	if data.Page > 1 {
		data.HasPrev = true
		data.PrevPage = data.Page - 1
	}

	if int64(data.Page*data.PageSize) < data.Total {
		data.HasNext = true
		data.NextPage = data.Page + 1
	}

	c.HTML(http.StatusOK, "students.html", data)
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

func (h *Handler) RenderCreateStudentPage(c *gin.Context) {
	c.HTML(http.StatusOK, "student_form.html", gin.H{
		"title":      "Create Student",
		"action":     "/students/create",
		"student":    dto.Student{},
		"submitText": "Create",
	})
}

func (h *Handler) CreateStudent(c *gin.Context) {
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")

	if firstName == "" || lastName == "" {
		c.String(http.StatusBadRequest, "first_name and last_name are required")
		return
	}

	if err := h.service.CreateStudent(firstName, lastName); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, "/students")
}

func (h *Handler) RenderEditStudentPage(c *gin.Context) {
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

	c.HTML(http.StatusOK, "student_form.html", gin.H{
		"title":      "Edit Student",
		"action":     fmt.Sprintf("/students/%d/edit", id),
		"student":    student,
		"submitText": "Update",
	})
}

func (h *Handler) UpdateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")

	if firstName == "" || lastName == "" {
		c.String(http.StatusBadRequest, "first_name and last_name are required")
		return
	}

	if err := h.service.UpdateStudent(id, firstName, lastName); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/students/%d", id))
}

func (h *Handler) DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteStudent(id); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, "/students")
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
