package main

import (
	"fmt"
	"net/rpc"
	"sync"
	"log"

	"client/service" //Path to the package which contains service definition
)

//Send the request to the load balancer
func sendRequest(client *rpc.Client, wg *sync.WaitGroup, args *service.Args, index int, resultMap *sync.Map) {
	defer wg.Done()

	response := &service.Response{}
	err := client.Call("SortArray.HandleRequest", args, response)
	if err != nil {
		log.Fatal("Error calling HandleRequest", err)
	}
	
	//The result and the corrisponding index of the request are mapped together
	resultMap.Store(index, response.Result)
}

func main() {

	addr := "loadbalancer:8888" //Address and port on which load balancer is listening
	// Try to connect to load balancer
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		log.Fatal("Error connecting to load balancer", err)
	}
	defer client.Close()

	// WaitGroup to wait the end of goroutines' execution
	var wg sync.WaitGroup

	resultMap := &sync.Map{}

	//Array to sort that are sent in the requests
	requests := [5][]int{{0,4,1,3}, {1,9,8,7,3}, {5,2,4,10}, {3,12,6,2,5,11,4}, {4,8,3,2}}
	
	fmt.Println("Array to sort:")
	for j:=0; j<= 4; j++ {
		fmt.Println(requests[j])
	}

	//Multiple requests are sent to the load balancer asynchronously
	for i := 0; i <= 4; i++ {
		wg.Add(1)
		args := &service.Args{requests[i]}
		go sendRequest(client, &wg, args, i, resultMap)
	}

	wg.Wait()
	
	fmt.Println("Sorted array:")

	//All the responses in the map are printed with the order specified in the index of the request
	for i := 0; i <= 4; i++ {
		result, found := resultMap.Load(i)
		if found {
			fmt.Println(result)
		} else {
			fmt.Println("Response not found for index", i)
		}
	}

}
