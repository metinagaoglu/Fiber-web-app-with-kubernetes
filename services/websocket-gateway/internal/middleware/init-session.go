package websocket

import (
	"context"
	"net"
	middleware "websocket-gateway/pkg/session/middleware"
)

func InitSessionMiddleware(ctx context.Context, conn net.Conn) {
	middleware.InitSession(ctx, conn)
}
