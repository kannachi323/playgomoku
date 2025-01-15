package server

import "playgomoku/backend/game"

type PlayerData struct {
    Id  string `json:"id"`
    Name string `json:"name"`
	Color game.Color `json:"color"`
}

type ConnData struct {
    RoomID  string  `json:"roomID"`
    Player *game.Player  `json:"player"`
    Type string `json:"type"`
}

