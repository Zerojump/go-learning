package grpc

import (
	"google.golang.org/grpc"
	"log"
	"context"
	"testing"
	"time"
	"io"
)


func TestClient(t *testing.T) {
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

	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			if err := stream.Send(&String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		log.Println(reply.GetValue())
	}
}
