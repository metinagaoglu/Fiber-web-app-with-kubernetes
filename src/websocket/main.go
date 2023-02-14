package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// JSON-RPC request and response structs
type jsonRPCRequest struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	ID     uint64          `json:"id"`
}

type jsonRPCResponse struct {
	ID     uint64          `json:"id"`
	Result json.RawMessage `json:"result"`
	Error  interface{}     `json:"error"`
}

// Upgrader to upgrade incoming HTTP requests to websocket connections
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// handleRPCRequest is a helper function to handle incoming JSON-RPC requests
func handleRPCRequest(ctx context.Context, wg *sync.WaitGroup, conn *websocket.Conn, request jsonRPCRequest) {
	defer wg.Done()

	//TODO: build controllers for this routing
	switch request.Method {
	case "example_method":
		var params map[string]interface{}
		if err := json.Unmarshal(request.Params, &params); err != nil {
			log.Println("Error decoding JSON-RPC params:", err)
			return
		}

		// Do something with the decoded params
		response := jsonRPCResponse{
			ID:     request.ID,
			Result: json.RawMessage(`{"example_result": "Success"}`),
		}
		if err := conn.WriteJSON(response); err != nil {
			log.Println("Error writing JSON-RPC response:", err)
			return
		}
	case "sleep":
		time.Sleep(8 * time.Second)
		response := jsonRPCResponse{
			ID:     request.ID,
			Result: json.RawMessage(`{"sleep": "Success"}`),
		}
		if err := conn.WriteJSON(response); err != nil {
			log.Println("Error writing JSON-RPC response:", err)
			return
		}
	default:
		log.Println("Received unsupported JSON-RPC method:", request.Method)
		// Do something with the decoded params
		response := jsonRPCResponse{
			ID:     request.ID,
			Result: json.RawMessage(`{"example_result": " unsupported JSON-RPC method"}`),
		}
		if err := conn.WriteJSON(response); err != nil {
			log.Println("Error writing JSON-RPC response:", err)
			return
		}
		return
	}
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading HTTP connection to websocket:", err)
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	for {
		select {
		case <-ctx.Done():
			return
		default:
			var request jsonRPCRequest
			if err := conn.ReadJSON(&request); err != nil {
				log.Println("Error reading JSON-RPC request:", err)
				return
			}
			wg.Add(1)
			go handleRPCRequest(ctx, &wg, conn, request)
		}
	}
	wg.Wait()
}

func main() {
	http.HandleFunc("/websocket", websocketHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
