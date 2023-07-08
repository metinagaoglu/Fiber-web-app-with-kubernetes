package handlers

import(
    "net"
    "context"
    "fmt"
    "encoding/json"

    authRoutes "websocket-gateway/pkg/auth/routes"
    sessionRoutes "websocket-gateway/pkg/session/routes"
)


type Message struct {
    Route string `json:"route"`
    Payload string `json:"payload"`
}

type MessageHandler interface {
    HandleMessage(conn *net.Conn,ctx context.Context, route string, message string)
}

func Run(conn *net.Conn,ctx context.Context, message string) {


    // Parse and dispatch message
    var decodedMessage Message
    err := json.Unmarshal([]byte(message), &decodedMessage)
    if err != nil {
        fmt.Println("Error decoding message")
        return
    }

    fmt.Println("Message route ->", decodedMessage.Route)
    fmt.Println("Message payload ->", decodedMessage.Payload)
    

    handlers := make(map[string]MessageHandler)
    handlers["auth"] = &authRoutes.AuthHandler{}
    handlers["init-session"] = &sessionRoutes.InitSessionHandler{}

    if handler, ok := handlers[string(decodedMessage.Route)]; ok {
        go handler.HandleMessage(conn, ctx, decodedMessage.Route, decodedMessage.Payload)
    }
}