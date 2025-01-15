package server

import "playgomoku/backend/game"

type PlayerData struct {
    PlayerID  string `json:"id"`
	Color game.Color `json:"color"`
}

type ConnData struct {
    RoomID  string  `json:"roomID"`
    Player *game.Player  `json:"player"`
    Type string `json:"type"`
}

