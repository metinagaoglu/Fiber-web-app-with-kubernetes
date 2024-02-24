package websocket

import (
	"fmt"
	"log"
	"strconv"
	epoll "websocket-gateway/internal/epoll"
	wsutil "github.com/gobwas/ws/wsutil"
)

/* This function consume rabbitmq  */
func HandleQueueMessage(payload string) {
	fmt.Println("P", payload)

	connId, _ := strconv.Atoi(payload)

	connection := epoll.GetConnById(connId)

	if connection == nil {
		log.Printf("Failed to get connection %v", connId)
		return
	}

	wsutil.WriteServerMessage(connection, 1, []byte("From another service"))
}
