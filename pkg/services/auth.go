package services

import (
	"context"
	"log"

	"github.com/hebobibun/grpc-auth-crud/pkg/db"
	"github.com/hebobibun/grpc-auth-crud/pkg/models"
	pb "github.com/hebobibun/grpc-auth-crud/pkg/pb"
	"github.com/hebobibun/grpc-auth-crud/pkg/utils"
)

type AuthServer struct {
	H   db.Handler
	Jwt utils.JWTwrapper
}

func (s *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if res := s.H.DB.Where("email = ?", req.Email).First(&user); res.Error == nil {
		return &pb.RegisterResponse{
			Status:  "error",
			Message: "Email already exists",
		}, nil
	}

	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)
	user.Name = req.Name
	user.RoleId = req.RoleId

	err := s.H.DB.Create(&user)
	log.Println(err)
	if err != nil {
		return &pb.RegisterResponse{
			Status:  "error",
			Message: "Error creating user",
		}, nil
	}

	return &pb.RegisterResponse{
		Status:  "success",
		Message: "User registered successfully",
	}, nil

}

func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if res := s.H.DB.Where("email = ?", req.Email).First(&user); res.Error != nil {
		return &pb.LoginResponse{
			Status:  "error",
			Message: "User not found",
		}, nil
	}

	if utils.CheckPassword(req.Password, user.Password) == false {
		return &pb.LoginResponse{
			Status:  "error",
			Message: "Invalid password",
		}, nil
	}

	token, err := s.Jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		return &pb.LoginResponse{
			Status:  "error",
			Message: "Error generating token",
		}, nil
	}

	return &pb.LoginResponse{
		Status:  "success",
		Message: "User logged in successfully",
		Data: &pb.LoginData{
			AccessToken: token,
		},
	}, nil
}

func (s *AuthServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return &pb.UpdateUserResponse{
		Status:  "success",
		Message: "User updated successfully",
	}, nil
}
