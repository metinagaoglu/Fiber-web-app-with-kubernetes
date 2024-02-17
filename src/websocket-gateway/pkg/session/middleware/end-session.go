package session

import (
	"context"
	"net"

	session "websocket-gateway/pkg/session"
	pb "websocket-gateway/pkg/session/pb"
)


func EndSession(conn net.Conn, ctx context.Context) {

	// Write Node ID
	userId := ctx.Value("userId")

	// String cast


	client := session.InitServiceClient()
	client.EndSession(context.Background(), &pb.EndSessionRequest{
		UserId: userId.(int64),
	})

	ctx.Done()
}
