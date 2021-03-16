package terminal

import (
	"fmt"
	"strconv"
)

type Color int

const (
	Black   Color = 30
	Red     Color = 31
	Green   Color = 32
	Yellow  Color = 33
	Blue    Color = 34
	Magenta Color = 35
	Cyan    Color = 36
	White   Color = 37
)

type printer struct {
}

func NewPrinter() *printer {
	return &printer{}
}

func (p *printer) WithColor(c Color, s string) *printer {
	return p.StartColor(c).Print(s).Reset()
}

func (p *printer) StartColor(c Color) *printer {
	fmt.Print("\x1b[" + strconv.Itoa(int(c)) + "m")
	return p
}

func (p *printer) WithColorBold(c Color, s string) *printer {
	return p.StartColorBold(c).Print(s).Reset()
}

func (p *printer) StartColorBold(c Color) *printer {
	fmt.Print("\x1b[" + strconv.Itoa(int(c)) + ";1m")
	return p
}

func (p *printer) WithColor256(code int, s string) *printer {
	return p.StartColor256(code).Print(s).Reset()
}

func (p *printer) StartColor256(code int) *printer {
	fmt.Print("\x1b[38;5;" + strconv.Itoa(code) + "m")
	return p
}

func (p *printer) WithBgColor(c Color, s string) *printer {
	return p.StartBgColor(c).Print(s).Reset()
}

func (p *printer) StartBgColor(c Color) *printer {
	fmt.Print("\x1b[" + strconv.Itoa(int(c+10)) + "m")
	return p
}

func (p *printer) WithBgColorBold(c Color, s string) *printer {
	return p.StartBgColorBold(c).Print(s).Reset()
}

func (p *printer) StartBgColorBold(c Color) *printer {
	fmt.Print("\x1b[" + strconv.Itoa(int(c+10)) + ";1m")
	return p
}

func (p *printer) WithBgColor256(code int, s string) *printer {
	return p.StartBgColor256(code).Print(s).Reset()
}

func (p *printer) StartBgColor256(code int) *printer {
	fmt.Print("\x1b[48;5;" + strconv.Itoa(code) + "m")
	return p
}

func (p *printer) WithBold(s string) *printer {
	return p.StartBold().Print(s).Reset()
}

func (p *printer) StartBold() *printer {
	fmt.Print("\x1b[1m")
	return p
}

func (p *printer) WithUnderline(s string) *printer {
	return p.StartUnderline().Print(s).Reset()
}

func (p *printer) StartUnderline() *printer {
	fmt.Print("\x1b[4m")
	return p
}

func (p *printer) WithReversed(s string) *printer {
	return p.StartReversed().Print(s).Reset()
}

func (p *printer) StartReversed() *printer {
	fmt.Print("\x1b[7m")
	return p
}

func (c *printer) Print(s string) *printer {
	fmt.Print(s)
	return c
}

func (c *printer) Reset() *printer {
	fmt.Print("\x1b[0m")
	return c
}
