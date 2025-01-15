package server

import (
	"fmt"
	"io"
	"playgomoku/backend/game"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type Server struct {
	rooms map[string]*Room
	queue []string
	mu    sync.Mutex
	startTime time.Time
}

func NewServer() *Server {
	defaultRooms := make(map[string]*Room)
	defaultQueue := make([]string, 0)

	for i := 0; i < 5; i++ {
		newRoomID := uuid.New().String()
		defaultRooms[newRoomID] = &Room{
			roomID: uuid.New().String(),
			conns: make(map[*websocket.Conn]bool),
			game: nil,
			players: make(map[*game.Player]*websocket.Conn),
		}
		defaultQueue = append(defaultQueue, newRoomID)
	}

	return &Server{
		rooms: defaultRooms,
		queue: defaultQueue,
		startTime: time.Now(),
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
    // Handle incoming WebSocket connection
    s.mu.Lock()
    defer s.mu.Unlock() // Ensures the mutex is always released, even on early return

    ws.Write([]byte(fmt.Sprintf("Incoming WebSocket connection from: %s", ws.RemoteAddr())))

    if len(s.queue) == 0 {
        ws.Write([]byte("error: no available rooms"))
        ws.Close()
        return
    }

    roomID := s.queue[0]
    s.queue = s.queue[1:]
    room := s.rooms[roomID]

    params := ws.Request().URL.Query()

    playerID := params.Get("pid")
    if playerID == "" {
        ws.Write([]byte("error: missing player ID"))
        ws.Close()
        return
    }

    clr := params.Get("clr")
    color, err := strconv.Atoi(clr)
    if err != nil {
        ws.Write([]byte("error: invalid color"))
        ws.Close()
        return
    }

    newPlayer := &game.Player{
        PlayerID: playerID,
        Color:    game.Color(color),
    }

    room.addConnection(ws, newPlayer)

    s.mu.Unlock()

    s.gameLoop(ws, room)
    room.removeConnection(ws)
}


func (s *Server) gameLoop(ws *websocket.Conn, room *Room) {
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