package main

import (
    "log"
    "net"

    pb "grpc-user-payment-services/gen/payment"
    "grpc-user-payment-services/database"
    "google.golang.org/grpc"
)

// Server is used to implement the PaymentServiceServer
type Server struct {
    pb.UnimplementedPaymentServiceServer
}

func main() {
    // Connect to the database
    database.ConnectDB() // Just call this without error handling

    // Set up a TCP listener on port 50052
    listener, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    // Create a new gRPC server
    grpcServer := grpc.NewServer()

    // Register the payment service
    pb.RegisterPaymentServiceServer(grpcServer, &Server{})

    log.Println("Payment Service is running on port 50052")

    // Start serving
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
