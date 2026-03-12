package grades

import (
	"distributed/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, service StudentService, db *gorm.DB) {
	h := NewHandler(service)
	authHandler := NewAuthHandler(db)

	r.POST("/login", authHandler.Login)

	students := r.Group("/students")
	{
		students.GET("", h.GetAllStudents)
		students.GET("/:id", h.GetStudentByID)
		students.GET("/:id/grades", h.GetGradesByStudentID)
	}

	protectedStudents := r.Group("/students")
	protectedStudents.Use(middleware.AuthRequired())
	{
		protectedStudents.POST("", h.CreateStudent)
		protectedStudents.PUT("/:id", h.UpdateStudent)
		protectedStudents.DELETE("/:id", h.DeleteStudent)
		protectedStudents.POST("/:id/grades", h.AddGrade)
	}

	grades := r.Group("/grades")
	{
		grades.GET("/:id", h.GetGradeByID)
	}

	protectedGrades := r.Group("/grades")
	protectedGrades.Use(middleware.AuthRequired())
	{
		protectedGrades.PUT("/:id", h.UpdateGrade)
		protectedGrades.DELETE("/:id", h.DeleteGrade)
	}
}
