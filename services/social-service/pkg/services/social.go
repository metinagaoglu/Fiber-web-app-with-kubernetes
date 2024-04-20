package services

import (
	"context"
	"fmt"
	"net/http"

	"social-service/pkg/db"
	// "social-service/pkg/models"
	"social-service/pkg/pb"
)

type Server struct {
	pb.UnsafeSocialServiceServer
	H db.Handler
}

func (s *Server) AddFollower(ctx context.Context, req *pb.AddFollowerRequest) (*pb.AddFollowerResponse, error) {
	fmt.Println(req)
	return &pb.AddFollowerResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) RemoveFollower(ctx context.Context, req *pb.RemoveFollowerRequest) (*pb.RemoveFollowerResponse, error) {
	fmt.Println(req)
	return &pb.RemoveFollowerResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) GetFollowers(ctx context.Context, req *pb.GetFollowersRequest) (*pb.GetFollowersResponse, error) {
	return &pb.GetFollowersResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) CountFollowers(ctx context.Context, req *pb.CountFollowersRequest) (*pb.CountFollowersResponse, error) {
	return &pb.CountFollowersResponse{
		Status: http.StatusOK,
	}, nil
}
