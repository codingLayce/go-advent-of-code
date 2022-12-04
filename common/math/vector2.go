package math

import "fmt"

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

func (v Vector2) Equal(other Vector2) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v Vector2) ToString() string {
	return fmt.Sprintf("%s:%s", v.X, v.Y)
}
