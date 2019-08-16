package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"log"
)

type HelloServiceImpl struct {
}

func (*HelloServiceImpl) Hello(ctx context.Context, arg *String) (*String, error) {
	return &String{Value: "hello:" + arg.GetValue()}, nil
}

func main() {
	server := grpc.NewServer()
	RegisterHelloServiceServer(server,new(HelloServiceImpl))

	listener, err := net.Listen("tcp", ":8574")
	if err!= nil {
		log.Fatalln("server listen err:",err)
	}

	server.Serve(listener)
}
