package game

func (g Game) makeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol

	switch dir {
	case "UP":
		newRow = oldRow - 1
		if newRow < 0 {
			newRow = len(g.Maze) - 1
		}
	case "DOWN":
		newRow = newRow + 1
		if newRow == len(g.Maze) {
			newRow = 0
		}
	case "RIGHT":
		newCol = newCol + 1
		if newCol == len(g.Maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(g.Maze[0]) - 1
		}
	}
	return
}

func (g Game) Move(oldRow, oldCol int, dir string) (nRow, nCol int) {
	nRow, nCol = g.makeMove(oldRow, oldCol, dir)
	// if block
	if g.Maze[nRow][nCol] == '#' {
		nRow = oldRow
		nCol = oldCol
	}
	return
}
