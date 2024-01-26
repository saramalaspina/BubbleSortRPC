The aim of the project is to implement a mechanism to support in RCP the transparency to replication through a server-side load balancer that acts as a proxy between the clients and the replicated servers that provide
the RPC service.

We realized a service that implements the bubble sort algorithm and we used it to sort vectors of integers that are sent by the client in the request.

The client request is handled by a load balancer which then routes it to the replicated servers following a Round Robin scheduling.

We containerized our distributed application by creating a docker file for every component (client, load balancer and server) and a compose file to orchestrate the container on the same local host.

Steps to build the images and run the containers and the application:

1. Open the BubbleSortRPC project folder in your terminal.

2. Build the Client's image
   
        cd client
   
        docker build -t client:1 .

3. Build the LoadBalancer's image
   
        cd ..
   
        cd loadbalancer
   
        docker build -t loadbalancer:1 .

4. Build the Server's image
   
        cd ..
   
        cd server
   
        docker build -t server:1 .

5. Run the containers
   
	   cd ..
   
	   docker-compose up

Then you can see on your terminal the connection of the servers and the load balancer and the results of the RPC call.





