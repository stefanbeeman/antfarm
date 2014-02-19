package antfarm

import (
	"fmt"
	"github.com/daviddengcn/go-colortext"
)

type console struct{}

func (this console) log(log string) {
	fmt.Println(log)
}

func (this console) meh(log string) {
	ct.ChangeColor(ct.White, false, ct.None, false)
	fmt.Println(log)
	ct.ResetColor()
}

func (this console) success(log string) {
	ct.ChangeColor(ct.Green, true, ct.None, false)
	fmt.Println(log)
	ct.ResetColor()
}

func (this console) info(log string) {
	ct.ChangeColor(ct.Blue, true, ct.None, false)
	fmt.Println(log)
	ct.ResetColor()
}

func (this console) warning(log string) {
	ct.ChangeColor(ct.Yellow, true, ct.None, false)
	fmt.Println(log)
	ct.ResetColor()
}

func (this console) danger(log string) {
	ct.ChangeColor(ct.Red, true, ct.None, false)
	fmt.Println(log)
	ct.ResetColor()
}

func (this console) spooky(log string) {
	ct.ChangeColor(ct.Magenta, true, ct.None, false)
	fmt.Println(log)
	ct.ResetColor()
}

func (this console) fun(log string) {
	ct.ChangeColor(ct.Cyan, true, ct.None, false)
	fmt.Println(log)
	ct.ResetColor()
}

var afc = new(console)
