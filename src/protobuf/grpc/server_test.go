package grpc

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"log"
	"testing"
	"io"
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

	listener, err := net.Listen("tcp", ":8574")
	if err != nil {
		log.Fatalln("server listen err:", err)
	}

	server.Serve(listener)
}
