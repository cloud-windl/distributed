package main

import (
	"distributed/pkg/config"
	"distributed/pkg/middleware"
	"distributed/portal"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID())
	portal.RegisterRoutes(r)

	r.LoadHTMLGlob("portal/*.html")

	port := config.GetEnv("PORT", "4000")

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}

//func main() {
//	err := portal.InmportTemplates()
//	if err != nil {
//		stlog.Fatal(err)
//	}
//	host, port := "localhost", "5000"
//	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
//
//	r := registry.Registration{
//		ServiceName: registry.PortalService,
//		ServiceURL:  serviceAddress,
//		RequiredServices: []registry.ServiceName{
//			registry.LogService,
//			registry.GradingService,
//		},
//		ServiceUpdateURL: serviceAddress + "/services",
//		HeartBeatURL:     serviceAddress + "/heartbeat",
//	}
//
//	ctx, err := service.Start(
//		context.Background(),
//		host,
//		port,
//		r,
//		portal.RegisterHandlers,
//	)
//	if err != nil {
//		stlog.Fatal(err)
//	}
//	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
//		log.SetClientLogger(logProvider, r.ServiceName)
//	}
//	<-ctx.Done()
//	fmt.Println("Shutting down portal")
//}
