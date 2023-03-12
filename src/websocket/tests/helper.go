package tests

import (
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/google/uuid"
)

func WsHandlerTester(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "ws://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

func GenerateID() string {
	return uuid.New().String()
}
