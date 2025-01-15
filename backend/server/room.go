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

	if rm.conns == nil {
		rm.conns = make(map[*websocket.Conn]bool)
	}

	//need to check if player id is already in the room (can't connect > 1 times)
	if _, ok := rm.players[player]; ok {
		fmt.Println("Player already in room")
		return	
	}
	
	rm.conns[ws] = true

	rm.players[player] = ws

	fmt.Printf("Player %s joined room %s\n", player.Name, rm.roomID)

}


func (rm *Room) removeConnection(ws *websocket.Conn) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	// Remove the WebSocket connection
	delete(rm.conns, ws)

	// Close the WebSocket
	_ = ws.Close()

	fmt.Printf("Connection removed from room %s\n", rm.roomID)

	// Check if the room is empty
	if len(rm.conns) == 0 {
		fmt.Printf("Room %s is now empty\n", rm.roomID)
	}
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


