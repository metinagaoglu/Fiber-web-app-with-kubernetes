package websocket

import (
	"github.com/gorilla/websocket"
)

var rooms = make(map[string][]*Client)

func (client *Client) Join(groupName string) {
	if _, ok := rooms[groupName]; !ok {
		rooms[groupName] = []*Client{}
	}

	if client.Rooms[groupName] {
		return
	}

	rooms[groupName] = append(rooms[groupName], client)
	client.Rooms[groupName] = true
}

func (client *Client) Leave(groupName string) {
	if clients, ok := rooms[groupName]; ok {
		for i, c := range clients {
			if c == client {
				rooms[groupName] = append(clients[:i], clients[i+1:]...)
				delete(client.Rooms, groupName)
				break
			}
		}
	}
}

func BroadcastToGroup(groupName string, msg []byte) {
	if clients, ok := rooms[groupName]; ok {
		for _, c := range clients {
			c.Conn.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
