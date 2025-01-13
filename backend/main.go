package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"playgomoku/backend/api"
	"playgomoku/backend/middleware"
)

func main() {
    r := chi.NewRouter()
	r.Use(middleware.JSONMiddleware)


	api.NewGameRoutes(r)
	api.NewBoardRoutes(r)

	http.ListenAndServe(":3000", r)

}