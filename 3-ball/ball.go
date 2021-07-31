package main

import (
	"flag"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

const BallRadius = 7.0

type Ball struct {
	vector   *Vector2
	velocity *Vector2
	radius   float64
	speed    float64
	color    uint32
}

var ballSpeed = flag.Int("ball-speed", 4, "Speed of the ball")

func NewBall(x, y float64) *Ball {
	ball := &Ball{
		vector: &Vector2{X: x, Y: y},
		speed:  float64(*ballSpeed),
		radius: BallRadius,
		color:  0xffffff,
	}

	velocity := Vector2{X: float64(1 + rand.Intn(5)), Y: float64(rand.Intn(5))}
	ball.velocity = velocity.Normalize().MultiplyNum(ball.speed)
	return ball
}

func (self *Ball) Rect() *sdl.Rect {
	size := uint16(self.radius * 2)
	x := self.vector.X - self.radius
	y := self.vector.Y - self.radius
	return &sdl.Rect{X: int32(x), Y: int32(y), W: int32(size), H: int32(size)}
}

func (self *Ball) Draw(renderer *sdl.Renderer) {
	err := renderer.FillRect(self.Rect())
	if err != nil {
		return 
	}
}
