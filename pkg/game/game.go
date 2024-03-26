package game

import "errors"

type Game struct {
	Maze []string
}

// NewGame
func NewGame() *Game {
	return &Game{}
}

// ===============
// Error messages
// ===============
var ErrNotFile = errors.New("could find not a file")

