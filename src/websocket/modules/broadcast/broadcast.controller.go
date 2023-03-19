package broadcast

import (
	"context"
	"fmt"
	"encoding/json"

	. "go-apps-with-kubernetes/libs/websocket"
	"github.com/gorilla/websocket"
)

func HandleJoin(ctx context.Context,conn *Client, request map[string]interface{}) {
	conn.Join("my_room")
	conn.WriteMessage(websocket.TextMessage,[]byte("request"))
}

func HandleBroadcastToRoom(ctx context.Context,conn *Client, request map[string]interface{}) {
		message := request["message"].(string)

		response := JsonRPCResponse{
			Result:  json.RawMessage(fmt.Sprintf(`{"message": "%s", "sender":"%s"}`, message, "TODO: auth")),
			ID:      "123",
		}

		BroadcastJsonToGroup("my_room", response)
}

func HandleLeave(ctx context.Context,conn *Client, request map[string]interface{}) {
	conn.Leave("my_room")
}

func HandlePing(ctx context.Context,conn *Client, request map[string]interface{}) {
	conn.WriteMessage(websocket.TextMessage,[]byte("pong"))
}