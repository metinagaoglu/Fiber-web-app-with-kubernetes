package websocket

import (
	"context"
	"net"
	EndSession "websocket-gateway/pkg/session/middleware"
)

func EndSessionMiddleware(conn net.Conn, ctx context.Context) {
	EndSession.EndSession(conn, ctx)
}
