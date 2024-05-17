package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/danicat/simpleansi"
)

// Config
type Config struct {
	Player       string        `json:"player,omitempty"`
	Ghost        string        `json:"ghost,omitempty"`
	Wall         string        `json:"wall,omitempty"`
	Dot          string        `json:"dot,omitempty"`
	Pill         string        `json:"pill,omitempty"`
	Death        string        `json:"death,omitempty"`
	Space        string        `json:"space,omitempty"`
	GhostBlue    string        `json:"ghost_blue,omitempty"`
	PillDuration time.Duration `json:"pill_duration,omitempty"`
	UseEmoji     bool          `json:"use_emoji,omitempty"`
}

var cfg Config

func loadConfig(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	return nil
}

// Sprite is a struct that represents a sprite in the game
type sprite struct {
	row      int
	col      int
	startRow int
	startCol int
}

var (
	maze    []string
	player  sprite
	ghosts  []*ghost
	score   int
	numDots int
)

type GhostStatus string

const (
	GhostStatusNormal GhostStatus = "Normal"
	GhostStatusBlue   GhostStatus = "Blue"
)

type ghost struct {
	position sprite
	status   GhostStatus
}

var lives = 1

// File Path
var fileMaze = "data/maze/maze01.txt"

// config flag
var (
	configFile = flag.String("config-file", "config.json", "path to custom configuration file")
	mazeFile   = flag.String("maze-file", fileMaze, "path to a custom maze file")
)

func loadMaze(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'P':
				player = sprite{row, col, row, col}
			case 'G':
				ghosts = append(ghosts, &ghost{sprite{row, col, row, col}, GhostStatusNormal})
			case '.':
				numDots++
			}
		}
	}

	return nil
}

func readInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	} else if cnt >= 3 {
		if buffer[0] == 0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil
			}
		}
	}

	return "", nil
}

func getLivesAsEmoji() string {
	buf := bytes.Buffer{}
	for i := lives; i > 0; i-- {
		buf.WriteString(cfg.Player)
	}
	return buf.String()
}

func printScreen() {
	simpleansi.ClearScreen()
	for _, line := range maze {
		for _, chr := range line {
			switch chr {
			case '#':
				fmt.Print(simpleansi.WithBlueBackground(cfg.Wall))
			case '.':
				fmt.Print(cfg.Dot)
			default:
				fmt.Print(cfg.Space)
			}
		}
		fmt.Println()
	}

	moveCursor(player.row, player.col)
	fmt.Print(cfg.Player)

	for _, g := range ghosts {
		moveCursor(g.position.row, g.position.col)
		if g.status == GhostStatusNormal {
			fmt.Printf(cfg.Ghost)
		} else if g.status == GhostStatusBlue {
			fmt.Printf(cfg.GhostBlue)
		}
	}
	moveCursor(len(maze)+1, 0)

	livesRemaining := strconv.Itoa(lives)
	if cfg.UseEmoji {
		livesRemaining = getLivesAsEmoji()
	}
	fmt.Printf("%vDots: %v\t %vGhosts: %v\r\n", cfg.Dot, numDots, cfg.Ghost, len(ghosts))
	fmt.Println("Score:", score, "\tLives:", livesRemaining)
}

func makeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol

	switch dir {
	case "UP":
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(maze) - 1
		}
	case "DOWN":
		newRow = newRow + 1
		if newRow == len(maze) {
			newRow = 0
		}
	case "RIGHT":
		newCol = newCol + 1
		if newCol == len(maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(maze[0]) - 1
		}
	}

	if maze[newRow][newCol] == '#' {
		newRow = oldRow
		newCol = oldCol
	}

	return
}

func movePlayer(dir string) {
	player.row, player.col = makeMove(player.row, player.col, dir)

	removeDot := func(row, col int) {
		maze[row] = maze[row][0:col] + " " + maze[row][col+1:]
	}

	switch maze[player.row][player.col] {
	case '.':
		numDots--
		score++
		removeDot(player.row, player.col)
	case 'X':
		score += 10
		removeDot(player.row, player.col)
		go processPill()
	}
}

