package game

import (
	"os"
)

func (g *Game)ReadInput() (string, error) {
	buffer := make([]byte, 128)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}
	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	}
	return "", nil
}
