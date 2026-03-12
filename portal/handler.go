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

func (h *Handler) RenderLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (h *Handler) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	token, err := h.service.Login(username, password)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}

	c.SetCookie("token", token, 3600*24, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/students")
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/login")
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

	token, err := c.Cookie("token")
	isLoggedIn := err == nil && token != ""

	studentsResp, err := h.service.GetStudents(page, pageSize, keyword, token)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	data := StudentsPageData{
		List:       studentsResp.List,
		Total:      studentsResp.Total,
		Page:       studentsResp.Page,
		PageSize:   studentsResp.PageSize,
		Keyword:    keyword,
		IsLoggedIn: isLoggedIn,
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

	token, err := c.Cookie("token")
	isLoggedIn := err == nil && token != ""

	student, err := h.service.GetStudentByID(id, token)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	data := StudentDetailPageData{
		Student:    student,
		IsLoggedIn: isLoggedIn,
	}

	c.HTML(http.StatusOK, "student.html", data)
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

	token, _ := c.Cookie("token")

	if err := h.service.CreateStudent(firstName, lastName, token); err != nil {
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

	token, _ := c.Cookie("token")

	student, err := h.service.GetStudentByID(id, token)
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

	token, _ := c.Cookie("token")

	if err := h.service.UpdateStudent(id, firstName, lastName, token); err != nil {
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

	token, _ := c.Cookie("token")

	if err := h.service.DeleteStudent(id, token); err != nil {
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

	token, _ := c.Cookie("token")

	if err := h.service.AddGrade(id, g, token); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/students/%v", id))
}

func (h *Handler) RenderEditGradePage(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	gradeID, err := strconv.Atoi(c.Param("grade_id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	token, _ := c.Cookie("token")

	student, err := h.service.GetStudentByID(studentID, token)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var target dto.Grade
	found := false
	for _, g := range student.Grades {
		if g.ID == gradeID {
			target = g
			found = true
			break
		}
	}
	if !found {
		c.String(http.StatusNotFound, "grade not found")
		return
	}

	c.HTML(http.StatusOK, "grade_form.html", gin.H{
		"title":      "Edit Grade",
		"action":     fmt.Sprintf("/students/%d/grades/%d/edit", studentID, gradeID),
		"studentID":  studentID,
		"grade":      target,
		"submitText": "Update",
	})
}

func (h *Handler) UpdateGrade(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	gradeID, err := strconv.Atoi(c.Param("grade_id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	score, err := strconv.ParseFloat(c.PostForm("score"), 32)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	g := dto.Grade{
		Title: c.PostForm("title"),
		Type:  c.PostForm("type"),
		Score: float32(score),
	}

	token, _ := c.Cookie("token")

	if err := h.service.UpdateGrade(gradeID, g, token); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/students/%d", studentID))
}

func (h *Handler) DeleteGrade(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	gradeID, err := strconv.Atoi(c.Param("grade_id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	token, _ := c.Cookie("token")

	if err := h.service.DeleteGrade(gradeID, token); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/students/%d", studentID))
}
