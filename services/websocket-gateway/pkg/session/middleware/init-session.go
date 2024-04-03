package session

import (
	"context"
	"fmt"
	"github.com/gobwas/ws/wsutil"
	"net"
	"strconv"

	epoll "websocket-gateway/internal/epoll"
	session "websocket-gateway/pkg/session"
	pb "websocket-gateway/pkg/session/pb"
)

func InitSession(ctx context.Context, conn net.Conn) {

	// Write Node ID
	nodeId := ctx.Value("nodeId")
	userID := ctx.Value("userId")

	fmt.Println("nodeId:", nodeId)
	fmt.Println("userID:", userID)

	connectionId := epoll.GetIdFromConn(conn)

	client := session.InitServiceClient()
	client.StartSession(context.Background(), &pb.StartSessionRequest{
		UserId:       userID.(int64),
		NodeId:       nodeId.(string),
		ConnectionId: strconv.Itoa(connectionId),
	})

	ctx.Done()
	wsutil.WriteServerMessage(conn, 1, []byte("inti-session"))
}
