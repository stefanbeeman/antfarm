package antfarm

const (
	SIMPLE  int = 1500
	COMPLEX int = 3000
)

type Actor interface {
	tic(*World)
}

type Action struct {
	delay    int
	complete func()
}

func makeWaitAction(duration int) Action {
	return Action{
		duration,
		func() {},
	}
}
