package main

import (
    "context"
    "fmt"
    "log"

    pb "grpc-user-payment-services/gen/payment"
    userpb "grpc-user-payment-services/gen/user"
    "google.golang.org/grpc"
    "grpc-user-payment-services/database"
)


func (s *Server) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {

    userConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        return nil, fmt.Errorf("failed to connect to user service: %v", err)
    }
    defer userConn.Close()

    userClient := userpb.NewUserServiceClient(userConn)

    userResp, err := userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId})
    if err != nil {
        return &pb.PaymentResponse{
            Success: false,
            Message: fmt.Sprintf("Error retrieving user: %v", err),
        }, nil
    }

    log.Printf("Processing payment for user: %s", userResp.User.Email)

    transaction := database.Transaction{
        UserID: userResp.User.Id,
        Amount: req.Amount,
        Status: "success",
    }
    if err := database.DB.Create(&transaction).Error; err != nil {
        return nil, fmt.Errorf("failed to create transaction: %v", err)
    }

    return &pb.PaymentResponse{
        Success: true,
        Message: "Payment processed successfully",
    }, nil
}

func (s *Server) RefundTransaction(ctx context.Context, req *pb.RefundRequest) (*pb.RefundResponse, error) {

    var transaction database.Transaction
    if err := database.DB.First(&transaction, req.TransactionId).Error; err != nil {
        return nil, fmt.Errorf("transaction not found: %v", err)
    }

    refund := database.Transaction{
        UserID: transaction.UserID,
        Amount: -transaction.Amount, 
        Status: "refunded",           
    }
    if err := database.DB.Create(&refund).Error; err != nil {
        return nil, fmt.Errorf("failed to create refund transaction: %v", err)
    }

    return &pb.RefundResponse{
        Success: true,
        Message: "Refund processed successfully",
    }, nil
}

func (s *Server) CheckBalance(ctx context.Context, req *pb.BalanceRequest) (*pb.BalanceResponse, error) {
    userConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        return nil, fmt.Errorf("failed to connect to user service: %v", err)
    }
    defer userConn.Close()

    userClient := userpb.NewUserServiceClient(userConn)

    userResp, err := userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId})
    if err != nil {
        return nil, fmt.Errorf("user not found: %v", err)
    }

    log.Printf("Checking balance for user: %s", userResp.User.Email)

    var totalBalance float64
    if err := database.DB.Model(&database.Transaction{}).Where("user_id = ?", req.UserId).Select("sum(amount)").Scan(&totalBalance).Error; err != nil {
        return nil, fmt.Errorf("failed to calculate balance: %v", err)
    }

    return &pb.BalanceResponse{
        Balance: totalBalance,
    }, nil
}


