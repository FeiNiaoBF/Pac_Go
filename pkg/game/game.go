package game

import (
	"errors"
	"log"
	"os"
	"os/exec"


)

type Game struct {
	Maze   []string
}

// NewGame
func NewGame() *Game {
	return &Game{}
}

// ===============
// Error messages
// ===============
var ErrNotFile = errors.New("could find not a file")

// initialize game
func (g *Game) InitialiseGame() {
	// Check if the terminal supports cbreak mode
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalln("unable to activate cbreak mode:", err)
	}
}

// restore terminal to cooked mode
func (g *Game) CleanupGame() {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalln("unable to restore cooked mode:", err)
	}
}

