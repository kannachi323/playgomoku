export interface Conn {
  roomID: string;
  player: Player
  type: string
}

export interface Player {
    id: string
    name: string
    color: Number
}
