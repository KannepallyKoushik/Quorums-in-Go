package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
	"strings"
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
	listener, err := net.Listen("tcp", ":8086")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 8086)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}
}

func (a *API) Write(r int, reply *string) error {
	time.Sleep(3)
	x = r
	version = version + 1
	*reply = "succesfull"
	return nil
}

func (a *API) WriteVersion(req string, reply *string) error {
	time.Sleep(time.Second * 3)
	r := strings.Split(req, "-")
	xval, err := strconv.Atoi(r[0])
	if err != nil {
		fmt.Println("Cannot convert string to int")
	}
	versionval, err := strconv.Atoi(r[1])
	if err != nil {
		fmt.Println("Cannot convert string to int")
	}

	x = xval
	version = versionval
	*reply = "succesfull"
	return nil
}

func (a *API) Read(empty string, reply *string) error {
	response := fmt.Sprintf("%d-%d", x, version)
	*reply = response
	return nil
}
