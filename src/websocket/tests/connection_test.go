package tests

import (
	"net/http"
	"testing"

	. "go-apps-with-kubernetes/libs/websocket"
)

func TestConnection(t *testing.T) {
	e := WsHandlerTester(t)

	jsonRPC, _ := GenerateRequest("ping", map[string]interface{}{"message": "hi"}, GenerateID())

	ws := e.GET("/websocket").WithWebsocketUpgrade().
		Expect().
		Status(http.StatusSwitchingProtocols).
		Websocket()
	defer ws.Disconnect()

	ws.WriteText(string(jsonRPC)).
		Expect().
		TextMessage().Body().IsEqual("pong")

}
