package broadcast

import (
	"context"
	. "go-apps-with-kubernetes/libs/websocket"
)

func HandleJoin(ctx context.Context,conn *Client, request JsonRPCRequest) {
	conn.Join("my_room")
}

func HandleBroadcastToRoom(ctx context.Context,conn *Client, request JsonRPCRequest) {
	BroadcastToGroup("my_room", []byte("request"))
}

func HandleLeave(ctx context.Context,conn *Client, request JsonRPCRequest) {
	conn.Leave("my_room")
}