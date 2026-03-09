package main

import (
	"distributed/pkg/config"
	"distributed/pkg/middleware"
	"distributed/registry"

	"github.com/gin-gonic/gin"
)

func main() {
	registry.SetupRegistryService()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID())
	registry.RegisterRoutes(r)

	port := config.GetEnv("PORT", "3000")

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}

//func main() {
//	registry.SetupRegistryService()
//	http.Handle("/services", &registry.RegistryService{})
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	var srv http.Server
//	srv.Addr = registry.ServerPort
//
//	go func() {
//		log.Println(srv.ListenAndServe())
//		cancel()
//	}()
//
//	go func() {
//		fmt.Println("Registry service started. Press any key to stop.")
//		var s string
//		fmt.Scanln(&s)
//		srv.Shutdown(ctx)
//		cancel()
//	}()
//
//	<-ctx.Done()
//	fmt.Println("Shutting down registry service")
//}
