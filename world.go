package antfarm

import (
	"strconv"
	"time"
)

type World interface {
	tic()
	Start()
	Sleep(int)
	Stop()
	RunFor(int)
	addActor(Actor)
}

type BasicWorld struct {
	Grid2D
	Now       int
	Materials []Material
	Actors    []Actor
	pacemaker *time.Ticker
}

func (this *BasicWorld) tic() {
	this.Now++
	for _, actor := range this.Actors {
		actor.tic(this)
	}
}

func (this *BasicWorld) Start() {
	for t := range this.pacemaker.C {
		_ = t
		this.tic()
	}
}

func (this *BasicWorld) Sleep(tics int) {
	duration, _ := time.ParseDuration(strconv.Itoa(tics*100) + "ms")
	time.Sleep(duration)
}

func (this *BasicWorld) Stop() {
	this.pacemaker.Stop()
}

func (this *BasicWorld) RunFor(tics int) {
	for i := 0; i < tics; i++ {
		this.tic()
	}
}

func (this *BasicWorld) addActor(a Actor) {
	this.Actors = append(this.Actors, a)
}

func MakeWorld(data string, width int, height int, worms int) World {
	g := makeGrid2D(width, height)
	m := LoadMaterials(data)
	u := make([]Actor, 0)
	pm := time.NewTicker(time.Millisecond * 100)
	w := BasicWorld{g, 0, m, u, pm}

	for y, row := range w.Cells {
		for x, _ := range row {
			p := Point{x, y}
			c := makeCell(p, w.Materials[0], false)
			if x == 0 || y == 0 || x == (width-1) || y == (height-1) {
				c.setSolid(true)
			}
			w.set(p, c)
		}
	}

	for n := 0; n < worms; n++ {
		rp := w.random()
		a := makeWorm(rp)
		w.addActor(a)
	}
	return &w
}
