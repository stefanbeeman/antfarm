package common

type Display interface{}

type Displayable interface {
	Display() Display
}
