package common

import (
	"fmt"
	"github.com/mgutz/ansi"
)

type Logger interface {
	Log(string)
	Dev(string)
	Meh(string)
	Success(string)
	Worked(string)
	Info(string)
	Warning(string)
	Problem(string)
	Danger(string)
	Broke(string)
	Spooky(string)
	Fun(string)
}

type BasicLogger struct{}

func (this BasicLogger) Log(log string) {
	out := ansi.ColorCode("white") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) Dev(log string) {
	out := ansi.ColorCode("white+bu") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) Meh(log string) {
	out := ansi.ColorCode("gray") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) Success(log string) {
	out := ansi.ColorCode("green") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) Worked(log string) {
	out := ansi.ColorCode("green+bu") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) Info(log string) {
	out := ansi.ColorCode("blue") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) Warning(log string) {
	out := ansi.ColorCode("yellow") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) Problem(log string) {
	out := ansi.ColorCode("yellow+bu") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) Danger(log string) {
	out := ansi.ColorCode("red") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) Broke(log error) {
	out := ansi.ColorCode("red+bu") + log.Error() + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) Spooky(log string) {
	out := ansi.ColorCode("magenta") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) Important(log string) {
	out := ansi.ColorCode("cyan") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

var Console = BasicLogger{}
