package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"playgomoku/backend/game"
	"playgomoku/backend/game/structs"
	"playgomoku/backend/middleware"

	"github.com/go-chi/chi/v5"
)

func GetGame(w http.ResponseWriter, r *http.Request) {
    gameID := chi.URLParam(r, "gameID")
    fmt.Printf("Received request for gameID: %s\n", gameID)

    game, exists := game.GetGame(gameID)
    if !exists {
        http.Error(w, "Game not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(game)
}

func AddGame(w http.ResponseWriter, r *http.Request) {
	var gameData struct {
		P1 structs.Player `json:"p1"`
		P2 structs.Player `json:"p2"`
	}

	if err := json.NewDecoder(r.Body).Decode(&gameData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	gameID, success := game.AddGame(15, &gameData.P1, &gameData.P2)
	if !success {
		http.Error(w, "Error creating game", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"gameID": gameID})
}


func GetNumGames(w http.ResponseWriter, r *http.Request) {
	numGames := game.GetNumGames()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"numGames": numGames})
}

func GetPlayerTurn(w http.ResponseWriter, r *http.Request) {
	game := r.Context().Value(middleware.GameKey("game")).(*game.Game).GameManager
	playerTurn := game.PlayerTurn

	json.NewEncoder(w).Encode(playerTurn)
}

func NewGameRoutes(r chi.Router) {
	r.Get("/game/{gameID}", GetGame)
	r.Get("/game/num-games", GetNumGames)
	r.Post("/game", AddGame)
}

