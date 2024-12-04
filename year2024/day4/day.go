package day4

import (
	"bufio"
	"fmt"
)

func SolvePart1(reader *bufio.Scanner) (string, error) {
	grid := getInput(reader)

	count := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'X' {
				// check up
				if row-3 >= 0 {
					if fmt.Sprintf("%c%c%c%c", grid[row][col], grid[row-1][col], grid[row-2][col], grid[row-3][col]) == "XMAS" {
						count++
					}
				}
				// check down
				if row+3 < len(grid) {
					if fmt.Sprintf("%c%c%c%c", grid[row][col], grid[row+1][col], grid[row+2][col], grid[row+3][col]) == "XMAS" {
						count++
					}
				}
				// check left
				if col-3 >= 0 {
					if fmt.Sprintf("%c%c%c%c", grid[row][col], grid[row][col-1], grid[row][col-2], grid[row][col-3]) == "XMAS" {
						count++
					}
				}
				// check right
				if col+3 < len(grid[row]) {
					if fmt.Sprintf("%c%c%c%c", grid[row][col], grid[row][col+1], grid[row][col+2], grid[row][col+3]) == "XMAS" {
						count++
					}
				}
				// check up left
				if row-3 >= 0 && col-3 >= 0 {
					if fmt.Sprintf("%c%c%c%c", grid[row][col], grid[row-1][col-1], grid[row-2][col-2], grid[row-3][col-3]) == "XMAS" {
						count++
					}
				}
				// check up right
				if row-3 >= 0 && col+3 < len(grid[row]) {
					if fmt.Sprintf("%c%c%c%c", grid[row][col], grid[row-1][col+1], grid[row-2][col+2], grid[row-3][col+3]) == "XMAS" {
						count++
					}
				}
				// check down left
				if row+3 < len(grid) && col-3 >= 0 {
					if fmt.Sprintf("%c%c%c%c", grid[row][col], grid[row+1][col-1], grid[row+2][col-2], grid[row+3][col-3]) == "XMAS" {
						count++
					}
				}
				// check down right
				if row+3 < len(grid) && col+3 < len(grid[row]) {
					if fmt.Sprintf("%c%c%c%c", grid[row][col], grid[row+1][col+1], grid[row+2][col+2], grid[row+3][col+3]) == "XMAS" {
						count++
					}
				}
			}
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	grid := getInput(reader)

	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'A' {
				if row-1 < 0 || row+1 >= len(grid) || col-1 < 0 || col+1 >= len(grid[row]) {
					continue
				}

				upLeft := grid[row-1][col-1]
				upRight := grid[row-1][col+1]
				downLeft := grid[row+1][col-1]
				downRight := grid[row+1][col+1]

				if upLeft == downRight || upRight == downLeft {
					continue
				}

				// Both up to down
				if upLeft == 'M' && downRight == 'S' && upRight == 'M' && downLeft == 'S' {
					count++
				}
				// Both down to up
				if upLeft == 'S' && downRight == 'M' && upRight == 'S' && downLeft == 'M' {
					count++
				}
				// Both left to right
				if upLeft == 'M' && downRight == 'S' && downLeft == 'M' && upRight == 'S' {
					count++
				}
				// Both right to left
				if upLeft == 'S' && downRight == 'M' && downLeft == 'S' && upRight == 'M' {
					count++
				}
			}
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func getInput(reader *bufio.Scanner) []string {
	var grid []string
	for reader.Scan() {
		grid = append(grid, reader.Text())
	}
	return grid
}
