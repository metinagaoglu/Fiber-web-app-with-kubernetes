package routes

import (
	"context"
	"log"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
	. "go-apps-with-kubernetes/libs/websocket"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func HandleRPCRequest(ctx context.Context, wg *sync.WaitGroup, conn *Client, request JsonRPCRequest) {
	defer wg.Done()

	handler := GetHandlerByType(request.Method)
	params := request.GetParams()

	if handler.Handler == nil {
		log.Println("Error decoding JSON-RPC params:")
		conn.WriteMessage(websocket.TextMessage, []byte("Error decoding JSON-RPC params"))
		return
	}
	log.Println("[HANDLER]: ", request.Method)

	validate = validator.New()

	if handler.Rules != nil {
		errs := validate.ValidateMap(params, handler.Rules)
		if len(errs) > 0 {
			log.Println(errs)
			return
		}
	}

	handler.Handler(ctx, conn, params)
}
