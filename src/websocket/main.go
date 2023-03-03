package main

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	. "go-apps-with-kubernetes/libs/websocket"
	. "go-apps-with-kubernetes/routes"

	amqpclient "go-apps-with-kubernetes/connections"
)

// Upgrader to upgrade incoming HTTP requests to websocket connections
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading HTTP connection to websocket:", err)
		return
	}
	wsConn := &Client{Conn: conn, Rooms: make(map[string]bool)}
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
			go HandleRPCRequest(ctx, &wg, wsConn, request)
		}
	}
	wg.Wait()
}

func main() {

	// Connection to RabbitMQ
	conn := amqpclient.Connection()

	conn.Close()

	http.HandleFunc("/websocket", websocketHandler)
	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
