package services

import (
	"log"
    "context"
    "net/http"


    "session-service/pkg/pb"
)

type Server struct {
	pb.UnsafeSessionServiceServer
}

func (s *Server) StartSession(ctx context.Context, req *pb.StartSessionRequest) (*pb.StartSessionResponse, error) {
		log.Println("[/register]: Registering user")

    return &pb.StartSessionResponse{
        Status: http.StatusCreated,
    }, nil
}

func (s *Server) EndSession(ctx context.Context, req *pb.EndSessionRequest) (*pb.EndSessionResponse, error) {
	  log.Println("[/login]: Logging in user")

    return &pb.EndSessionResponse{
        Status: http.StatusOK,
    }, nil
}

func (s *Server) GetSession(ctx context.Context, req *pb.GetSessionRequest) (*pb.GetSessionResponse, error) {
		log.Println("[/validate]: Validating token")
		
		return &pb.GetSessionResponse{
        Status: http.StatusOK,
				NodeId: "asdasd",
    }, nil
}
