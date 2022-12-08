package day8

import (
	"fmt"
	"strconv"

	"adc/common/slices"
	"adc/common/strings"
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

func parseBoard(lines []string) ([][]int, int, int) {
	width := len(lines[0])
	height := len(lines)
	var board [][]int

	for row := 0; row < height; row++ {
		board = append(board, []int{})
		for col := 0; col < width; col++ {
			value, _ := strconv.Atoi(string(lines[row][col]))
			board[row] = append(board[row], value)
		}
	}

	return board, width, height
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
	maxScenic := 0
	board, width, height := parseBoard(lines)

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if score := scenicScore(board, row, col, width, height); score > maxScenic {
				maxScenic = score
			}
		}
	}

	return fmt.Sprintf("%d", maxScenic), nil
}

func scenicScoreForDirection(max int, view []int) int {
	for idx, tree := range view {
		if tree >= max {
			return idx + 1
		}
	}
	return len(view)
}

func scenicScore(board [][]int, row, col, width, height int) int {
	score := 1
	tree := board[row][col]

	if col != width-1 {
		score *= scenicScoreForDirection(tree, board[row][col+1:])
	}

	if col != 0 {
		score *= scenicScoreForDirection(tree, slices.Reverse(board[row][:col]))
	}

	if row != height-1 {
		view := 0
		for _, line := range board[row+1:] {
			view++
			if tree >= line[col] {
				break
			}
		}
		score *= view
	}

	if row != 0 {
		view := 0
		for _, line := range slices.Reverse(board[:row]) {
			view++
			if tree >= line[col] {
				break
			}
		}
		score *= view
	}

	return score
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
	for i := row - 1; i >= 0; i-- {
		value, _ := strconv.Atoi(string(lines[i][col]))
		toTop = append(toTop, value)
	}
	if _, topMax := slices.Max(toTop); topMax < tree {
		return 1
	}

	// tree --> bottom
	var toBottom []int
	for i := row + 1; i < height; i++ {
		value, _ := strconv.Atoi(string(lines[i][col]))
		toBottom = append(toBottom, value)
	}
	if _, bottomMax := slices.Max(toBottom); bottomMax < tree {
		return 1
	}

	return 0
}
