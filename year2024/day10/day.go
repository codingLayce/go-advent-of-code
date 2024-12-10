package day

import (
	"bufio"
	"fmt"
	"strconv"
)

func SolvePart1(reader *bufio.Scanner) (string, error) {
	grid := getInput(reader)

	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 0 {
				count += countPaths(grid, row, col, -1)
			}
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func countPaths(grid [][]int, row, col int, prev int) int {
	if prev == 9 { // Path reached
		return 1
	}
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) { // Out of bound
		return 0
	}
	if grid[row][col] != prev+1 { // Path stops because wrong value
		return 0
	}
	count := countPaths(grid, row, col+1, prev+1) // Move right
	count += countPaths(grid, row+1, col, prev+1) // Move down
	count += countPaths(grid, row, col-1, prev+1) // Move left
	count += countPaths(grid, row-1, col, prev+1) // Move up

	return count
}

func SolvePart2(reader *bufio.Scanner) (string, error) {

	return fmt.Sprintf("%d", 1), nil
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
