package main

import (
	"net/http"

	"playgomoku/backend/server"
)

func main() {
    s := server.CreateServer()
	s.MountHandlers()
	
	http.ListenAndServe(":3000", s.Router)
}
