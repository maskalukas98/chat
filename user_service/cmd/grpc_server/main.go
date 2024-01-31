package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	grpc2 "user_service/src/infrastructure/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatal("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	grpc2.RegisterGrpcServices(grpcServer)

	log.Print("Grpc server started on port 9000")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server gRPC server over port 9000: %v", err)
	} else {
		log.Print("Grpc server started on port 9000")
	}
}
