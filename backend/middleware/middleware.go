package middleware

import (
	"context"
	"fmt"
	"net/http"
	"playgomoku/backend/game"

	"github.com/go-chi/chi/v5"
)

type GameKey string

func JSONMiddleware(next http.Handler) http.Handler {
	fmt.Println("got to json middleware first")
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}

func GameMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gameID := chi.URLParam(r, "gameID")
		game, exists := game.GetGame(gameID)
		if !exists {
			http.Error(w, "Game not found", http.StatusNotFound)
			return
		}
		
		gameKey := GameKey("game")
		ctx := context.WithValue(r.Context(), gameKey, game)


		fmt.Println("Setting game in context for gameID:", gameID)

		retrievedGame := ctx.Value(gameKey)
		if retrievedGame != nil {
			fmt.Printf("Game object set in context: %+v\n", retrievedGame)
		} else {
			fmt.Println("No game object found in context")
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}