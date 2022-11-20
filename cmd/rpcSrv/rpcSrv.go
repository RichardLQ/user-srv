package main

import (
	"context"
	"fmt"
	"github.com/RichardLQ/user-srv/proto/stub"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithInsecure()}
	err := stub.RegisterUserServiceHandlerFromEndpoint(context.Background(),
		gwmux, "localhost:8081", opt)
	if err != nil {
		log.Fatal(fmt.Errorf("启动rpc失败:%+v", err))
	}
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()
}
