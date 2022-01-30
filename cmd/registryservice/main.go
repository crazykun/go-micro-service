package main

import (
	"context"
	"fmt"
	"go-micro-service/registry"
	"log"
	"net/http"
)

func main() {
	http.Handle("/services", &registry.RegistryService{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServicePort

	// 监听错误
	go func() {
		// 打印错误
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	// 监听主动关闭
	go func() {
		fmt.Println("Registry service is running on port: " + registry.ServicePort + ". Press any key to exit.")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Shutting down registry service")
}
