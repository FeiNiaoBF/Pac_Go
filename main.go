package main

import (
	"fmt"
	"log"

	"github.com/FeiNiaoBF/Pac_Go/pkg/game"
)

const (
	mazeFilePath = "data/maze/maze01.txt"
)

func main() {
	game := game.NewGame()
	// initialize game
	game.InitialiseGame()
	defer game.CleanupGame()
	// load game resources
	err := game.Load(mazeFilePath)
	if err != nil {
		log.Printf("Error loading maze file: \n%v", err)
		return
	}
	// game loop
	for {
		// update screen
		game.ToString()
		// process input
		input, err := game.ReadInput()
		if err != nil {
			log.Printf("error reading input: %q", err)
			break
		}
		if input == "ESC" {
			break
		}
		// process movement

		// process collisions

		// check game over
		fmt.Println("Hello, Pac Go!")
		break

		// repeat
	}
}
