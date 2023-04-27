package main

import (
	"ezyevent-location-api/internal/handler"
	"ezyevent-location-api/internal/proto"
	route "ezyevent-location-api/internal/route_guide"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	//gRPC Connection establishing using tcp
	con, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	//Initializing Fiber
	app := fiber.New()
	handler.InitRoutes(app)

	//Initializing Grpc
	rpc := grpc.NewServer()
	proto.RegisterLocationDataServiceServer(rpc, &route.Server{})

	//Handling gRPC server errors
	rpcError := rpc.Serve(con)
	if rpcError != nil {
		log.Fatal(rpcError)
	}

	//Finally
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatalln(app.Listen("0.0.0.0:" + port))
}
