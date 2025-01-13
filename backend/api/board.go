package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"playgomoku/backend/game"
	"playgomoku/backend/middleware"

	"github.com/go-chi/chi/v5"
)

func GetBoard(w http.ResponseWriter, r *http.Request) {
	game := r.Context().Value(middleware.GameKey("game")).(*game.Game).GameManager.Board
	if game == nil {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}


	json.NewEncoder(w).Encode(game)
}

func GetLastMove(w http.ResponseWriter, r *http.Request) {
	moves := r.Context().Value(middleware.GameKey("game")).(*game.Game).GameManager.Moves
	if len(moves) == 0 {
		http.Error(w, "No moves found", http.StatusNotFound)
		return
	}

	lastMove := moves[len(moves)-1]

	json.NewEncoder(w).Encode(lastMove)
}

func AddMove(w http.ResponseWriter, r *http.Request) {
	var moveData struct {
		Row int `json:"row"`
		Col int `json:"col"`
	}

	gm := r.Context().Value(middleware.GameKey("game")).(*game.Game).GameManager
	board, playerTurn := gm.Board, gm.PlayerTurn

	
	if err := json.NewDecoder(r.Body).Decode(&moveData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	row, col, color := moveData.Row, moveData.Col, playerTurn.Color

	
	if !board.PlaceStone(row, col, color) {
        err := fmt.Errorf("invalid move at row: %d, col: %d", row, col)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
    } else if !gm.SwitchPlayers() {
		http.Error(w, "error switching players", http.StatusInternalServerError)
		return
	}


	// return ok status 200
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	
}

func ClearBoard(w http.ResponseWriter, r *http.Request) {
	board := r.Context().Value(middleware.GameKey("game")).(*game.Game).GameManager.Board
	board.ClearStones()

	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func NewBoardRoutes(r chi.Router) {
	r.With(middleware.GameMiddleware).Route("/game/{gameID}/board", func(r chi.Router) {
		r.Get("/", GetBoard)
		r.Delete("/", ClearBoard)
		r.Post("/", AddMove)
		
		r.Get("/last-move", GetLastMove)
	})
}

