package routes

import (
	"fmt"
	"net"
	"github.com/gobwas/ws/wsutil"
	"context"

		epoll "websocket-gateway/internal/epoll"
)

type InitSessionHandler struct{}

type InitSessionHandlerBody struct {
    Token  string `json:"token"`
}

func (h *InitSessionHandler) HandleMessage(conn *net.Conn,ctx context.Context, route string, payload string) {
		fmt.Println("id:", epoll.GetIdFromConn(*conn))
		ctx.Done()
		wsutil.WriteServerMessage(*conn, 1, []byte("Session initialized"))
}