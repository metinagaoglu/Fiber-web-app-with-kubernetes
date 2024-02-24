package websocket

import (
	"context"
	"net"
	middleware "websocket-gateway/pkg/session/middleware"
)

func EndSessionMiddleware(ctx context.Context, conn net.Conn) {
	middleware.EndSession(ctx, conn)
}
