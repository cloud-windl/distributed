package portal

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	h := NewHandler()

	r.GET("/", h.RedirectToStudents)
	r.GET("/students", h.RenderStudents)
	r.GET("/students/:id", h.RenderStudent)
	r.POST("/students/:id/grades", h.AddGrade)
}
