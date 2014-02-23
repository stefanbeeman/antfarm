package af

import (
	"errors"
	"fmt"
	"github.com/mgutz/ansi"
)

type Logger interface {
	log(string)
	dev(string)
	meh(string)
	success(string)
	worked(string)
	info(string)
	warning(string)
	problem(string)
	danger(string)
	broke(string)
	spooky(string)
	fun(string)
	test()
}

type BasicLogger struct{}

func (this BasicLogger) log(log string) {
	out := ansi.ColorCode("white") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) dev(log string) {
	out := ansi.ColorCode("white+bu") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) meh(log string) {
	out := ansi.ColorCode("gray") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) success(log string) {
	out := ansi.ColorCode("green") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) worked(log string) {
	out := ansi.ColorCode("green+bu") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) info(log string) {
	out := ansi.ColorCode("blue") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) warning(log string) {
	out := ansi.ColorCode("yellow") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) problem(log string) {
	out := ansi.ColorCode("yellow+bu") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) danger(log string) {
	out := ansi.ColorCode("red") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) broke(log error) {
	out := ansi.ColorCode("red+bu") + log.Error() + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) spooky(log string) {
	out := ansi.ColorCode("magenta") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) important(log string) {
	out := ansi.ColorCode("cyan") + log + ansi.ColorCode("reset")
	fmt.Println(out)
}

func (this BasicLogger) test() {
	this.log("Log")
	this.dev("Dev")
	this.meh("Meh")
	this.success("Success")
	this.worked("Worked")
	this.info("Info")
	this.warning("Warning")
	this.problem("Problem")
	this.danger("Danger")
	err := errors.New("Broke")
	this.broke(err)
	this.spooky("Spooky")
	this.important("Important")
}

var console = BasicLogger{}
