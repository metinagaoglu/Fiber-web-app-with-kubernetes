package handlers

import (
	"context"
	"fmt"
	"net"

	"encoding/json"
	authroutes "websocket-gateway/pkg/auth/routes"
	sessionroutes "websocket-gateway/pkg/session/routes"

	"github.com/go-playground/validator/v10"
)

type Message struct {
	Route   *string `json:"route" validate:"required" default:""`
	Payload *string `json:"payload",json`
}

type MessageHandler interface {
	HandleMessage(conn *net.Conn, ctx context.Context, route string, message string)
}

func Run(connPtr *net.Conn, ctx context.Context, message []byte) {
	// Parse and dispatch message
	var Validator = validator.New()
	var decodedMessage Message

	err := json.Unmarshal([]byte(message), &decodedMessage)
  if err != nil {
		conn := *connPtr
		conn.Close()
		return
  }

	err = Validator.Struct(decodedMessage)
	if err != nil {
		fmt.Println(err.Error())
		conn := *connPtr
		conn.Close()
		return
	}

	if decodedMessage.Payload == nil {
		decodedMessage.Payload = new(string)
	}

	// fmt.Println("Message route ->", *decodedMessage.Route)
	// fmt.Println("Message payload ->", *decodedMessage.Payload)
	//TODO: Implement validation and middleware layers here

	handlers := make(map[string]MessageHandler)
	handlers["auth"] = &authroutes.AuthHandler{}
	handlers["init-session"] = &sessionroutes.InitSessionHandler{}

	if handler, ok := handlers[string(*decodedMessage.Route)]; ok {
		go handler.HandleMessage(connPtr, ctx, *decodedMessage.Route, *decodedMessage.Payload)
	}
}
