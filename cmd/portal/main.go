package main

import (
	"context"
	mylog "distributed/log"
	"distributed/pkg/config"
	"distributed/pkg/middleware"
	"distributed/portal"
	"distributed/registry"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID())

	portal.RegisterRoutes(r)
	r.LoadHTMLGlob("portal/*.html")

	port := config.GetEnv("PORTAL_PORT", "5000")
	addr := ":" + port

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// 先启动 HTTP 服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("portal listen error: %v", err)
		}
	}()

	// 给自己一点时间，确保 /healthz 和 /registry/updates 已经能接收请求
	time.Sleep(2 * time.Second)

	reg := portal.BuildRegistration()
	if err := portal.RegisterToRegistry(reg); err != nil {
		log.Printf("register portal to registry failed: %v", err)
	} else {
		log.Printf("portal registered to registry: %s", reg.ServiceURL)
	}

	// 注册完成后，再尝试连接 log service
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		mylog.SetClientLogger(logProvider, registry.PortalService)
		log.Println("portal connected to log service")
	} else {
		log.Printf("portal failed to connect log service: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down portal...")

	if err := portal.DeregisterFromRegistry(reg.ServiceURL); err != nil {
		log.Printf("deregister portal from registry failed: %v", err)
	} else {
		log.Printf("portal deregistered from registry: %s", reg.ServiceURL)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("portal shutdown error: %v", err)
	}
}
