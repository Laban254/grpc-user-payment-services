package main

import (
    "log"
    "net"

    pb "grpc-user-payment-services/gen/user"
    "grpc-user-payment-services/database"

    "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Exported Server type
type Server struct {
    pb.UnimplementedUserServiceServer
}

func main() {
    // Connect to the database
    database.ConnectDB() 

    // Set up listener for gRPC server
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterUserServiceServer(grpcServer, &Server{})
	reflection.Register(grpcServer)

    log.Println("User Service is running on port 50051")
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
