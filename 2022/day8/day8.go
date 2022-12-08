package day8

import (
	"adc/common/slices"
	"adc/common/strings"
	"fmt"
	"strconv"
)

type Day8 struct {
	Input string
	Dir   string
}

func New() Day8 {
	return Day8{
		Input: "2022/day8/input.txt",
		Dir:   "2022/day8/",
	}
}

func (d Day8) ProcessPuzzle1(lines []string) (string, error) {
	total := 0
	width := len(lines[0])
	height := len(lines)

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			total += isVisible(lines, row, col, width, height)
		}
	}

	return fmt.Sprintf("%d", total), nil
}

func (d Day8) ProcessPuzzle2(lines []string) (string, error) {
	return "not implemented", nil
}

func isVisible(lines []string, row, col, width, height int) int {
	line := strings.ToIntSlice(lines[row])
	tree := line[col]

	// Edges
	if row == 0 || col == 0 || row == height-1 || col == width-1 {
		return 1
	}

	// tree --> right
	if _, rightMax := slices.Max(line[col+1:]); rightMax < tree {
		return 1
	}

	// tree --> left
	if _, leftMax := slices.Max(line[:col]); leftMax < tree {
		return 1
	}

	// tree --> top
	var toTop []int
	for i := row-1;i >=0; i-- {
		value, _ := strconv.Atoi(string(lines[i][col]))
		toTop = append(toTop, value)
	}
	if _, topMax := slices.Max(toTop); topMax < tree {
		return 1
	}

	// tree --> bottom
	var toBottom []int
	for i := row+1;i < height; i++ {
		value, _ := strconv.Atoi(string(lines[i][col]))
		toBottom = append(toBottom, value)
	}
	if _, bottomMax := slices.Max(toBottom); bottomMax < tree {
		return 1
	}

	return 0
}
