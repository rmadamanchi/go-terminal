package main

import (
	"fmt"
	"strconv"

	"github.com/rmadamanchi/go-terminal/terminal"
)

func main() {
	p := terminal.NewPrinter()
	p.WithColor(terminal.Black, "black ").
		WithColor(terminal.Red, "red ").
		WithColor(terminal.Green, "green ").
		WithColor(terminal.Yellow, "yellow ").
		WithColor(terminal.Magenta, "magenta ").
		WithColor(terminal.Cyan, "cyan ").
		WithColor(terminal.White, "white ")
	fmt.Println()
	fmt.Println()

	p.WithColorBold(terminal.Black, "bold-black ").
		WithColorBold(terminal.Red, "bold-red ").
		WithColorBold(terminal.Green, "bold-green ").
		WithColorBold(terminal.Yellow, "bold-yellow ").
		WithColorBold(terminal.Magenta, "bold-magenta ").
		WithColorBold(terminal.Cyan, "bold-cyan ").
		WithColorBold(terminal.White, "bold-white ")
	fmt.Println()
	fmt.Println()

	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			code := i*16 + j
			p.WithColor256(code, strconv.Itoa(code)).Print(" ")
		}
	}

	fmt.Println()
	fmt.Println()

	p.WithBgColor(terminal.Black, "bg-black ").
		WithBgColor(terminal.Red, "bg-red ").
		WithBgColor(terminal.Green, "bg-green ").
		WithBgColor(terminal.Yellow, "bg-yellow ").
		WithBgColor(terminal.Magenta, "bg-magenta ").
		WithBgColor(terminal.Cyan, "bg-cyan ").
		WithBgColor(terminal.White, "bg-white ")
	fmt.Println()
	fmt.Println()

	p.WithBgColorBold(terminal.Black, "bg-bold-black ").
		WithBgColorBold(terminal.Red, "bg-bold-red ").
		WithBgColorBold(terminal.Green, "bg-bold-green ").
		WithBgColorBold(terminal.Yellow, "bg-bold-yellow ").
		WithBgColorBold(terminal.Magenta, "bg-bold-magenta ").
		WithBgColorBold(terminal.Cyan, "bg-bold-cyan ").
		WithBgColorBold(terminal.White, "bg-bold-white ")
	fmt.Println()
	fmt.Println()

	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			code := i*16 + j
			p.WithBgColor256(code, strconv.Itoa(code)).Print(" ")
		}
	}

	fmt.Println()
	fmt.Println()

	p.WithBold("bold ")
	p.WithUnderline("underline ")
	p.WithReversed("reversed ")

	fmt.Println()
	fmt.Println()

	p.StartBgColor(terminal.Cyan).StartColor(terminal.Red).StartBold().StartUnderline().Print("cyan-bg-red-fg-bold-underline").Reset()
}
