package math

import "fmt"

type Vector2 struct {
	x int
	y int
}

func NewVector2(x, y int) Vector2 {
	return Vector2{
		x: x,
		y: y,
	}
}

func (v Vector2) Equal(other Vector2) bool {
	return v.x == other.x && v.y == other.y
}

func (v Vector2) ToString() string {
	return fmt.Sprintf("%s:%s", v.x, v.y)
}
