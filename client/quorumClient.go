package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"strings"
	"sync"
)

//waitgroups to let the main function wait until all goroutines are finished
var wg = sync.WaitGroup{}

var reply1 string
var reply2 string
var reply3 string
var reply4 string
var reply5 string
var reply6 string

func main() {

	client1, err1 := rpc.DialHTTP("tcp", "192.168.0.101:8081")

	if err1 != nil {
		log.Fatal("Connection error: ", err1)
	}
	client2, err2 := rpc.DialHTTP("tcp", "192.168.0.101:8082")

	if err2 != nil {
		log.Fatal("Connection error: ", err2)
	}
	client3, err3 := rpc.DialHTTP("tcp", "192.168.0.101:8083")

	if err3 != nil {
		log.Fatal("Connection error: ", err3)
	}
	client4, err4 := rpc.DialHTTP("tcp", "192.168.0.101:8084")

	if err4 != nil {
		log.Fatal("Connection error: ", err4)
	}
	client5, err5 := rpc.DialHTTP("tcp", "192.168.0.101:8085")

	if err5 != nil {
		log.Fatal("Connection error: ", err5)
	}
	client6, err6 := rpc.DialHTTP("tcp", "192.168.0.101:8086")

	if err6 != nil {
		log.Fatal("Connection error: ", err6)
	}

	for true {
		fmt.Println("\n Select an option")
		fmt.Println("1.write into X")
		fmt.Println("2.Read X")
		fmt.Println("3.Exit")
		var option int

		fmt.Scanln(&option)
		switch option {
		case 1:
			//writing into write quorum i.e., server1, server2, server3, server4
			fmt.Println("Enter the value of X which u want to Update")
			var x int
			fmt.Scanln(&x)
			wg.Add(4)
			go rpcCallWrite(client1, x)
			go rpcCallWrite(client2, x)
			go rpcCallWrite(client3, x)
			go rpcCallWrite(client4, x)
			wg.Wait()
			fmt.Println("\nAll the Replicas in Write Quorum are Updated Succesfully\n")
		case 2:
			//Read the value from read quorum i.e., server4, server5, server6
			client4.Call("API.Read", "", &reply4)
			client5.Call("API.Read", "", &reply5)
			client6.Call("API.Read", "", &reply6)

			res4 := strings.Split(reply4, "-")
			res5 := strings.Split(reply5, "-")
			res6 := strings.Split(reply6, "-")

			r40, err := strconv.Atoi(res4[0])
			if err != nil {
				fmt.Println("Cannot convert string to int")
			}
			r41, err := strconv.Atoi(res4[1])
			if err != nil {
				fmt.Println("Cannot convert string to int")
			}

			r51, err := strconv.Atoi(res5[1])
			if err != nil {
				fmt.Println("Cannot convert string to int")
			}

			r61, err := strconv.Atoi(res6[1])
			if err != nil {
				fmt.Println("Cannot convert string to int")
			}

			if r41 != r51 || r41 != r61 {
				fmt.Println("\nAll the version Numbers in each replica differs")
				fmt.Println("The most recently Updated Value of x is ", r40)
				fmt.Println("Now the other replicas will get updated with Latest value with Peer to Peer Communication")
				client5.Call("API.WriteVersion", reply4, &reply5)
				client6.Call("API.WriteVersion", reply4, &reply6)

			} else {
				fmt.Println("\nAll the Replicas have Consistent value of x which is", r40)
			}
		case 3:
			fmt.Println()
			fmt.Println("Thankyou for using our Application")
			os.Exit(3)

		default:
			fmt.Println("Choose from the options given")
		}

	}
}

func rpcCallWrite(clientObj *rpc.Client, x int) {
	clientObj.Call("API.Write", x, &reply1)
	wg.Done()
}
