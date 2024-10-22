package main

import (
    "context"
    "fmt"
    pb "grpc-user-payment-services/gen/payment"
    "grpc-user-payment-services/database"
)



// ProcessPayment processes a payment request
func (s *Server) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
    var account struct {
        Balance float64 `gorm:"column:balance"`
    }

    // Check user balance in PostgreSQL using GORM
    if err := database.DB.Model(&account).Where("user_id = ?", req.UserId).First(&account).Error; err != nil {
        return &pb.PaymentResponse{
            Success: false,
            Message: "User not found",
        }, nil
    }

    if account.Balance < req.Amount {
        return &pb.PaymentResponse{
            Success: false,
            Message: "Insufficient funds",
        }, nil
    }

    // Deduct amount from balance
    newBalance := account.Balance - req.Amount
    if err := database.DB.Model(&account).Where("user_id = ?", req.UserId).Update("balance", newBalance).Error; err != nil {
        return nil, fmt.Errorf("failed to update balance: %v", err)
    }

    return &pb.PaymentResponse{
        Success: true,
        Message: "Payment processed successfully",
    }, nil
}

// CheckBalance retrieves the balance for a user
func (s *Server) CheckBalance(ctx context.Context, req *pb.BalanceRequest) (*pb.BalanceResponse, error) {
    var account struct {
        Balance float64 `gorm:"column:balance"`
    }

    // Retrieve user balance using GORM
    if err := database.DB.Model(&account).Where("user_id = ?", req.UserId).First(&account).Error; err != nil {
        return nil, fmt.Errorf("user not found: %v", err)
    }

    return &pb.BalanceResponse{
        Balance: account.Balance,
    }, nil
}
