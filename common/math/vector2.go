package math

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X int
	Y int
}

func NewVector2(x, y int) Vector2 {
	return Vector2{
		X: x,
		Y: y,
	}
}

func (v Vector2) Distance(other Vector2) int {
	return int(math.Sqrt(math.Pow(float64(other.X-v.X), 2) + math.Pow(float64(other.Y-v.Y), 2)))
}

func (v Vector2) DistanceManathan(other Vector2) int {
	return int(math.Abs(float64(other.X-v.X)) + math.Abs(float64(other.Y-v.Y)))
}

func (v Vector2) Direction(other Vector2) Vector2 {
	v.X -= other.X
	v.Y -= other.Y
	return v.Normalize()
}

func (v Vector2) Normalize() Vector2 {
	m := v.Magnitude()
	if m > 0 {
		v.X /= m
		v.Y /= m
	}
	return v
}

func (v Vector2) Magnitude() int {
	return int(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v Vector2) Add(other Vector2) Vector2 {
	v.X += other.X
	v.Y += other.Y
	return v
}

func (v Vector2) Minus(other Vector2) Vector2 {
	v.X -= other.X
	v.Y -= other.Y
	return v
}

func (v Vector2) Equal(other Vector2) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v Vector2) ToString() string {
	return fmt.Sprintf("%d:%d", v.X, v.Y)
}
