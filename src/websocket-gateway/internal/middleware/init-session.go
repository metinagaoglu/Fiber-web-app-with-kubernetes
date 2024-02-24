package websocket

import (
	"context"
	"net"
	middleware "websocket-gateway/pkg/session/middleware"
)

func InitSessionMiddleware(conn net.Conn, ctx context.Context) {
	middleware.InitSession(conn, ctx)
}
