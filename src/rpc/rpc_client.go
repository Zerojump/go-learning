package main

import (
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8547")
	if err != nil {
		log.Fatalln("diali err:", err)
	}

	var reply string
	client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatalln("call err:", err)
	}

	log.Println("reply from server:", reply)
}
