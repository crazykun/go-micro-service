package main

import (
	"context"
	"fmt"
	"go-micro-service/log"
	"go-micro-service/service"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	ctx, err := service.Start(
		context.Background(),
		"Log Service",
		host,
		port,
		log.RegisterHandler,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down log service")

}
