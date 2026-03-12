package portal

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	h := NewHandler()

	r.GET("/login", h.RenderLoginPage)
	r.POST("/login", h.Login)
	r.GET("/logout", h.Logout)

	r.GET("/", h.RedirectToStudents)

	authorized := r.Group("/")
	authorized.Use(LoginRequired())
	{
		authorized.GET("/students", h.RenderStudents)
		authorized.GET("/students/create", h.RenderCreateStudentPage)
		authorized.POST("/students/create", h.CreateStudent)

		authorized.GET("/students/:id", h.RenderStudent)
		authorized.GET("/students/:id/edit", h.RenderEditStudentPage)
		authorized.POST("/students/:id/edit", h.UpdateStudent)
		authorized.POST("/students/:id/delete", h.DeleteStudent)

		authorized.POST("/students/:id/grades", h.AddGrade)

		authorized.GET("/students/:id/grades/:grade_id/edit", h.RenderEditGradePage)
		authorized.POST("/students/:id/grades/:grade_id/edit", h.UpdateGrade)
		authorized.POST("/students/:id/grades/:grade_id/delete", h.DeleteGrade)
	}
}
