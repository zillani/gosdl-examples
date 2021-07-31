package main

import "math"

type Vector2 struct {
	X, Y float64
}

func (self *Vector2) Normalize() *Vector2 {
	length := self.Length()
	return &Vector2{X: self.X / length, Y: self.Y / length}
}

func (self *Vector2) MultiplyNum(other float64) *Vector2 {
	return &Vector2{X: self.X * other, Y: self.Y * other}
}

func (self *Vector2) Plus(other *Vector2) *Vector2 {
	return &Vector2{X: self.X + other.X, Y: self.Y + other.Y}
}

func (self *Vector2) PlusEqual(other *Vector2) *Vector2 {
	return &Vector2{
		X: self.X + other.X,
		Y: self.Y + other.Y,
	}
}

func (self *Vector2) Minus(other *Vector2) *Vector2 {
	return &Vector2{X: (self.X - other.X), Y: (self.Y - other.Y)}
}

func (self *Vector2) Length() float64 {
	return math.Sqrt((self.X * self.X) + (self.Y * self.Y))
}
