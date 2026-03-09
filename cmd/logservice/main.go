package main

import (
	"distributed/log"
	"distributed/pkg/config"
	"distributed/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Run("./distributed.log")

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID())
	log.RegisterRoutes(r)

	port := config.GetEnv("PORT", "4000")

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}

//func main() {
//	log.Run("./distributed.log")
//	host, post := "localhost", "4000"
//	serviceAddress := fmt.Sprintf("http://%s:%s", host, post)
//
//	r := registry.Registration{
//		ServiceName:      registry.LogService,
//		ServiceURL:       serviceAddress,
//		RequiredServices: []registry.ServiceName{},
//		ServiceUpdateURL: serviceAddress + "/services",
//		HeartBeatURL:     serviceAddress + "/heartbeat",
//	}
//
//	ctx, err := service.Start(
//		context.Background(),
//		host,
//		post,
//		r,
//		log.RegisterHandlers,
//	)
//
//	if err != nil {
//		stlog.Fatalln(err)
//	}
//	<-ctx.Done()
//
//	fmt.Println("Shutting down log service.")
//}
