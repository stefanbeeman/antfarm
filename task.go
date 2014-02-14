package antfarm

type Task struct {
	complete func() bool
	invalid  func() bool
	next     func() Action
	reset    func()
}
