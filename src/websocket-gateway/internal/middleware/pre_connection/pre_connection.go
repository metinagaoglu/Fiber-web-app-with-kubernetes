package preconnection

import(
    "net/http"
    "fmt"

    authMiddleware "websocket-gateway/pkg/auth/middleware"
)


// logRequestMiddleware logs basic info of a HTTP request
// RemoteAddr: Network address that sent the request (IP:port)
// Proto: Protocol version
// Method: HTTP method
// URL: Request URL
func logRequestMiddleware(r *http.Request) bool {
    fmt.Printf("LOG %s - %s %s %s\n", r.RemoteAddr, r.Proto, r.Method, r.URL)
    return true    
}

// secureHeadersMiddleware adds two basic security headers to each HTTP response
//TODO: set user
func authHeadersMiddleware(r *http.Request) bool {
	token := r.Header.Get("Authorization")
  if len(token) == 0 {
    //TODO: set false
    return true
  }
  return authMiddleware.AuthByToken(token)
}

// TODO: Run with Array
func Run(r *http.Request) bool {
  logRequestMiddleware(r)

  return authHeadersMiddleware(r)
}

