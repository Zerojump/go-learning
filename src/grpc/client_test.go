package grpc

import (
	"google.golang.org/grpc"
	"log"
	"context"
	"testing"
	"time"
	"io"
	"google.golang.org/grpc/credentials"
	"fmt"
)

const (
	OPEN_TLS = true
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

//go test -v client_test.go hello.pb.go -test.run TestTLS_Client
func TestTLS_Client(t *testing.T) {
	creds, err := credentials.NewClientTLSFromFile("tls/server.pem", "hello.cmy.fun")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial("localhost:8574", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	//初始化客户端
	client := NewHelloServiceClient(conn)
	//调用方法
	reply, err := client.Hello(context.Background(), &String{Value: "hello"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("reply from server", reply.GetValue())
}

//自定义认证
type customCredential struct{}

//实现自定义认证接口
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  app_id,
		"appkey": app_key,
	}, nil
}

//自定义认证是否开启TLS
func (c customCredential) RequireTransportSecurity() bool {
	return OPEN_TLS
}

func TestInterceptorClient(t *testing.T) {
	creds, err := credentials.NewClientTLSFromFile("tls/server.pem", "hello.cmy.fun")
	if err != nil {
		fmt.Errorf("%v", err)
	}
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(new(customCredential)), grpc.WithUnaryInterceptor(interceptorClient))
	if err != nil {
		fmt.Errorf("%v", err)
	}
	defer conn.Close()

	//初始化客户端
	client := NewHelloServiceClient(conn)
	//调用方法
	reply, err := client.Hello(context.Background(), &String{Value: "hello"})
	if err != nil {
		fmt.Errorf("%v", err)
	} else {
		fmt.Println("reply from server:", reply.GetValue())
	}
}

func interceptorClient(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("method:%s, req:%v resp:%v duration:%s, err:%v\n", method, req, reply, time.Since(start), err)
	return err

}
