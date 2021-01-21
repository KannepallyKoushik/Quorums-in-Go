package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type API int

//The data that needs to be updated
var x int
var version int

func main() {
	x = 0
	version = 0

	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 8081)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}
}

func (a *API) Write(r int, reply *string) error {
	time.Sleep(time.Second * 1)
	x = r
	version = version + 1
	*reply = "succesfull"
	return nil
}
func (a *API) Read(empty string, reply *string) error {
	response := fmt.Sprintf("%d-%d", x, version)
	*reply = response
	return nil
}
