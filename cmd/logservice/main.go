package main

import (
	"context"
	"fmt"
	"go-micro-service/log"
	"go-micro-service/registry"
	"go-micro-service/service"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName: "Log Service",
		ServiceUrl:  serviceAddress,
	}

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandler,
	)

	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down log service")

}
