package server

import (
	"fmt"
	"io"
	"playgomoku/backend/game"
	"strconv"
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
	defaultRooms := make(map[string]*Room)

	defaultRooms["sexy"] = &Room{
		roomID: "sexy",
		conns: make(map[*websocket.Conn]bool),
		game: nil,
		players: make(map[*game.Player]*websocket.Conn),
	}

	return &Server{
		rooms: defaultRooms,
		startTime: time.Now(),
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
    ws.Write([]byte(fmt.Sprintf("Incoming WebSocket connection from: %s\n", ws.RemoteAddr())))

    room := s.rooms["sexy"]
    if room == nil {
        fmt.Println("error: room not found")
        return
    }

    params := ws.Request().URL.Query()
    playerID := params.Get("pid")
    if playerID == "" {
        fmt.Println("error: player ID not found")
        return
    }

    clr := params.Get("clr")
    color, err := strconv.Atoi(clr)
    if err != nil {
        fmt.Println("error: invalid color")
        ws.Close()
        return
    }

    newPlayer := &game.Player{
        PlayerID: playerID,
        Color:    game.Color(color),
    }

    room.addConnection(ws, newPlayer)
    
	fmt.Println("Player joined room:", room.roomID)

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