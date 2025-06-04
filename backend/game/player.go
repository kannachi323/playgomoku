package game

type Player struct {
    ID          string    `json:"id"`
    Username    string    `json:"username"`
}

func NewPlayers(p1 Player, p2 Player) []*Player {
    newPlayers := make([]*Player, 2)
    newPlayers[0] = &p1
    newPlayers[1] = &p2

    return newPlayers
}