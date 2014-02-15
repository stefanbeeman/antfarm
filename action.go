package antfarm

type Action struct {
	delay    int
	complete func()
}
