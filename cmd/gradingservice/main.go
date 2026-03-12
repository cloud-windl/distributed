// @title Distributed Grades Service API
// @version 1.0
// @description Student and grade management service based on Gin + MySQL
// @host localhost:6001
// @BasePath /
package main

import (
	"distributed/grades"
	"distributed/pkg/config"
	"distributed/pkg/db"
	"distributed/pkg/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "distributed/docs"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID())

	mysqlDB, err := db.NewMySQL()
	if err != nil {
		panic(err)
	}

	if err := grades.AutoMigrate(mysqlDB); err != nil {
		panic(err)
	}

	if err := grades.SeedStudents(mysqlDB); err != nil {
		panic(err)
	}

	repo := grades.NewMySQLStudentRepo(mysqlDB)
	service := grades.NewService(repo)
	grades.RegisterRoutes(r, service)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := config.GetEnv("GRADES_PORT", "6001")
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
