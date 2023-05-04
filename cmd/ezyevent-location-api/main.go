package main

import (
	"ezyevent-location-api/internal/proto"
	route "ezyevent-location-api/internal/route_guide"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	//gRPC Connection establishing using tcp
	con, err := net.Listen("tcp", ":8181")
	if err != nil {
		log.Fatal(err)
	}

	//Initializing Grpc
	rpc := grpc.NewServer()
	proto.RegisterLocationDataServiceServer(rpc, &route.Server{})

	//Handling gRPC server errors
	rpcError := rpc.Serve(con)
	if rpcError != nil {
		log.Fatal(rpcError)
	}
}
