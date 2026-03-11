package main

import (
	"distributed/grades"
	"distributed/pkg/config"
	"distributed/pkg/db"
	"distributed/pkg/middleware"

	"github.com/gin-gonic/gin"
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

	_ = mysqlDB.AutoMigrate(&grades.StudentModel{}, &grades.GradeModel{})

	repo := grades.NewMySQLStudentRepo(mysqlDB)
	service := grades.NewService(repo)
	grades.RegisterRoutes(r, service)

	var count int64
	mysqlDB.Model(&grades.StudentModel{}).Count(&count)
	if count == 0 {
		mysqlDB.Create(&grades.StudentModel{FirstName: "Tom", LastName: "Jerry"})
		mysqlDB.Create(&grades.StudentModel{FirstName: "Alice", LastName: "Smith"})
	}

	//repo := grades.NewMemoryStudentRepo()
	//service := grades.NewService(repo)
	//grades.RegisterRoutes(r, service)

	port := config.GetEnv("GRADES_PORT", "6000")

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}

//func main() {
//	host, port := "localhost", "6000"
//	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
//
//	r := registry.Registration{
//		ServiceName:      registry.GradingService,
//		ServiceURL:       serviceAddress,
//		RequiredServices: []registry.ServiceName{registry.LogService},
//		ServiceUpdateURL: serviceAddress + "/services",
//		HeartBeatURL:     serviceAddress + "/heartbeat",
//	}
//	ctx, err := service.Start(
//		context.Background(),
//		host,
//		port,
//		r,
//		grades.RegisterHandlers,
//	)
//	if err != nil {
//		stlog.Fatal(err)
//	}
//
//	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
//		fmt.Printf("Logging service found at: %s\n", logProvider)
//		log.SetClientLogger(logProvider, r.ServiceName)
//	}
//
//	<-ctx.Done()
//	fmt.Println("Shutting down grading service")
//}
