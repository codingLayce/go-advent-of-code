package day8

import (
	"advent/lib/vector"
	"bufio"
	"fmt"
)

func SolvePart1(reader *bufio.Scanner) (string, error) {
	grid := getInput(reader)
	points := getAntennasPositions(grid)

	maxX := len(grid) - 1
	maxY := len(grid[0]) - 1
	antinodes := make(map[vector.Vec2]struct{})
	// Look for all combinaisons of the same antenna type
	for _, positions := range points {
		for _, a := range positions {
			for _, b := range positions {
				if a == b {
					continue
				}
				// Determine the direction vector between the 2 points
				direction := b.Sub(a)
				// Apply the direction to 'b' in order to find the next aligned point
				point := b.Add(direction)
				// Only counts if it's in bound
				if point.X >= 0 && point.X <= maxX && point.Y >= 0 && point.Y <= maxY {
					antinodes[point] = struct{}{}
				}
			}
		}
	}

	return fmt.Sprintf("%d", len(antinodes)), nil
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	grid := getInput(reader)
	points := getAntennasPositions(grid)

	maxX := len(grid) - 1
	maxY := len(grid[0]) - 1
	antinodes := make(map[vector.Vec2]struct{})
	// Look for all combinaisons of the same antenna type
	for _, positions := range points {
		for _, a := range positions {
			for _, b := range positions {
				if a == b {
					continue
				}
				// Determine the direction vector between the 2 points
				direction := b.Sub(a)
				// Starts from b and apply the direction until it's out of bound
				point := b
				for ; point.X >= 0 && point.X <= maxX && point.Y >= 0 && point.Y <= maxY; point = point.Add(direction) {
					antinodes[point] = struct{}{}
				}
			}
		}
	}

	return fmt.Sprintf("%d", len(antinodes)), nil
}

func getInput(reader *bufio.Scanner) [][]byte {
	var res [][]byte
	for reader.Scan() {
		res = append(res, []byte(reader.Text()))
	}
	return res
}

func getAntennasPositions(grid [][]byte) map[byte][]vector.Vec2 {
	points := make(map[byte][]vector.Vec2)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != '.' {
				_, ok := points[grid[row][col]]
				if ok {
					points[grid[row][col]] = append(points[grid[row][col]], vector.NewVec2(row, col))
				} else {
					points[grid[row][col]] = []vector.Vec2{vector.NewVec2(row, col)}
				}
			}
		}
	}
	return points
}
