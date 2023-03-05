package broadcast

import (
	"context"

	. "go-apps-with-kubernetes/libs/websocket"
	"github.com/gorilla/websocket"
)

func HandleJoin(ctx context.Context,conn *Client, request map[string]interface{}) {
	conn.Join("my_room")
	conn.WriteMessage(websocket.TextMessage,[]byte("request"))
}

func HandleBroadcastToRoom(ctx context.Context,conn *Client, request map[string]interface{}) {
	message := request["message"].(string)

	BroadcastToGroup("my_room", []byte(message))
	conn.WriteMessage(websocket.TextMessage,[]byte("Broadcaseted"))
}

func HandleLeave(ctx context.Context,conn *Client, request map[string]interface{}) {
	conn.Leave("my_room")
}