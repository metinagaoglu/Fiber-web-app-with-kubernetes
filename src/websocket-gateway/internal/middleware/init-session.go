package websocket

import (
	"context"
	"net"
	InitSession "websocket-gateway/pkg/session/middleware"
)

func InitSessionMiddleware(conn net.Conn, ctx context.Context) {
	InitSession.InitSession(conn, ctx)
}
