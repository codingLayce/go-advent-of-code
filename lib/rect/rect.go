package rect

import "advent/lib/vector"

type Rect struct {
	Position vector.Vec2
	Width    int
	Height   int
}

func NewRect(pos vector.Vec2, w, h int) Rect {
	return Rect{
		Position: pos,
		Width:    w,
		Height:   h,
	}
}

func (r Rect) IsInside(point vector.Vec2) bool {
	return point.X >= r.Position.X && point.X <= r.Position.X+r.Width && point.Y >= r.Position.Y && point.Y <= r.Position.Y+r.Height
}
