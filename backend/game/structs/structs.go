package structs

type Result int
const (
    Win Result = iota
    Draw
    Pending 
)
type Status int
const (
    Active Status = iota
    Offline
)

type Color int
const (
    _ Color = iota
    White
    Black
)

type Move struct {
    R int
    C int
    Player Player
}

type Player struct {
    ID    int
    Name  string
    Color Color
}

type Room struct {
    RoomID int
    Players [2]Player
    Status string

}
