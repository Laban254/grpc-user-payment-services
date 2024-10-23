package main

import (
    "context"
    "fmt"
    "log"

    pb "grpc-user-payment-services/gen/payment"
    userpb "grpc-user-payment-services/gen/user"
    "google.golang.org/grpc"
)


func (s *Server) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
    log.Printf("Received payment request for user ID: %d, amount: %f", req.UserId, req.Amount)

    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        return nil, fmt.Errorf("failed to connect to user service: %v", err)
    }
    defer conn.Close()

    userClient := userpb.NewUserServiceClient(conn)

    userResp, err := userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId})
    if err != nil {
        return &pb.PaymentResponse{
            Success: false,
            Message: fmt.Sprintf("Error retrieving user: %v", err),
        }, nil
    }

    log.Printf("Processing payment for user: %s", userResp.User.Email)
    log.Printf("Processing payment of %f for user ID: %d", req.Amount, userResp.User.Id)

    return &pb.PaymentResponse{
        Success: true,
        Message: "Payment processed successfully",
    }, nil
}

func (s *Server) CheckBalance(ctx context.Context, req *pb.BalanceRequest) (*pb.BalanceResponse, error) {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        return nil, fmt.Errorf("failed to connect to user service: %v", err)
    }
    defer conn.Close()

    userClient := userpb.NewUserServiceClient(conn)

    userResp, err := userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId})
    if err != nil {
        return nil, fmt.Errorf("user not found: %v", err)
    }

    log.Printf("Checking balance for user: %s", userResp.User.Email)

    balance := 100.0

    return &pb.BalanceResponse{
        Balance: balance,
    }, nil
}
