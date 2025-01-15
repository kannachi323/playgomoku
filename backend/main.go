package main

import (
	"net/http"

	"playgomoku/backend/server"

	"golang.org/x/net/websocket"
)
func main() {
	server := server.NewServer()

	http.Handle("/ws", websocket.Handler(server.HandleWS))
	http.ListenAndServe(":3000", nil)
}
