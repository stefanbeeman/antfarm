package antfarm

type Task struct {
	complete func() bool
	next     func() Action
}
