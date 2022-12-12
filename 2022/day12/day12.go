package day12

import (
	"fmt"

	"adc/common/math"
	"adc/common/queue"
)

type Day12 struct {
	Input string
	Dir   string
}

func New() Day12 {
	return Day12{
		Input: "2022/day12/input.txt",
		Dir:   "2022/day12/",
	}
}

var directions = []math.Vector2{
	math.NewVector2(1, 0),  // right
	math.NewVector2(-1, 0), // left
	math.NewVector2(0, 1),  // down
	math.NewVector2(0, -1), // up
}

type Case struct {
	value     int32
	elevation int32
	pos       math.Vector2
	cost      int
}

func resetCosts(grid [][]*Case) {
	for _, row := range grid {
		for _, col := range row {
			col.cost = 0
		}
	}
}

// To solve the puzzle 1 I'll use the A* Algorithm
// doc: https://www.redblobgames.com/pathfinding/a-star/introduction.html

func (d Day12) ProcessPuzzle1(lines []string) (string, error) {
	grid := parseGrid(lines)
	start := getStart(grid)
	end := getEnd(grid)

	return fmt.Sprintf("%d", pathFinding(grid, start, end)), nil
}

func (d Day12) ProcessPuzzle2(lines []string) (string, error) {
	grid := parseGrid(lines)
	end := getEnd(grid)

	possibleStarts := []*Case{getStart(grid)}
	for _, row := range grid {
		for _, col := range row {
			if col.value == 'a' {
				possibleStarts = append(possibleStarts, col)
			}
		}
	}

	min := 999999999
	for _, start := range possibleStarts {
		resetCosts(grid)
		cost := pathFinding(grid, start, end)
		if cost < min && cost > 0 {
			min = cost
		}
	}

	return fmt.Sprintf("%d", min), nil
}

func pathFinding(grid [][]*Case, start, end *Case) int {
	frontier := queue.NewPriorityQueue[*Case]()
	frontier.Push(start, 0)

	for !frontier.IsEmpty() {
		current := frontier.Pop()

		for _, next := range getNeighbors(grid, current.pos) {
			newCost := current.cost + 1
			if next.cost == 0 || newCost < next.cost {
				next.cost = newCost
				priority := newCost + end.pos.DistanceManathan(next.pos)
				frontier.Push(next, priority)
			}
		}
	}

	return end.cost
}

func getNeighbors(grid [][]*Case, pos math.Vector2) []*Case {
	var neighbors []*Case
	for _, direction := range directions {
		lookup := pos.Add(direction)
		if lookup.X < 0 || lookup.Y < 0 || lookup.Y >= len(grid) || lookup.X >= len(grid[0]) {
			continue
		}
		if grid[lookup.Y][lookup.X].elevation <= grid[pos.Y][pos.X].elevation+1 {
			neighbors = append(neighbors, grid[lookup.Y][lookup.X])
		}
	}

	return neighbors
}

func getStart(grid [][]*Case) *Case {
	for _, row := range grid {
		for _, col := range row {
			if col.value == 'S' {
				return col
			}
		}
	}
	return nil
}

func getEnd(grid [][]*Case) *Case {
	for _, row := range grid {
		for _, col := range row {
			if col.value == 'E' {
				return col
			}
		}
	}
	return nil
}

func parseGrid(lines []string) [][]*Case {
	var grid [][]*Case

	for y, line := range lines {
		var lineGrid []*Case
		for x, char := range line {
			tmp := char
			if char == 'S' {
				tmp = 'a'
			} else if char == 'E' {
				tmp = 'z'
			}
			lineGrid = append(lineGrid, &Case{value: char, pos: math.NewVector2(x, y), elevation: tmp - 'a'})
		}
		grid = append(grid, lineGrid)
	}
	return grid
}
