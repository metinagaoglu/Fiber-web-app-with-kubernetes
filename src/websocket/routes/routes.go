package routes

import (
	"context"

	. "go-apps-with-kubernetes/libs/websocket"

	. "go-apps-with-kubernetes/modules/broadcast"
)

type MessageHandler func(ctx context.Context, conn *Client, request JsonRPCRequest)

var handlers = make(map[string]MessageHandler)

func GetHandlerByType(messageType string) MessageHandler {
	handlers["broadcast-room"] = HandleBroadcastToRoom
	handlers["join"] = HandleJoin
	handlers["leave"] = HandleLeave
	return handlers[messageType]
}
