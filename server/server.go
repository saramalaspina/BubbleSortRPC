package main

import (
	"fmt"
	"net"
	"net/rpc"
	"log"
	"os"

	"server/bubbleSort" //Path to the package which contains service definition
)


func main() {

	//Create an instance of struct which implements service interface
	bubbleSort := new(bubbleSort.BubbleSort)
	rpc.RegisterName("BubbleSort", bubbleSort)

	//Takes the port of the current server from the environment variable specified in the compose file
	port := os.Getenv("PORT")
	if(port == ""){
		fmt.Println("Error to get env port")
		os.Exit(1)
	}
	addr := ":"+port

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Error Listen",err)
	}

	defer listener.Close()

	fmt.Println("Server started on port",port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}