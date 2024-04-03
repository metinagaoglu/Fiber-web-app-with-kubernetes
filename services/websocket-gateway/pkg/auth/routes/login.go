package routes

import (
	"context"
	"fmt"
	"github.com/gobwas/ws/wsutil"
	"net"

	epoll "websocket-gateway/internal/epoll"
	auth "websocket-gateway/pkg/auth"
	pb "websocket-gateway/pkg/auth/pb"
)

type AuthHandler struct{}

type AuthRequestBody struct {
	Token string `json:"token"`
}

func (h *AuthHandler) HandleMessage(conn *net.Conn, ctx context.Context, route string, payload string) {
	// "foo" mesajı işleme kodu burada
	b := AuthRequestBody{}

	// if err := ctx.BindJSON(&b); err != nil {
	//     wsutil.WriteServerMessage(*conn, 1, []byte("Login error"))
	//     return
	// }

	client := auth.InitServiceClient()
	res, err := client.Validate(context.Background(), &pb.ValidateRequest{
		Token: b.Token,
	})

	if err != nil {
		fmt.Println(err)
		//ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	fmt.Println("id:", epoll.GetIdFromConn(*conn))
	fmt.Println("payload:", payload)
	fmt.Println("res:", res)
	ctx.Done()
	wsutil.WriteServerMessage(*conn, 1, []byte("test"))
}
