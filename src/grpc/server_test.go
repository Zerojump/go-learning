package grpc

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"log"
	"testing"
	"io"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"fmt"
	"golang.org/x/net/trace"
	"net/http"
)

const (
	port    = ":8574"
	app_id  = "cid"
	app_key = "ckey"
)

type HelloServiceImpl struct {
}

func (*HelloServiceImpl) Hello(ctx context.Context, arg *String) (*String, error) {
	log.Println("msg from client:", arg.Value)
	return &String{Value: "hello:" + arg.GetValue()}, nil
}

func (*HelloServiceImpl) Channel(stream HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		log.Println("msg from client:", args.Value)
		reply := &String{Value: "hello:" + args.GetValue()}
		if err := stream.Send(reply); err != nil {
			return err
		}
	}

}

//gRPC和标准库的RPC框架有一个区别，gRPC生成的接口并不支持异步调用。不过我们可以在多个Goroutine之间安全地共享gRPC底层的HTTP/2链接，
// 因此可以通过在另一个Goroutine阻塞调用的方式模拟异步调用。
//RPC是远程函数调用，因此每次调用的函数参数和返回值不能太大，否则将严重影响每次调用的响应时间。
// 因此传统的RPC方法调用对于上传和下载较大数据量场景并不适合。同时传统RPC模式也不适用于对时间不确定的订阅和发布模式。为此，gRPC框架针对服务器端和客户端分别提供了流特性。

//protoc --go_out=plugins=grpc:. hello.proto
//GOPATH=$PROJECT_DIR/src/go-learning:$GOPATH #gosetup
//go test -c -i -o /tmp/___TestServer_in_protobuf_grpc__1_ protobuf/grpc #gosetup
//go tool test2json -t /tmp/___TestServer_in_protobuf_grpc__1_ -test.v -test.run ^TestServer$ #gosetup
func TestServer(t *testing.T) {
	server := grpc.NewServer()
	RegisterHelloServiceServer(server, new(HelloServiceImpl))

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("server listen err:", err)
	}

	server.Serve(listener)
}

//server.key
//openssl genrsa -out server.key 2048
//openssl ecparam -genkey -name secp384r1 -out server.key

//server.pem
//openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
//go test -v server_test.go hello.pb.go -test.run TestTLS_Server
func TestTLS_Server(t *testing.T) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	//TLS 认证
	creds, err := credentials.NewServerTLSFromFile("tls/server.pem", "tls/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credential %v", err)
	}

	//实例化grpc Server，并开启TLS认证
	server := grpc.NewServer(grpc.Creds(creds))
	//注册HelloService
	RegisterHelloServiceServer(server, new(HelloServiceImpl))

	log.Println("Listen on localhost" + port + " with TLS")
	server.Serve(listener)
}

func TestInterceptorServer(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost"+port)
	if err != nil {
	    fmt.Errorf("%v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("tls/server.pem", "tls/server.key")
	if err != nil {
		fmt.Errorf("%v", err)
	}

	server := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(interceptor))
	RegisterHelloServiceServer(server,new(HelloServiceImpl))
	log.Println("Listen on localhost" + port + " with TLS and Interceptor")
	server.Serve(listener)
}

func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	err := auth(ctx)
	if err != nil {
		return nil, err
	}
	//继续执行
	return handler(ctx, req)
}

func auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "无token认证信息")
	}
	var appid, appkey string
	if val, ok := md["appid"]; ok {
		appid = val[0]
	}
	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}

	if appid != app_id || appkey != app_key {
		return status.Errorf(codes.Unauthenticated, "Invalid token! appid:%s, appkey:%s", appid, appkey)
	}
	return nil
}

func TestTraceServer(t *testing.T) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	server := grpc.NewServer()
	RegisterHelloServiceServer(server,new(HelloServiceImpl))

	go startTrace()

	fmt.Println("Listen on",port)
	server.Serve(listener)
}

func startTrace()  {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true,true
	}
	go http.ListenAndServe(":8704", nil)
	fmt.Println("Trace listen on 8704")

}