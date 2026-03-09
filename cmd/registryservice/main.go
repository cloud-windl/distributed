package main

import (
	"distributed/registry"

	"github.com/gin-gonic/gin"
)

func main() {
	registry.SetupRegistryService()

	r := gin.Default()
	registry.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
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
