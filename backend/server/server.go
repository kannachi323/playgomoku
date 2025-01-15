package server

import (
	"fmt"
	"io"
	"playgomoku/backend/game"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	rooms map[string]*Room
	mu    sync.Mutex
	startTime time.Time
}

func NewServer() *Server {
	return &Server{
		rooms: make(map[string]*Room),
		startTime: time.Now(),
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
	fmt.Println(ws.RemoteAddr(), " just connected")

	var data ConnData
    if err := websocket.JSON.Receive(ws, &data); err != nil {
        fmt.Println("Error receiving JSON data:", err)
        return
    }

	s.mu.Lock()
	room, exists := s.rooms[data.RoomID]
	if !exists {
		room = &Room{
			roomID: data.RoomID,
			conns:  make(map[*websocket.Conn]bool),
			game: nil,
			players: make(map[*game.Player]*websocket.Conn),
			
		}
		s.rooms[data.RoomID] = room
	}

	s.mu.Unlock()

	//at this point, we should have the user data rdy from the websocket
	room.addConnection(ws, data.Player)

	
	if len(room.players) == 2 {
		currentPlayers := [2]*game.Player{nil, nil}

		i := 0
		for player, _:= range room.players {
			currentPlayers[i] = player
			i += 1
		}

		fmt.Println("game started!!!: ", currentPlayers[0], "vs", currentPlayers[1])
		room.game = game.CreateGame(15)
	}

	s.readLoop(ws, room)

	room.removeConnection(ws)
}

func (s *Server) readLoop(ws *websocket.Conn, room *Room) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}
		msg := buf[:n]

		room.broadcast(msg)
	}
}