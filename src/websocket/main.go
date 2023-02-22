package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	. "go-apps-with-kubernetes/libs"
)

// declaring a struct
type Message struct {
	// defining struct variables
	Chat string
}

// Upgrader to upgrade incoming HTTP requests to websocket connections
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// handleRPCRequest is a helper function to handle incoming JSON-RPC requests
func handleRPCRequest(ctx context.Context, wg *sync.WaitGroup, conn *WebsocketConn, request JsonRPCRequest) {
	defer wg.Done()

	//TODO: better way to handle this
	switch request.Method {
	case "message":
		var params map[string]interface{}
		if err := json.Unmarshal(request.Params, &params); err != nil {
			log.Println("Error decoding JSON-RPC params:", err)
			return
		}

		var data Message
		json.Unmarshal(request.Params, &data)

		// Do something with the decoded params
		conn.BroadcastExceptOne(ResponseBuilder(request.ID, json.RawMessage(fmt.Sprintf(`{"message": "%s"}`, data.Chat))))

	case "sleep":
		time.Sleep(8 * time.Second)
		response := JsonRPCResponse{
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
		response := JsonRPCResponse{
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
	wsConn := &WebsocketConn{Conn: conn}
	defer wsConn.RemoveConnection()
	wsConn.AddConnection()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	for {
		select {
		case <-ctx.Done():
			return
		default:
			var request JsonRPCRequest
			if err := conn.ReadJSON(&request); err != nil {
				log.Println("Error reading JSON-RPC request:", err)
				return
			}
			wg.Add(1)
			go handleRPCRequest(ctx, &wg, wsConn, request)
		}
	}
	wg.Wait()
}

func main() {
	http.HandleFunc("/websocket", websocketHandler)
	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
