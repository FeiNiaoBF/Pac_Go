package game

import (
	"bufio"
	"os"
)

// loadMaze is load maze file to game
// return maze and error
func (g *Game) loadMaze(file string) ([]string, error) {
	var maze []string
	f, err := os.Open(file)
	if err != nil {
		return nil, ErrNotFile
	}
	// defer puts f.Close() in the *call stack*
	defer f.Close()
	// buffered I/O
	scanner := bufio.NewScanner(f)
	// reading data such as a file of newline-delimited lines of text
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}
	return maze, nil
}

// Load is load maze file to game
// return error
func (g *Game) Load(file string) error {
	maze, err := g.loadMaze(file)
	if err != nil {
		return err
	}
	g.Maze = maze
	// find player
	return nil
}
