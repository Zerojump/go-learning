package main

import (
	"net/rpc"
	"net"
	"log"
)

type HelloService struct{}

func (p *HelloService) Hello(req string, reply *string) error {
	*reply = "hello:" + req
	log.Println("msg from client:", req)
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":8547")
	if err != nil {
		log.Fatalln("listenTcp err:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalln("accept err:", err)
	}

	rpc.ServeConn(conn)

}
