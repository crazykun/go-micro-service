package service

import (
	"context"
	"fmt"
	"go-micro-service/registry"
	"log"
	"net/http"
)

func Start(ctx context.Context, host, port string, reg registry.Registration, registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, reg.ServiceName, host, port)
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = ":" + port

	// 错误后停止
	go func() {
		log.Println(srv.ListenAndServe())
		err := registry.ShutdownService(registry.ServiceUrl)
		if err != nil {
			log.Println(err)
		}
		cancel()
	}()

	// 手动输入任意字符后停止
	go func() {
		fmt.Printf("%v started. Press any key to stop...\n", serviceName)
		var s string
		fmt.Scanln(&s)
		err := registry.ShutdownService(registry.ServiceUrl)
		if err != nil {
			log.Println(err)
		}
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx

}
