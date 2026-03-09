package grades

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	h := NewHandler()

	studentsRoutes := r.Group("/students")
	{
		studentsRoutes.GET("", h.GetAllStudents)
		studentsRoutes.GET("/:id", h.GetStudentByID)
		studentsRoutes.POST("/:id/grades", h.AddGrade)
	}
}
