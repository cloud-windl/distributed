package grades

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, service StudentService) {
	h := NewHandler(service)

	students := r.Group("/students")
	{
		students.GET("", h.GetAllStudents)
		students.GET("/:id", h.GetStudentByID)
		students.POST("", h.CreateStudent)
		students.PUT("/:id", h.UpdateStudent)
		students.DELETE("/:id", h.DeleteStudent)

		students.GET("/:id/grades", h.GetGradesByStudentID)
		students.POST("/:id/grades", h.AddGrade)
	}

	grades := r.Group("/grades")
	{
		grades.GET("/:id", h.GetGradeByID)
		grades.PUT("/:id", h.UpdateGrade)
		grades.DELETE("/:id", h.DeleteGrade)
	}
}
