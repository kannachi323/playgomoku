package server

type GameState struct {
    Board [][]*Stone   `json:"board"`
    Size    int     `json:"size"`
    Players []Player `json:"players"`
    Turn    string  `json:"turn"`
    Status  string  `json:"status"`
}


type Stone struct {
    R     int       `json:"r"`
    C     int       `json:"c"`
    Color string    `json:"color"`
}

type Player struct {
    ID          string    `json:"id"`
    Username    string    `json:"username"`
}
