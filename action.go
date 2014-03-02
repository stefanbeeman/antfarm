package af

import "fmt"

const (
	SIMPLE  int = 15
	COMPLEX int = 30
)

type Action interface {
	complete() bool
	tic()
}

type BasicAction struct {
	warmUp int
	coolDown int
	action func()
}

func (this *BasicAction) tic() {
	if this.warmUp > 0 {
		this.warmUp--
	} else if this.warmUp == 0 {
		this.action()
	} else {
		this.coolDown--
	}
}

func (this *BasicAction) complete() bool { return this.coolDown <= 0 }

func MakeMoveAction(u Unit, l Location, time int) Action {
	fmt.Println("Going to... ", l)
	return &BasicAction{0, time, func(){u.SetPosition(l)} }
}

func MakeWaitAction(i int) Action {
	fmt.Println("Waiting")
	return &BasicAction{0, i, func(){}}
}
