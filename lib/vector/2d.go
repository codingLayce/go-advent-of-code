package vector

import "math"

var (
	LeftDir   = NewVec2(0, -1)
	RightDir  = NewVec2(0, 1)
	TopDir    = NewVec2(-1, 0)
	BottomDir = NewVec2(1, 0)
)

type Vec2 struct {
	X int
	Y int
}

func NewVec2(x, y int) Vec2 {
	return Vec2{X: x, Y: y}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	v.X -= other.X
	v.Y -= other.Y
	return v
}

func (v Vec2) Add(other Vec2) Vec2 {
	v.X += other.X
	v.Y += other.Y
	return v
}

func (v Vec2) Minus(other Vec2) Vec2 {
	v.X -= other.X
	v.Y -= other.Y
	return v
}

func (v Vec2) Mul(val int) Vec2 {
	v.X *= val
	v.Y *= val
	return v
}

func (v Vec2) Magnitude() int {
	return int(math.Sqrt(float64(v.X)*float64(v.X) + float64(v.Y)*float64(v.Y)))
}
