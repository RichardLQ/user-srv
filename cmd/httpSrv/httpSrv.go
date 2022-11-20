package main

import (
	"github.com/RichardLQ/confs"
	"github.com/RichardLQ/user-srv/client"
	"github.com/RichardLQ/user-srv/proto/stub"
	"google.golang.org/grpc"
	"net"
)

func main() {
	confs.NewStart().BinComb(&client.Global)
	rpcServer := grpc.NewServer()
	stub.RegisterUserServiceServer(rpcServer, stub.UserService)
	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)
}
