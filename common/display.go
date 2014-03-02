package common

type Display interface {
	Glyph() string
}

type Displayable interface {
	Display() Display
}
