package tests

import (
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
	. "go-apps-with-kubernetes/libs/websocket"
)

func TestBroadcast(t *testing.T) {

	// Create a websocket client using gorilla package
	ws := NewClient("ws://localhost:8080/websocket")

	// Connect to the websocket server
	err := ws.Connect()
	if err != nil {
		t.Fatal(err)
	}

	// Close the connection when the test is done
	defer ws.Close()

	// Wait a message from the server
	_, message, err := ws.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}

}
