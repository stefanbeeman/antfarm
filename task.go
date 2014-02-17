package antfarm

type Task interface {
	complete() bool
	getAction() Action
}

type BasicTask struct {
	fnComplete func() bool
	fnGet      func() Action
}

func (this BasicTask) complete() bool {
	return this.fnComplete()
}

func (this BasicTask) getAction() Action {
	return this.fnGet()
}
