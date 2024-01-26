package service

import (
	"net/rpc"
	"sync"
	"log"
)

//Request of RPC call
type Args struct {
	Array []int
}

//Result of RPC call
type Response struct {
	Result []int
}

type Server struct {
	Address string
}

//LoadBalancer service for RPC
type LoadBalancer struct {
	Servers []*Server
	Current int
	Mutex   sync.Mutex
}

//Function used to select the next server to handle the request using Round Robin scheduling
func (lb *LoadBalancer) GetNextServer() string {
	lb.Mutex.Lock()
	defer lb.Mutex.Unlock()
	server := lb.Servers[lb.Current]
	lb.Current = (lb.Current + 1) % len(lb.Servers) //Round Robin 
	return server.Address
}

//Function used to route the request to the selected server
func (lb *LoadBalancer) HandleRequest(args Args, response *Response) error {
	serverAddress := lb.GetNextServer()

	client, err := rpc.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatal("Error dialing", err)
	}
	defer client.Close()

	err = client.Call("BubbleSort.Sort", args, &response.Result)
	if err != nil {
		log.Fatal("Error calling the service ", err)
	}
	
	return nil
}


