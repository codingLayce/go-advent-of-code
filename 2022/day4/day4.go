package day4

import (
	"fmt"
	"strconv"
	"strings"

	"adc/common/math"
)

type Day4 struct {
	Input string
	Dir   string
}

func New() Day4 {
	return Day4{
		Input: "2022/day4/input.txt",
		Dir:   "2022/day4/",
	}
}

func (d Day4) ProcessPuzzle1(lines []string) (string, error) {
	total := 0

	for _, line := range lines {
		a, b := extractPairs(line)
		if isInside(b, a.X) && isInside(b, a.Y) || isInside(a, b.X) && isInside(a, b.Y) {
			total++
		}
	}

	return fmt.Sprintf("%d", total), nil
}

func (d Day4) ProcessPuzzle2(lines []string) (string, error) {
	total := 0

	for _, line := range lines {
		a, b := extractPairs(line)
		if isInside(a, b.X) || isInside(a, b.Y) || isInside(b, a.X) || isInside(b, a.Y) {
			total++
		}
	}

	return fmt.Sprintf("%d", total), nil
}

func isInside(segment math.Vector2, point int) bool {
	return point >= segment.X && point <= segment.Y
}

func extractPairs(line string) (math.Vector2, math.Vector2) {
	arr := strings.Split(line, ",")
	res := [2]math.Vector2{}

	for i, e := range arr {
		r := strings.Split(e, "-")
		x, _ := strconv.Atoi(r[0])
		y, _ := strconv.Atoi(r[1])
		res[i] = math.NewVector2(x, y)
	}

	return res[0], res[1]
}
