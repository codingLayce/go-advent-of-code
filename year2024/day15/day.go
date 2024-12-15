package day

import (
	"advent/lib/vector"
	"bufio"
	"fmt"
)

var directions = map[byte]vector.Vec2{
	'<': vector.LeftDir,
	'>': vector.RightDir,
	'v': vector.BottomDir,
	'^': vector.TopDir,
}

func SolvePart1(reader *bufio.Scanner) (string, error) {
	grid, moves := getInput(reader)

	robot := getRobotPos(grid)
	for _, move := range moves {
		dir := directions[move]
		attempt := robot.Add(dir)
		switch grid[attempt.X][attempt.Y] {
		case '.':
			grid[robot.X][robot.Y] = '.'
			robot = attempt
			grid[robot.X][robot.Y] = '@'
		case '#':
		case 'O':
			hasPushed := false
			grid, hasPushed = tryToPush(grid, attempt, dir)
			if hasPushed {
				grid[robot.X][robot.Y] = '.'
				robot = attempt
				grid[robot.X][robot.Y] = '@'
			}
		}
	}

	return fmt.Sprintf("%d", gpsCoordinates(grid)), nil
}

func tryToPush(grid [][]byte, pos, dir vector.Vec2) ([][]byte, bool) {
	if canBePush(grid, pos, dir) {
		return push(grid, pos, dir), true
	}
	return grid, false
}

func push(grid [][]byte, pos, dir vector.Vec2) [][]byte {
	grid[pos.X][pos.Y] = '.'
	next := pos.Add(dir)
	for ; grid[next.X][next.Y] == 'O'; next = next.Add(dir) {

	}
	grid[next.X][next.Y] = 'O'

	return grid
}

func canBePush(grid [][]byte, pos, dir vector.Vec2) bool {
	if grid[pos.X][pos.Y] == '.' {
		return true
	}
	if grid[pos.X][pos.Y] == '#' {
		return false
	}
	return canBePush(grid, pos.Add(dir), dir)
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	grid, moves := getInput(reader)
	grid = modifyGridForPart2(grid)
	cells := gridToCells(grid)

	robot := getRobotPos(grid)
	for _, move := range moves {
		fmt.Println(string(move))
		dir := directions[move]
		attempt := robot.Add(dir)
		switch grid[attempt.X][attempt.Y] {
		case '.':
			grid[robot.X][robot.Y] = '.'
			robot = attempt
			grid[robot.X][robot.Y] = '@'
		case '#':
		case '[', ']':

		}

		debug(grid)
		fmt.Println()
	}

	return fmt.Sprintf("%d", gpsCoordinates(grid)), nil
}

func gridToCells(grid [][]byte) []*Cell {
	var cells []*Cell
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			cells = append(cells, &Cell{value: grid[row][col], pos: vector.NewVec2(row, col)})
		}
	}

	for _, cell := range cells {
		cell.top = findCell(cells, cell.pos.Add(vector.TopDir))
		cell.left = findCell(cells, cell.pos.Add(vector.LeftDir))
		cell.right = findCell(cells, cell.pos.Add(vector.RightDir))
		cell.down = findCell(cells, cell.pos.Add(vector.BottomDir))
	}

	return cells
}

func findCell(cells []*Cell, pos vector.Vec2) *Cell {
	for _, cell := range cells {
		if cell.pos == pos {
			return cell
		}
	}
	return nil
}

type Cell struct {
	value byte
	pos   vector.Vec2
	left  *Cell
	top   *Cell
	right *Cell
	down  *Cell
}

func (c *Cell) CanMove(dir vector.Vec2) bool {
	switch dir {
	case vector.LeftDir:
		return c.left != nil && c.left.value == '.'
	case vector.RightDir:
		return c.right != nil && c.right.value == '.'
	case vector.TopDir:
		return c.top != nil && c.top.value == '.'
	case vector.BottomDir:
		return c.down != nil && c.down.value == '.'
	default:
		return false
	}
}

func modifyGridForPart2(grid [][]byte) [][]byte {
	newGrid := make([][]byte, len(grid))
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			switch grid[row][col] {
			case '#':
				newGrid[row] = append(newGrid[row], '#', '#')
			case 'O':
				newGrid[row] = append(newGrid[row], '[', ']')
			case '.':
				newGrid[row] = append(newGrid[row], '.', '.')
			case '@':
				newGrid[row] = append(newGrid[row], '@', '.')
			}
		}
	}
	return newGrid
}

func gpsCoordinates(grid [][]byte) int {
	sum := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'O' {
				sum += 100*row + col
			}
		}
	}
	return sum
}

func getRobotPos(grid [][]byte) vector.Vec2 {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '@' {
				return vector.NewVec2(row, col)
			}
		}
	}
	return vector.Vec2{} // Shouldn't happen
}

func debug(grid [][]byte) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			fmt.Printf("%c ", grid[row][col])
		}
		fmt.Printf("\n")
	}
}

func getInput(reader *bufio.Scanner) ([][]byte, []byte) {
	reader.Scan()
	var grid [][]byte
	for line := reader.Text(); line != ""; line = reader.Text() {
		grid = append(grid, []byte(reader.Text()))
		reader.Scan()
	}

	var moves []byte
	for reader.Scan() {
		moves = append(moves, []byte(reader.Text())...)
	}

	return grid, moves
}
