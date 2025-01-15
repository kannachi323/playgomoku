package server

import (
	"fmt"
	"playgomoku/backend/game"
	"sync"

	"golang.org/x/net/websocket"
)

type Room struct {
	roomID  string
	conns   map[*websocket.Conn]bool // Store all client connections
	game    *game.Game
	mu      sync.Mutex
	players map[*game.Player]*websocket.Conn
	
}

func (rm *Room) addConnection(ws *websocket.Conn, player *game.Player) {
    rm.mu.Lock()
    defer rm.mu.Unlock()

    if rm.conns == nil || rm.players == nil {
        ws.Write([]byte("error: broken resources"))
	}

    rm.conns[ws] = true
    rm.players[player] = ws

	rm.mu.Lock()
	ws.Write([]byte(fmt.Sprintf("Player %s joined room %s\n", player.PlayerID, rm.roomID)))
	rm.mu.Unlock()
}


func (rm *Room) removeConnection(ws *websocket.Conn) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	// Remove the WebSocket connection
	delete(rm.conns, ws)

	// Close the WebSocket
	_ = ws.Close()

	fmt.Printf("Connection removed from room %s\n", rm.roomID)
}

func (rm *Room) broadcast(message []byte) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	for ws := range rm.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(message); err != nil {
				fmt.Printf("Broadcast error: %v\n", err)
			}
		}(ws)
	}
}

func (rm *Room) startGame() {
    rm.game = game.CreateGame(15)
    rm.broadcast([]byte("Game has started!"))
}



