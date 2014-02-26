package af

import (
	"fmt"
)

const (
  MAX_WEIGHT int = 1
)

var DESIRES = map[Point]float64 {
  Point{1,1}: float64(2),
  Point{18,18}: float64(2),
  Point{1,18}: float64(1),
  Point{18,1}: float64(1),
}

type Thinker interface {
	think(*BasicUnit, World) Task
}

type BasicThinker struct{}

func (this BasicThinker) think(u *BasicUnit, w World) Task {
	task := NO_TASK
	for task.complete() {
    desiresSlice := []DesirePoint{}
    for k,v := range DESIRES {
      desiresSlice = append(desiresSlice, DesirePoint{k, v})
    }
		fmt.Println("I'm heading to ", desiresSlice)
		task = BasicTask{
			func() bool {
        for _, d := range desiresSlice {
          if u.Position.equals(d.Position) {
            delete(DESIRES, d.Position)
            return true
          }
        }
				return false
			},
			u.mover.moveTo(u, desiresSlice),
		}
	}

	return task
}
