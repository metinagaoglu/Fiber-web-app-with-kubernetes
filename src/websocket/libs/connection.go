package libs

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

// WebsocketConn is a wrapper around websocket.Conn
type WebsocketConn struct {
	*websocket.Conn
}

var connections = make([]*WebsocketConn, 0) // slice of connections

func (conn *WebsocketConn) BroadcastExceptOne(message JsonRPCResponse) {
	for _, client := range connections {
		if client != conn {
			err := client.WriteJSON(message)
			if err != nil {
				log.Printf("error sending message to client: %v", err)
				conn.RemoveConnection()
			}
		}
	}
}

func (conn *WebsocketConn) AddConnection() {
	connections = append(connections, conn)
	fmt.Println("Client count is now: ", len(connections))

}

func (conn *WebsocketConn) RemoveConnection() {
	for i, client := range connections {
		if client == conn {
			connections = append(connections[:i], connections[i+1:]...)
			break
		}
	}
}

func (conn *WebsocketConn) Close() error {
	conn.RemoveConnection()
	return conn.Conn.Close()
}
