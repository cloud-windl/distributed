// @title Distributed Grades Service API
// @version 1.0
// @description Student and grade management service based on Gin + MySQL
// @host localhost:6000
// @BasePath /
package main

import (
	"context"
	"distributed/grades"
	"distributed/pkg/config"
	"distributed/pkg/db"
	"distributed/pkg/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	repo, err := grades.InitMySQLRepo(mysqlDB)
	if err != nil {
		panic(err)
	}

	service := grades.NewService(repo)
	grades.RegisterRoutes(r, service, mysqlDB)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := config.GetEnv("GRADES_PORT", "6000")
	addr := ":" + port

	reg := grades.BuildRegistration()
	if err := grades.RegisterToRegistry(reg); err != nil {
		log.Printf("register to registry failed: %v", err)
	} else {
		log.Printf("registered service to registry: %s", reg.ServiceURL)
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("grades service listen error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down grades service...")

	if err := grades.DeregisterFromRegistry(reg.ServiceURL); err != nil {
		log.Printf("deregister from registry failed: %v", err)
	} else {
		log.Printf("deregistered service from registry: %s", reg.ServiceURL)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}
}
