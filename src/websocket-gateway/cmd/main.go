package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"syscall"

	websocket "websocket-gateway/internal/websocket"
	logger "websocket-gateway/pkg/logger"
)



func main() {
	// Increase resources limitations
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	// Enable pprof hooks
	go func() {
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Fatalf("pprof failed: %v", err)
		}
	}()


	go websocket.Start()
	logger.Info("main.go", "Starting gateway on 8000")

	http.HandleFunc("/", websocket.WsHandler)
	if err := http.ListenAndServe("0.0.0.0:8000", nil); err != nil {
		log.Fatal(err)
	}
}
