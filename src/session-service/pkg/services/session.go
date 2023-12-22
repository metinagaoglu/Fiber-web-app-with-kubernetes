package services

import (
	"log"
    "context"
    "net/http"
    "fmt"

    "session-service/pkg/pb"
    "github.com/redis/go-redis/v9"
)

type Server struct {
	pb.UnsafeSessionServiceServer
    Rdb *redis.Client
}

func (s *Server) StartSession(ctx context.Context, req *pb.StartSessionRequest) (*pb.StartSessionResponse, error) {
	log.Println("[/register]: Registering user")

    sessionKey := fmt.Sprintf("session:%d", req.UserId)
    s.Rdb.HSet(ctx,sessionKey, "nodeId", req.NodeId)
    s.Rdb.HSet(ctx,sessionKey, "connectionId", req.ConnectionId)

    return &pb.StartSessionResponse{
        Status: http.StatusCreated,
    }, nil
}

func (s *Server) EndSession(ctx context.Context, req *pb.EndSessionRequest) (*pb.EndSessionResponse, error) {
	log.Println("[/login]: Logging in user")

    sessionKey := fmt.Sprintf("session:%d", req.UserId)
    s.Rdb.Del(ctx, sessionKey)

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
