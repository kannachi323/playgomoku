package game

import (
	"fmt"
	"playgomoku/backend/game/board"
	"playgomoku/backend/game/structs"
	"sync"

	"github.com/google/uuid"
)

var (
    games = make(map[string]*Game)
    mu    sync.Mutex
    numGames int
)

type Game struct {
    GameManager *GameManager
}

type GameManager struct {
    Status  structs.Status
    Result  structs.Result
    Board   *board.Board
    Player1 structs.Player
    Player2 structs.Player
    PlayerTurn structs.Player
    Moves []structs.Move
}

func CreateGameManager(p1 *structs.Player, p2 *structs.Player, boardSize int) *GameManager {
    return &GameManager{
        Status: structs.Active,
        Result: structs.Pending,
        Board: board.CreateBoard(boardSize),
        Player1: *p1,
        Player2: *p2,
        PlayerTurn: *p1,
        Moves: make([]structs.Move, 0),
    }
}

func AddGame(boardSize int, p1 *structs.Player, p2 *structs.Player) (string, bool) {
    mu.Lock()
    defer mu.Unlock()


    gameID := uuid.New().String()

    gameManager := CreateGameManager(p1, p2, boardSize);

    newGame := &Game{
        GameManager: gameManager,
    }

    games[gameID] = newGame

    numGames += 1

    return gameID, true
}

func GetGame(gameID string) (*Game, bool) {
    mu.Lock()
    defer mu.Unlock()
    game, exists := games[gameID]
    if !exists {
        fmt.Printf("Game with ID %s not found\n", gameID)
    }
    return game, exists
}

func GetNumGames() int {
    return numGames
}

func (gm *GameManager) SwitchPlayers() bool {
    id := gm.PlayerTurn.ID
    
    if id == gm.Player1.ID {
        gm.PlayerTurn = gm.Player2
    } else if id == gm.Player2.ID {
        gm.PlayerTurn = gm.Player1
    }

    return gm.PlayerTurn.ID != id //make sure we switched players
}