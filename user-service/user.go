package main

import (
    "context"
    "fmt"
    pb "grpc-user-payment-services/gen/user"
    "grpc-user-payment-services/database"
)


// CreateUser method for the Server type
func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
    // Create a new user instance
    user := database.User{Name: req.Name, Email: req.Email}

    // Save user using GORM
    if err := database.DB.Create(&user).Error; err != nil {
        return nil, fmt.Errorf("failed to create user: %v", err)
    }

    return &pb.CreateUserResponse{
        Success: true,
        Message: "User created successfully",
        User:    &pb.User{Id: int32(user.ID), Name: user.Name, Email: user.Email}, // Convert uint to int32
    }, nil
}

// GetUser method for the Server type
func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
    var user database.User

    // Retrieve user using GORM
    if err := database.DB.First(&user, req.Id).Error; err != nil {
        return nil, fmt.Errorf("user not found: %v", err)
    }

    return &pb.GetUserResponse{
        User: &pb.User{Id: int32(user.ID), Name: user.Name, Email: user.Email}, // Convert uint to int32
    }, nil
}
