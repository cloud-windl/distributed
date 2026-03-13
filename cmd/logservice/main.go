package main

import (
	"context"
	mylog "distributed/log"
	"distributed/pkg/config"
	"distributed/pkg/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	mylog.Run("./distributed.log")

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID())

	mylog.RegisterRoutes(r)

	port := config.GetEnv("LOG_PORT", "4000")
	addr := ":" + port

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// 先启动 HTTP 服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("log service listen error: %v", err)
		}
	}()

	time.Sleep(2 * time.Second)

	reg := mylog.BuildRegistration()
	if err := mylog.RegisterToRegistry(reg); err != nil {
		log.Printf("register log service to registry failed: %v", err)
	} else {
		log.Printf("log service registered to registry: %s", reg.ServiceURL)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down log service...")

	if err := mylog.DeregisterFromRegistry(reg.ServiceURL); err != nil {
		log.Printf("deregister log service from registry failed: %v", err)
	} else {
		log.Printf("log service deregistered from registry: %s", reg.ServiceURL)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("log service shutdown error: %v", err)
	}
}
