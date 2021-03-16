package terminal

import (
	"fmt"
	"strconv"
)

func CursorUp(n int) {
	fmt.Print("\x1b[" + strconv.Itoa(n) + "A")
}

func CursorDown(n int) {
	fmt.Print("\x1b[" + strconv.Itoa(n) + "B")
}

func CursorRight(n int) {
	fmt.Print("\x1b[" + strconv.Itoa(n) + "C")
}
func CursorLeft(n int) {
	fmt.Print("\x1b[" + strconv.Itoa(n) + "D")
}
