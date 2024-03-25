package game

import (
	"bufio"
	"os"
)

// load the maze file
var maze []string

// loadMaze is load maze file to game
func (game *Game) loadMaze(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
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
	return nil
}

// Load
func (game *Game) Load(file string) error {
	err := game.loadMaze(file)
	if err != nil {
		return err
	}
	game.Maze = maze
	return nil
}

// ToString
func (game *Game) ToString() string {
	var result string
	for _, line := range game.Maze {
		result += line + "\n"
	}
	return result
}
