package main

import (
    "context"
    "fmt"
    "log"
    "errors"

    pb "grpc-user-payment-services/gen/user"
    "grpc-user-payment-services/database"
    "gorm.io/gorm"
)

// CreateUser method for the Server type
func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
    // Create a new user instance
    user := database.User{Name: req.Name, Email: req.Email}

    // Save user using GORM
    if err := database.DB.Create(&user).Error; err != nil {
        return nil, fmt.Errorf("failed to create user: %v", err)
    }

    // Log the created user details
    log.Printf("Created user with ID: %d and Email: %s", user.ID, user.Email)

    return &pb.CreateUserResponse{
        Success: true,
        Message: "User created successfully",
        User:    &pb.User{Id: int32(user.ID), Name: user.Name, Email: user.Email},
    }, nil
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
    var user database.User

    // Log the incoming request
    log.Printf("Received GetUser request for user ID: %d", req.Id)

    // Retrieve user using GORM by ID
    err := database.DB.Where("id = ?", req.Id).First(&user).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, fmt.Errorf("user not found: %d", req.Id)
    } else if err != nil {
        return nil, fmt.Errorf("failed to retrieve user: %v", err)
    }

    log.Printf("Retrieved user with ID: %d and Email: %s", user.ID, user.Email)

    return &pb.GetUserResponse{
        User: &pb.User{Id: int32(user.ID), Name: user.Name, Email: user.Email},
    }, nil
}
