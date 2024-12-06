package day6

import (
	"bufio"
	"fmt"
)

func SolvePart1(reader *bufio.Scanner) (string, error) {
	area := getInput(reader)
	visited := make(map[string]struct{})

	startX, startY := getStartingPosition(area)
	ok := true
	for x, y := startX, startY; ok; x, y, ok = guardMove(area, x, y) {
		_, alreadyVisited := visited[fmt.Sprintf("%d:%d", x, y)]
		if !alreadyVisited {
			visited[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
		}
	}

	return fmt.Sprintf("%d", len(visited)), nil
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	visited := make(map[string]struct{})
	area := getInput(reader)
	startX, startY := getStartingPosition(area)
	ok := true
	for x, y := startX, startY; ok; x, y, ok = guardMove(area, x, y) {
		_, alreadyVisited := visited[fmt.Sprintf("%d:%d", x, y)]
		if !alreadyVisited {
			visited[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
		}
	}

	// Now I know all the positions the guard will be.
	// I need to simulate his path by adding a crate one by one on each position.

	// To determine if he is in a loop, I'm going to check if he reaches a position more than a given threshold.
	// I've started with 50 (execution time around 10s) and the answer was correct.
	// I've tuned it to 20 (execution time around 6s).

	count := 0
	delete(visited, fmt.Sprintf("%d:%d", startX, startY))
	for position := range visited {
		var crateX, crateY int
		_, _ = fmt.Sscanf(position, "%d:%d", &crateX, &crateY) // No error check because it's only indexes
		area[crateX][crateY] = '#'

		steps := make(map[string]int)
		ok = true
		for x, y := startX, startY; ok; x, y, ok = guardMove(area, x, y) {
			steps[fmt.Sprintf("%d:%d", x, y)]++
			if steps[fmt.Sprintf("%d:%d", x, y)] > 20 {
				count++
				break
			}
		}
		area = resetArea(area, startX, startY, crateX, crateY)
	}

	return fmt.Sprintf("%d", count), nil
}

func resetArea(area [][]byte, startX, startY, crateX, crateY int) [][]byte {
	// Clear guard position
	for row := 0; row < len(area); row++ {
		for col := 0; col < len(area[row]); col++ {
			switch area[row][col] {
			case '>', 'v', '<', '^':
				area[row][col] = '.'
			}
		}
	}
	// Clear crate
	area[crateX][crateY] = '.'
	// Position the guard
	area[startX][startY] = '^'
	return area
}

// Either turn or advance.
// Returns the new position or the previous one if only turned.
// Returns a boolean indicating if the guard in still in the map.
func guardMove(area [][]byte, x, y int) (newX int, newY int, ok bool) {
	direction := area[x][y]
	switch direction {
	case '^':
		if x == 0 { // Leaving the map
			area[x][y] = '.'
			break
		}
		if area[x-1][y] == '#' { // Obstacle in front, turn
			area[x][y] = '>'
			return x, y, true
		}
		// Move forward
		area[x-1][y] = '^'
		area[x][y] = '.'
		return x - 1, y, true
	case '>':
		if y == len(area[x])-1 { // Leaving the map
			area[x][y] = '.'
			break
		}
		if area[x][y+1] == '#' { // Obstacle in front, turn
			area[x][y] = 'v'
			return x, y, true
		}
		// Move forward
		area[x][y+1] = '>'
		area[x][y] = '.'
		return x, y + 1, true
	case 'v':
		if x == len(area)-1 { // Leaving the map
			area[x][y] = '.'
			break
		}
		if area[x+1][y] == '#' { // Obstacle in front, turn
			area[x][y] = '<'
			return x, y, true
		}
		// Move forward
		area[x+1][y] = 'v'
		area[x][y] = '.'
		return x + 1, y, true
	case '<':
		if y == 0 { // Leaving the map
			area[x][y] = '.'
			break
		}
		if area[x][y-1] == '#' { // Obstacle in front, turn
			area[x][y] = '^'
			return x, y, true
		}
		// Move forward
		area[x][y-1] = '<'
		area[x][y] = '.'
		return x, y - 1, true
	}
	return x, y, false
}

func getStartingPosition(area [][]byte) (x int, y int) {
	for row := 0; row < len(area); row++ {
		for col := 0; col < len(area[row]); col++ {
			if area[row][col] == '^' {
				return row, col
			}
		}
	}
	return 0, 0
}

func getInput(reader *bufio.Scanner) [][]byte {
	var res [][]byte
	for reader.Scan() {
		res = append(res, []byte(reader.Text()))
	}
	return res
}
