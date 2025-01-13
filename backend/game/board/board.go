package board

import (
	"playgomoku/backend/game/structs"
)

type Board struct {
	Size int
	Grid [][]structs.Color
}

func CreateBoard(size int) *Board {
	grid := make([][]structs.Color, size) 
	for i := range grid {
		grid[i] = make([]structs.Color, size)
	}
	
	return &Board{
		Size: size, 
		Grid: grid,
	}
}

func (board *Board) PlaceStone(i int, j int, color structs.Color) bool {
	if !board.IsValidMove(i, j) || board.Grid[i][j] != 0 {
		return false;
	}
	board.Grid[i][j] = color
	return board.Grid[i][j] == color
	
}

func (board *Board) RemoveStone(i int, j int) bool {
	board.Grid[i][j] = 0
	return board.Grid[i][j] == 0
}

func(board *Board) ClearStones() {
	for i := range board.Grid {
		for j := range board.Grid[i] {
			board.Grid[i][j] = 0
		}
	}
}

func (board *Board) CheckWin(move structs.Move) bool {
	r, c := move.R, move.C
	color := move.Player.Color
	directions := [8][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	
	

	var dfs func(row int, col int, dr int, dc int, seq int) bool
	dfs = func (row int, col int, dr int, dc int, seq int) bool {
		if !board.IsValidMove(row, col) || board.Grid[row][col] != color {
			return false
		}
		if seq == 5 {
			return true
		}
	
		return dfs(row + dr, col + dc, dr, dc, seq + 1)
	}

	for _, direction := range directions {
		dr, dc := direction[0], direction[1]
		if dfs(r, c, dr, dc, 1) {
			return true;
		}
	}

	return false
	
}

func (board *Board) IsValidMove(r int, c int) bool {
	if r < 0 || r >= board.Size || c < 0 || c >= board.Size {
		return false
	}
	return true
}
