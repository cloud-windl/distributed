package main

import (
	"distributed/log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Run("./distributed.log")

	r := gin.Default()
	log.RegisterRoutes(r)

	if err := r.Run(":4000"); err != nil {
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
