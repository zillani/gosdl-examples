package main

import (
	"flag"
	"github.com/veandco/go-sdl2/sdl"
)

const paddleWidth = 8
const paddleHeight = 80

type World struct {
	running       bool
	pause         bool
	Height, Width int
	Screen        *sdl.Surface
	Ball          *Ball
	Paddle        *Paddle
}

type Paddle struct {
	vector               *Vector2
	target               *Vector2
	height, width, speed float64
	color                uint32
}

var paddleSpeed  = flag.Int("player-speed", 4, "Speed of player paddle")

func NewPaddle(x, y, w, h float64) *Paddle {
	return &Paddle{
		vector: &Vector2{X: x, Y: y},
		target: &Vector2{X: x, Y: y},
		height: h,
		width:  w,
		color:  0x6666ff,
		speed:  float64(*paddleSpeed),
	}
}

func (self *Paddle) Go(x, y float64) {
	self.target = &Vector2{X: self.vector.X, Y: y}
}

func (self *Paddle) Update(world *World) {
	goal := self.target.Minus(self.vector)

	if goal.Length() > self.speed {
		goal = goal.Normalize().MultiplyNum(self.speed)
	}

	future := self.vector.Plus(goal)
	if future.Y < (self.height / 2) {
		return
	}
	if (future.Y + (self.height / 2)) > float64(world.Height) {
		return
	}
	self.vector = future
}

func (self *Paddle) Draw(renderer *sdl.Renderer) {
	err := renderer.FillRect(self.Rect())
	if err != nil {
		return
	}
}

func (self *Paddle) Rect() *sdl.Rect {
	h := self.height
	w := self.width
	x := self.vector.X - w/2
	y := self.vector.Y - h/2
	return &sdl.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)}
}
func (self *Paddle) Hit(past, future *Vector2) (hit bool, place *Vector2) {
	// our front line
	halfHeight := self.height / 2
	halfWidth := self.width / 2
	x0, y0 := (self.vector.X + halfWidth), (self.vector.Y - halfHeight)
	x1, y1 := (self.vector.X + halfWidth), (self.vector.Y + halfHeight)

	return self.hitCore(x0, y0, x1, y1, past, future)
}
func (self *Paddle) hitCore(x0, y0, x1, y1 float64, past, future *Vector2) (hit bool, place *Vector2) {
	// line between past and future
	x2, y2 := past.X, past.Y
	x3, y3 := future.X, future.Y
	d := (x1-x0)*(y3-y2) - (y1-y0)*(x3-x2)

	if d < 0.001 {
		return
	} // never hit since parallel

	ab := ((y0-y2)*(x3-x2) - (x0-x2)*(y3-y2)) / d

	if ab > 0.0 && ab < 1.0 {
		cd := ((y0-y2)*(x1-x0) - (x0-x2)*(y1-y0)) / d
		if cd > 0.0 && cd < 1.0 {
			linx := x0 + ab*(x1-x0)
			liny := y0 + ab*(y1-y0)
			hit = true
			place = &Vector2{X: linx, Y: liny}
		}
	}

	// no hit
	return
}
