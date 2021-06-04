package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, serviceName, host, port string, registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, serviceName, host, port)
	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = ":" + port

	// 错误后停止
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	// 手动输入任意字符后停止
	go func() {
		fmt.Printf("%v started. Press any key to stop...\n", serviceName)
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx

}
