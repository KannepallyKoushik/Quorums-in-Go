package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client1, err1 := rpc.DialHTTP("tcp", "192.168.0.101:8081")

	if err1 != nil {
		log.Fatal("Connection error: ", err1)
	}

	fmt.Printf("%T", client1)
}
