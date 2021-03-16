package main

import (
	"fmt"
	"strconv"

	ansi "gitbuh.com/rmadamanchi/go-terminal/terminal"
)

func main() {
	p := ansi.NewPrinter()
	p.WithColor(ansi.Black, "black ").
		WithColor(ansi.Red, "red ").
		WithColor(ansi.Green, "green ").
		WithColor(ansi.Yellow, "yellow ").
		WithColor(ansi.Magenta, "magenta ").
		WithColor(ansi.Cyan, "cyan ").
		WithColor(ansi.White, "white ")
	fmt.Println()
	fmt.Println()

	p.WithColorBold(ansi.Black, "bold-black ").
		WithColorBold(ansi.Red, "bold-red ").
		WithColorBold(ansi.Green, "bold-green ").
		WithColorBold(ansi.Yellow, "bold-yellow ").
		WithColorBold(ansi.Magenta, "bold-magenta ").
		WithColorBold(ansi.Cyan, "bold-cyan ").
		WithColorBold(ansi.White, "bold-white ")
	fmt.Println()
	fmt.Println()

	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			code := i*16 + j
			p.WithColor256(code, strconv.Itoa(code)).Print(" ")
		}
	}
}
