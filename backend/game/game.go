package game

import (
	"sync"

	"github.com/google/uuid"
)
type Game struct {
    GameManager *GameManager
    GameID string
}

type GameManager struct {
    Status  Status
    Result  Result
    Board   *Board
    PlayerTurn int
    Moves []Move
    mu sync.Mutex
}

func CreateGame(boardSize int) *Game {
    gameID := uuid.New().String()

    gameManager := &GameManager{
        Status: Active,
        Result: Pending,
        Board: CreateBoard(boardSize),
        PlayerTurn: 1,
        Moves: make([]Move, 0),
    }

    game := &Game{
        GameManager: gameManager,
        GameID: gameID,
    }

    return game
}
