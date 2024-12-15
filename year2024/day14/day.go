package day

import (
	"advent/lib/rect"
	"advent/lib/vector"
	"bufio"
	"fmt"
	"math"
)

var tilesWidth, tilesHeight int

type Robot struct {
	Position vector.Vec2
	Velocity vector.Vec2
}

func (r *Robot) Move(steps int) {
	r.Position = r.Position.Add(r.Velocity.Mul(steps))

	// Move it in bounds
	r.Position.X %= tilesWidth
	r.Position.Y %= tilesHeight
	if r.Position.X < 0 {
		r.Position.X += tilesWidth
	}
	if r.Position.Y < 0 {
		r.Position.Y += tilesHeight
	}
}

func SolvePart1(reader *bufio.Scanner) (string, error) {
	robots := getInput(reader)
	for _, robot := range robots {
		robot.Move(100)
	}

	return fmt.Sprintf("%d", safety(robots)), nil
}

func safety(robots []*Robot) int {
	width := tilesWidth/2 - 1
	height := tilesHeight/2 - 1

	topLeft := rect.NewRect(vector.NewVec2(0, 0), width, height)
	topRight := rect.NewRect(vector.NewVec2(width+2, 0), width, height)
	bottomLeft := rect.NewRect(vector.NewVec2(0, height+2), width, height)
	bottomRight := rect.NewRect(vector.NewVec2(width+2, height+2), width, height)
	var quadrants [4]int
	for _, robot := range robots {
		if topLeft.IsInside(robot.Position) {
			quadrants[0]++
		} else if topRight.IsInside(robot.Position) {
			quadrants[1]++
		} else if bottomLeft.IsInside(robot.Position) {
			quadrants[2]++
		} else if bottomRight.IsInside(robot.Position) {
			quadrants[3]++
		}
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	robots := getInput(reader)
	maxSafe := math.MaxInt
	sec := 0
	for i := 1; i < tilesWidth*tilesHeight; i++ {
		for _, robot := range robots {
			robot.Move(1)
		}
		saf := safety(robots)
		if saf < maxSafe {
			maxSafe = saf
			sec = i
		}
	}

	return fmt.Sprintf("%d", sec), nil
}

func getInput(reader *bufio.Scanner) []*Robot {
	var robots []*Robot

	for reader.Scan() {
		robot := &Robot{}
		_, _ = fmt.Sscanf(reader.Text(), "p=%d,%d v=%d,%d", &robot.Position.X, &robot.Position.Y, &robot.Velocity.X, &robot.Velocity.Y)
		robots = append(robots, robot)
	}
	return robots
}
