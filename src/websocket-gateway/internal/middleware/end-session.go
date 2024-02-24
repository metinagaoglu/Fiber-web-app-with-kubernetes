package websocket

import (
	"context"
	"net"
	middleware "websocket-gateway/pkg/session/middleware"
)

func EndSessionMiddleware(conn net.Conn, ctx context.Context) {
	middleware.EndSession(conn, ctx)
}
