package websocket

import (
	"fmt"
	"net/http"

	middleware "websocket-gateway/pkg/auth/middleware"
)

// secureHeadersMiddleware adds two basic security headers to each HTTP response
func Auth(r *http.Request) (int64, error) {
	//return 1, nil
	var token string
	token = r.Header.Get("Authorization")

	// If no token provided, we can get auth header from query string
	if token == "" {
		token = r.URL.Query().Get("token")
	}

	if len(token) == 0 {
		err := fmt.Errorf("No token provided")
		return 0, err
	}
	return middleware.AuthByToken(token)
}
