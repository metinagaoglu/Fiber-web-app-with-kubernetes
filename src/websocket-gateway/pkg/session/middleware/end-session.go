package session

import (
	"context"
	"net"

	session "websocket-gateway/pkg/session"
	pb "websocket-gateway/pkg/session/pb"
)


func EndSession(ctx context.Context, conn net.Conn) {

	// Write Node ID
	userId := ctx.Value("userId")

	client := session.InitServiceClient()
	client.EndSession(context.Background(), &pb.EndSessionRequest{
		UserId: userId.(int64),
	})

	ctx.Done()
}