var (
	ghostsStatusMx sync.RWMutex
	pillMx         sync.Mutex
	pillTimer      *time.Timer
)

func updateGhosts(ghosts []*ghost, ghostStatus GhostStatus) {
	ghostsStatusMx.Lock()
	defer ghostsStatusMx.Unlock()
	for _, g := range ghosts {
		g.status = ghostStatus
	}
}

func processPill() {
	pillMx.Lock()
	updateGhosts(ghosts, GhostStatusBlue)
	for _, g := range ghosts {
		g.status = GhostStatusBlue
	}
	pillTimer = time.NewTimer(time.Second * cfg.PillDuration)
	pillMx.Unlock()
	<-pillTimer.C
	pillMx.Lock()
	for _, g := range ghosts {
		g.status = GhostStatusNormal
	}
	updateGhosts(ghosts, GhostStatusNormal)
	pillMx.Unlock()
}

func moveGhosts() {
	for _, g := range ghosts {
		dir := drawDirection()
		g.position.row, g.position.col = makeMove(g.position.row, g.position.col, dir)
	}
}

func initialise() {
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalln("unable to activate cbreak mode:", err)
	}
}

func cleanup() {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalln("unable to activate cooked mode:", err)
	}
}

func drawDirection() string {
	dir := rand.Intn(4)
	move := map[int]string{
		0: "UP",
		1: "DOWN",
		2: "RIGHT",
		3: "LEFT",
	}
	return move[dir]
}

func moveCursor(row, col int) {
	if cfg.UseEmoji {
		simpleansi.MoveCursor(row, col*2)
	} else {
		simpleansi.MoveCursor(row, col)
	}
}

func main() {
	flag.Parse()
	// initialise game
	initialise()
	defer cleanup()

	// load resources
	err := loadMaze(*mazeFile)
	if err != nil {
		log.Println("failed to load maze:", err)
		return
	}
	// load config
	err = loadConfig(*configFile)
	if err != nil {
		log.Println("failed to load configuration:", err)
		return
	}
	// process input
	input := make(chan string, 128)
	go func(ch chan<- string) {
		for {
			str, err := readInput()
			if err != nil {
				log.Println("failed to read input:", err)
				ch <- "ESC"
			}
			ch <- str
		}
	}(input)
	// game loop
	for {
		// update screen
		printScreen()
		// process movement
		select {
		case inp := <-input:
			if inp == "ESC" {
				lives = 0
			}
			movePlayer(inp)
		default:
		}
		moveGhosts()
		time.Sleep(200 * time.Millisecond)
		// process collisions
		for _, g := range ghosts {
			if player.row == g.position.row && player.col == g.position.col {
				ghostsStatusMx.RLock()
				if g.status == GhostStatusNormal {
					lives -= 1
					if lives != 0 {
						moveCursor(player.row, player.row)
						fmt.Print(cfg.Death)
						moveCursor(len(maze)+2, 0)
						ghostsStatusMx.RUnlock()
						updateGhosts(ghosts, GhostStatusNormal)
						// resetting player address
						time.Sleep(1 * time.Second)
						player.row, player.col = player.startRow, player.startCol
					}
				} else if g.status == GhostStatusBlue {
					ghostsStatusMx.RUnlock()
					updateGhosts([]*ghost{g}, GhostStatusNormal)
					g.position.row, g.position.col = g.position.startRow, g.position.startCol
				}
			}
		}
		printScreen()
		// check game over
		if numDots == 0 || lives == 0 {
			if lives == 0 {
				moveCursor(player.row, player.col)
				fmt.Print(cfg.Death)
				moveCursor(len(maze)+2, 0)
			}
			break
		}

		// repeat
		time.Sleep(100 * time.Millisecond)
	}
}
