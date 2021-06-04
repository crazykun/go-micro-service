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
	srv.Addr = registry.ServerPort

	go func() {
		// 打印错误
		log.Println(srv.ListenAndServe())
	}()

	<-ctx.Done()
	fmt.Println("Shutting down registry service")
}
