package players

import (
	"strings"

	"github.com/FeiNiaoBF/Pac_Go/pkg/game"
)

// Player struct
// This is the player for the PacMan game
type Player struct {
	Sprite sprite
}

// sprite struct
// tracking the player's position
type sprite struct {
	Row int
	Col int
}

// NewPlayer
// Create a new player
func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) FindPlayer(maze []string) {
	// find the player in the maze
	for row, line := range maze {
		// Only one player
		col := strings.Index(line, "P")
		p.Sprite = sprite{row, col}
		break
	}
}

func (p *Player) MovePlayer(g *game.Game, dir string) {
	p.Sprite.Row, p.Sprite.Col = g.Move(p.Sprite.Row, p.Sprite.Col, dir)
}
