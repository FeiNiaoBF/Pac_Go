package test_test

import (
	"reflect"
	"testing"
)

var test1String = []string{
	"################",
	"#G#..#   # ##G.#",
	"#.#... #.#    .#",
	"#.###.# .# # #.#",
	"#.    # .#    .#",
	"#.### # .#    .#",
	"#.#...# .   # .#",
	"#.#...#..#..#..#",
	"#.##X###### ## #",
	"#..........P#XG#",
	"################",
}

func TestLoad(t *testing.T) {
	t.Run("game load-test1", func(t *testing.T) {
		game := game.NewGame()
		game.Load("data/maze/maze01.txt")
		got := game.Maze
		want := test1String
		assertString(t, got, want)
	})

	t.Run("game load-test2 is error", func(t *testing.T) {
		g := game.NewGame()
		got := g.Load("data/maze/maze.txt")
		want := game.ErrNotFile
		assertError(t, got, want)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertString(t testing.TB, got, want []string) {
	t.Helper()
	if reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
