package main

import (
	"google.golang.org/grpc"
	"log"
	"context"
)

func main() {
	conn, err := grpc.Dial("localhost:8574", grpc.WithInsecure())
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		log.Fatalln("client dial err:", err)
	}
	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value: "hello"})
	if err != nil {
		log.Fatalln("server reply err:", err)
	}

	log.Println(reply.GetValue())

}
