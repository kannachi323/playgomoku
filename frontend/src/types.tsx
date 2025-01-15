export interface Conn {
  roomID: string;
  player: Player
  type: string
}

export interface Player {
    playerID: string
    color: number
}
