package main

import (
    "log"
    "net"

    pb "grpc-user-payment-services/gen/payment"
    "grpc-user-payment-services/database"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

// Server is used to implement the PaymentServiceServer
type Server struct {
    pb.UnimplementedPaymentServiceServer
}

func main() {
    database.ConnectDB()

    listener, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()

    pb.RegisterPaymentServiceServer(grpcServer, &Server{})

    reflection.Register(grpcServer)

    log.Println("Payment Service is running on port 50052")

    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
