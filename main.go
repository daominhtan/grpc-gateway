package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/daominhtan/grpc-gateway/config"
	grpc_gateway "github.com/daominhtan/grpc-gateway/proto"
	pb "github.com/daominhtan/grpc-gateway/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {

	fmt.Println("Request = ", request)

	name := request.Name
	response := &pb.HelloResponse{
		Message: "Hello " + name,
	}
	return response, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen(config.DefaultGRPCServerConfig.Network, config.DefaultGRPCServerConfig.Address)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterGreeterServer(s, &server{})
	// Serve gRPC server
	log.Println("Serving gRPC on connection ")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	conn, err := grpc.Dial(config.DefaultGRPCServerConfig.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux()
	// Register Greeter

	err = grpc_gateway.RegisterGreeterHandlerServer(context.Background(), mux, &server{})

	// err = pb.RegisterGreeterHandler(, mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    config.DefaultReverseConfig.Address,
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on connection")
	log.Fatalln(gwServer.ListenAndServe())
}
