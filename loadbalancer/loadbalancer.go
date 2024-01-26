package main

import (
	"fmt"
	"net"
	"net/rpc"
	"log"
	"os"
	"bufio"

	"loadbalancer/service"
)

func main() {

	//Open and read the configuration file to get the addresses of the replicated servers
	file, err := os.Open("configurationFile.txt")
	if(err != nil) {
		log.Fatal(err)
	}

	fileReader := bufio.NewScanner(file)
	fileReader.Split(bufio.ScanLines)

	var serverAddresses []string
	for fileReader.Scan() {
		serverAddresses = append(serverAddresses, fileReader.Text())
	}
	file.Close()

	var servers []*service.Server
	for _, addr := range serverAddresses {
		servers = append(servers, &service.Server{Address: addr})
	}

	//Create an instance of the service load balancer
	loadBalancer := &service.LoadBalancer{Servers: servers}

	server := rpc.NewServer()
	err = server.RegisterName("SortArray", loadBalancer)
	if err != nil {
		log.Fatal("Error in Register Name ", err)
	}

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("Error listening on port 8888 ", err)
	}

	fmt.Println("Load balancer started on port 8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection ", err)
			continue
		}
		go server.ServeConn(conn)
	}

}
