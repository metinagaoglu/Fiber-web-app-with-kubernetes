package websocket

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"fmt"
	"net/http"
	"context"

	epoll "websocket-gateway/internal/epoll"
	handlers "websocket-gateway/internal/handlers"
	pre_connection "websocket-gateway/internal/middleware/pre_connection"
)

func WsHandler(w http.ResponseWriter, r *http.Request) {
	epoller := epoll.GetEpollInstance()

	/*
	*	Pre connection
	*/
	pre_connection := pre_connection.Run(r)
	fmt.Println("pre", pre_connection)
	if !pre_connection {
		return
	}

	// Upgrade connection
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		return
	}

	if err := epoller.Add(conn); err != nil {
		log.Printf("Failed to add connection %v", err)
		conn.Close()
	}
}


func Start() {
	epoller := epoll.GetEpollInstance()
	ctx := context.Background()

	for {
		connections, err := epoller.Wait()
		if err != nil {
			log.Printf("Failed to epoll wait %v", err)
			continue
		}
		for _, conn := range connections {
			if conn == nil {
				break
			}
			if msg, _, err := wsutil.ReadClientData(conn); err != nil {
				if err := epoller.Remove(conn); err != nil {
					log.Printf("Failed to remove %v", err)
				}
				conn.Close()
			} else {
				// This is commented out since in demo usage, stdout is showing messages sent from > 1M connections at very high rate
					log.Printf("msg: %s", string(msg))

					handlers.Run(&conn, ctx, string(msg))
					err = wsutil.WriteServerMessage(conn, 1, msg)
					if err != nil {
						// handle error
					}
			}
		}
	}
}
