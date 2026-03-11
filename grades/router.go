package grades

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, service StudentService) {
	h := NewHandler(service)

	studentsRoutes := r.Group("/students")
	{
		studentsRoutes.GET("", h.GetAllStudents)
		studentsRoutes.GET("/:id", h.GetStudentByID)
		studentsRoutes.POST("/:id/grades", h.AddGrade)
	}
}
