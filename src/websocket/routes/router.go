package routes

import (
	"context"
	"log"
	"sync"

	. "go-apps-with-kubernetes/libs/websocket"
)

func HandleRPCRequest(ctx context.Context, wg *sync.WaitGroup, conn *Client, request JsonRPCRequest) {
	defer wg.Done()

	handler := GetHandlerByType(request.Method)
	if handler == nil {
		log.Println("Error decoding JSON-RPC params:")
		return
	}
	log.Println("[HANDLER]: ", request.Method)
	handler(ctx, conn, request)
}
