package af

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

func (this BasicAction) tic() {
	if this.warmUp > 0 {
		this.warmUp--
	} else if this.warmUp == 0 {
		this.action()
	} else {
		this.coolDown--
	}
}

func (this BasicAction) complete() bool { return this.coolDown <= 0 }

func MakeWaitAction(i int) Action {
	return BasicAction{0, i, func(){}}
}
