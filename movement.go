package af

type Mover interface {
	MovementAlg
}

type BasicMover struct {
	MovementAlg
}