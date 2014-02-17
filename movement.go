package antfarm

import (
  "math/rand"
  "fmt"
)

var (
	NORTH   = Point{0, 1}
	SOUTH   = Point{0, -1}
	EAST    = Point{1, 0}
	WEST    = Point{-1, 0}
	NOWHERE = Point{0, 0}
)

type Mover interface {
	moveTo(*BasicUnit, Point) (func() Action)
}

type AStarWalker struct {}

func (this AStarWalker) moveTo(u *BasicUnit, d Point) (func() Action) {
  closedSet := make(map[Point]Point)
  q := MakePositionHeap()

  q.Push(u.Position, 0)
  closedSet[u.Position] = u.Position
  done := false

  for !done && q.Size() > 0 {
    c, dist := q.Pop()
    if c == d {
      done = true
    } else {
      for _, n := range getNEWS(c) {
        cost, moveable := u.moveCost(n)
        if _, exists := closedSet[n]; !exists && moveable {
          weight := h(d,n) + dist + cost
          q.Push(n, weight)
          closedSet[n] = c
        }
      }
    }
  }

  if !done {
    // No possible path - deal with it later
    fmt.Println("unable to make path - bad things")
    return nil
  }

  b := d
  result := []Point{}
  for b != u.Position {
    result = append(result, b)
    b = closedSet[b]
  }
  l := len(result)

  return func() Action {
    return &BasicAction{
      1,
      func() {
        u.Position = result[l-1]
        l += -1
      },
    }
  }
}

func h(p1, p2 Point) float64 {
  return p1.distanceTo(p2)
}

func getNEWS(p Point) [4]Point {
  return [4]Point{Point{p.X+1,p.Y}, Point{p.X-1,p.Y}, Point{p.X,p.Y+1}, Point{p.X,p.Y-1}}
}

type RandomWalker struct {
	Speed int
}

func (this RandomWalker) move(u *BasicUnit, vector Point) {
	u.Position = u.Position.add(vector)
}

func (this RandomWalker) moveTo(u *BasicUnit, p Point) (func() Action) {
	return func() Action {
    act := BasicAction{
      1,
      func() {
        i := rand.Intn(4)
        u.Position = getNEWS(u.Position)[i]
      },
    }
    return &act
	}
}
