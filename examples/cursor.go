package main

import (
	"fmt"
	"time"

	"github.com/rmadamanchi/go-terminal/terminal"
)

func main() {

	for i := 0; i < 100; i++ {

		fmt.Print(i)
		fmt.Println("%")
		fmt.Print(i + 1)
		fmt.Print("%")

		time.Sleep(100 * time.Millisecond)

		if i != 99 {
			terminal.CursorUp(1)
			terminal.CursorLeft(100)
		}
	}
}
