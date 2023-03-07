package routes

import (
	"context"

	. "go-apps-with-kubernetes/libs/websocket"

	. "go-apps-with-kubernetes/modules/broadcast"
	. "go-apps-with-kubernetes/modules/broadcast/validations"
)

type MessageHandler func(ctx context.Context, conn *Client, request map[string]interface{})

type RequestHandler struct {
	Handler MessageHandler
	Rules   map[string]interface{}
}

var handlers = make(map[string]RequestHandler)

func GetHandlerByType(messageType string) RequestHandler {
	handlers["broadcast"] = RequestHandler{
		Handler: HandleBroadcastToRoom,
		Rules:   GetBroadcastValidationRules(),
	}
	handlers["join"] = RequestHandler{
		Handler: HandleJoin,
		Rules:   nil,
	}
	handlers["leave"] = RequestHandler{
		Handler: HandleLeave,
		Rules:   nil,
	}
	return handlers[messageType]
}
