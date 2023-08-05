package websocket

import(
    "net/http"
    "fmt"

    authMiddleware "websocket-gateway/pkg/auth/middleware"
)


// secureHeadersMiddleware adds two basic security headers to each HTTP response
func Auth(r *http.Request) (int64, error) {
	token := r.Header.Get("Authorization")
  if len(token) == 0 {
    err :=  fmt.Errorf("No token provided")
    return 0 , err
  }
  return authMiddleware.AuthByToken(token)
}


