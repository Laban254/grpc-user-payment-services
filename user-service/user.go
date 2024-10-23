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

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
    user := database.User{Name: req.Name, Email: req.Email}

    if err := database.DB.Create(&user).Error; err != nil {
        return nil, fmt.Errorf("failed to create user: %v", err)
    }

    log.Printf("Created user with ID: %d and Email: %s", user.ID, user.Email)

    return &pb.CreateUserResponse{
        Success: true,
        Message: "User created successfully",
        User:    &pb.User{Id: int32(user.ID), Name: user.Name, Email: user.Email},
    }, nil
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
    var user database.User

    log.Printf("Received GetUser request for user ID: %d", req.Id)

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

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
    var user database.User

    log.Printf("Received UpdateUser request for user ID: %d", req.Id)

    if err := database.DB.Where("id = ?", req.Id).First(&user).Error; err != nil {
        return nil, fmt.Errorf("user not found: %d", req.Id)
    }

    user.Name = req.Name
    user.Email = req.Email

    if err := database.DB.Save(&user).Error; err != nil {
        return nil, fmt.Errorf("failed to update user: %v", err)
    }

    log.Printf("Updated user with ID: %d and Email: %s", user.ID, user.Email)

    return &pb.UpdateUserResponse{
        Success: true,
        Message: "User updated successfully",
    }, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
    log.Printf("Received DeleteUser request for user ID: %d", req.Id)

    if err := database.DB.Where("id = ?", req.Id).Delete(&database.User{}).Error; err != nil {
        return nil, fmt.Errorf("failed to delete user: %v", err)
    }

    log.Printf("Deleted user with ID: %d", req.Id)

    return &pb.DeleteUserResponse{
        Success: true,
        Message: "User deleted successfully",
    }, nil
}