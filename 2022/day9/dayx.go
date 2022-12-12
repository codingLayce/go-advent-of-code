package day9

import (
	"fmt"
	gomath "math"
	"strconv"
	"strings"

	"adc/common/math"
)

type Day9 struct {
	Input string
	Dir   string
}

func New() Day9 {
	return Day9{
		Input: "2022/day9/input.txt",
		Dir:   "2022/day9/",
	}
}

var movements = map[string]math.Vector2{
	"R": {X: 1, Y: 0},
	"L": {X: -1, Y: 0},
	"U": {X: 0, Y: -1},
	"D": {X: 0, Y: 1},
}

func (d Day9) ProcessPuzzle1(lines []string) (string, error) {
	fmt.Println(5 / 2)
	return fmt.Sprintf("%d", process(lines, 2)), nil
}

func (d Day9) ProcessPuzzle2(lines []string) (string, error) {
	return fmt.Sprintf("%d", process(lines, 10)), nil
}

func process(lines []string, size int) int {
	visited := make(map[string]struct{})
	var rope []math.Vector2
	for i := 0; i < size; i++ {
		rope = append(rope, math.NewVector2(0, 0))
	}
	visited[rope[size-1].ToString()] = struct{}{}

	for _, line := range lines {
		direction, steps := parseLine(line)
		var vis []math.Vector2
		rope, vis = move(rope, direction, steps, size-1)
		for _, e := range vis {
			visited[e.ToString()] = struct{}{}
		}
	}
	return len(visited)
}

func move(rope []math.Vector2, direction string, steps int, tailIdx int) ([]math.Vector2, []math.Vector2) {
	var visited []math.Vector2
	for i := 0; i < steps; i++ {
		rope[0] = rope[0].Add(movements[direction])

		for idx := 1; idx < len(rope); idx++ {
			head := rope[idx-1]
			tail := rope[idx]
			delta := math.NewVector2(int(gomath.Abs(float64(head.X-tail.X))), int(gomath.Abs(float64(head.Y-tail.Y))))
			if delta.X == 2 || delta.Y == 2 { // knots aren't touching anymore
				if delta.X == 2 { // head: 5 tail: 3 delta: 2
					rope[idx].X = (tail.X + head.X) / 2 // 5+3=8/2=4 (tail is not touching head)
				} else { // move forward
					rope[idx].X = head.X
				}
				if delta.Y == 2 {
					rope[idx].Y = (tail.Y + head.Y) / 2
				} else {
					rope[idx].Y = head.Y
				}
			}
			if tailIdx == idx {
				visited = append(visited, rope[idx])
			}
		}
	}
	return rope, visited
}

func parseLine(line string) (string, int) {
	arr := strings.Split(line, " ")
	value, _ := strconv.Atoi(arr[1])
	return arr[0], value
}
