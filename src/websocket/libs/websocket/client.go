package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

// Client is a wrapper around websocket.Conn
type Client struct {
	*websocket.Conn
}

var connections = make([]*Client, 0) // slice of connections

func (conn *Client) BroadcastExceptOne(message JsonRPCResponse) {
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

func (conn *Client) AddConnection() {
	connections = append(connections, conn)
	fmt.Println("Client count is now: ", len(connections))

}

func (conn *Client) RemoveConnection() {
	for i, client := range connections {
		if client == conn {
			connections = append(connections[:i], connections[i+1:]...)
			break
		}
	}
}

func (conn *Client) Close() error {
	conn.RemoveConnection()
	return conn.Conn.Close()
}
