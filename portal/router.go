package portal

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	h := NewHandler()

	r.GET("/", h.RedirectToStudents)

	r.GET("/students", h.RenderStudents)
	r.GET("/students/create", h.RenderCreateStudentPage)
	r.POST("/students/create", h.CreateStudent)

	r.GET("/students/:id", h.RenderStudent)
	r.GET("/students/:id/edit", h.RenderEditStudentPage)
	r.POST("/students/:id/edit", h.UpdateStudent)
	r.POST("/students/:id/delete", h.DeleteStudent)

	r.POST("/students/:id/grades", h.AddGrade)
}
