package main

import (
	"fmt"
	"log"

	"github.com/FeiNiaoBF/Pac_Go/pkg/game"
)

const (
	filePath = "data/maze/maze01.txt"
)

func main() {
	game := game.NewGame()
	// initialize game

	// load game resources
	err := game.Load(filePath)
	if err != nil {
		log.Printf("Error loading maze file: \n%v", err)
		return
	}
	// game loop
	for {
		// update screen
		fmt.Print(game.ToString())
		// process input

		// process movement

		// process collisions

		// check game over
		fmt.Println("Hello, Pac Go!")
		break

		// repeat
	}
}
