package day

import (
	"bufio"
	"fmt"
	"strconv"

	"advent/lib/vector"
)

func SolvePart1(reader *bufio.Scanner) (string, error) {
	grid := getInput(reader)

	var (
		targets []vector.Vec2
		starts  []vector.Vec2
	)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 9 {
				targets = append(targets, vector.NewVec2(row, col))
			} else if grid[row][col] == 0 {
				starts = append(starts, vector.NewVec2(row, col))
			}
		}
	}

	count := 0
	for _, start := range starts {
		a := 0
		for _, target := range targets {
			if canReach(grid, start, target, 0) {
				a++
			}
		}
		count += a
	}

	return fmt.Sprintf("%d", count), nil
}

func canReach(grid [][]int, pos, target vector.Vec2, cur int) bool {
	if pos.X < 0 || pos.X >= len(grid) || pos.Y < 0 || pos.Y >= len(grid[pos.X]) {
		return false
	}
	if cur == 9 {
		return pos == target
	}
	if grid[pos.X][pos.Y] != cur {
		return false
	}

	if canReach(grid, pos.Add(vector.NewVec2(1, 0)), target, cur+1) { // down
		return true
	}
	if canReach(grid, pos.Minus(vector.NewVec2(1, 0)), target, cur+1) { // up
		return true
	}
	if canReach(grid, pos.Add(vector.NewVec2(0, 1)), target, cur+1) { // right
		return true
	}
	return canReach(grid, pos.Minus(vector.NewVec2(0, 1)), target, cur+1) // left
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	grid := getInput(reader)

	visited := make([][]int, len(grid))
	for row := 0; row < len(grid); row++ {
		visited[row] = make([]int, len(grid[row]))
		for col := 0; col < len(grid[row]); col++ {
			visited[row][col] = 0
		}
	}

	// Visit all cells starting from 9 to 0
	// By doing it this way I know each time the maximum number of paths from a cell to a 9.
	// example with a 3 :
	// 3210  1...  11..  111.  1112
	// 2101  ....  1...  11.1  1141
	// 1232  ..1.  .111  2111  2111
	for i := 9; i >= 0; i-- {
		for row := 0; row < len(grid); row++ {
			for col := 0; col < len(grid[row]); col++ {
				if grid[row][col] == i {
					visited = countPaths(grid, visited, row, col, i)
				}
			}
		}
	}

	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 0 {
				count += visited[row][col]
			}
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func countPaths(grid, visited [][]int, row, col, value int) [][]int {
	if value == 9 {
		visited[row][col] = 1
		return visited
	}

	count := 0
	if row-1 >= 0 && grid[row-1][col] == value+1 {
		count += visited[row-1][col]
	}
	if row+1 < len(grid) && grid[row+1][col] == value+1 {
		count += visited[row+1][col]
	}
	if col-1 >= 0 && grid[row][col-1] == value+1 {
		count += visited[row][col-1]
	}
	if col+1 < len(grid[row]) && grid[row][col+1] == value+1 {
		count += visited[row][col+1]
	}
	visited[row][col] = count
	return visited
}

func getInput(reader *bufio.Scanner) [][]int {
	var res [][]int
	for reader.Scan() {
		line := make([]int, len(reader.Text()))
		for i, ch := range reader.Text() {
			nb, _ := strconv.Atoi(string(ch))
			line[i] = nb
		}
		res = append(res, line)
	}
	return res
}
