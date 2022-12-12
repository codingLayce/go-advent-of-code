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
	value int32
	pos   math.Vector2
}

// To solve the puzzle 1 I'll use the A* Algorithm
// doc: https://www.redblobgames.com/pathfinding/a-star/introduction.html

func (d Day12) ProcessPuzzle1(lines []string) (string, error) {
	grid := parseGrid(lines)
	start := getStart(grid)
	end := getEnd(grid)

	frontier := queue.NewPriorityQueue[Case]()
	frontier.Push(start, 0)
	path := make(map[Case]*Case) // Reverse path to get to the given case
	path[start] = nil
	costs := make(map[Case]int) // Best cost to get to the given case
	costs[start] = 0

	for !frontier.IsEmpty() {
		current := frontier.Pop()

		if current.value == 'E' {
			break
		}

		for _, next := range filterNeighbors(current.value, getNeighbors(grid, current.pos)) {
			newCost := costs[current] + 1
			if _, ok := path[next]; !ok || newCost < costs[start] {
				costs[next] = newCost
				priority := newCost + end.pos.DistanceManathan(current.pos)
				frontier.Push(next, priority)
				path[next] = &current
			}
		}
	}

	steps := 0
	for current := path[end]; current != nil; {
		steps++
		current = path[*current]
	}

	return fmt.Sprintf("%d", steps), nil
}

func (d Day12) ProcessPuzzle2(lines []string) (string, error) {
	return "not implemented", nil
}

func filterNeighbors(current int32, neighbors []Case) []Case {
	if current == 'S' {
		return neighbors
	}

	var filtered []Case
	for _, neighbor := range neighbors {
		if neighbor.value == 'E' && current == 'z' {
			filtered = append(filtered, neighbor)
			continue
		}
		if neighbor.value == current || neighbor.value == current+1 { // check if the neighbor is one level up
			filtered = append(filtered, neighbor)
		}
	}

	return filtered
}

func getNeighbors(grid [][]Case, pos math.Vector2) []Case {
	var neighbors []Case
	for _, direction := range directions {
		lookup := pos.Add(direction)
		if lookup.X < 0 || lookup.Y < 0 || lookup.Y >= len(grid) || lookup.X >= len(grid[0]) {
			continue
		}
		neighbors = append(neighbors, grid[lookup.Y][lookup.X])
	}

	return neighbors
}

func getStart(grid [][]Case) Case {
	for _, row := range grid {
		for _, col := range row {
			if col.value == 'S' {
				return col
			}
		}
	}
	return Case{}
}

func getEnd(grid [][]Case) Case {
	for _, row := range grid {
		for _, col := range row {
			if col.value == 'E' {
				return col
			}
		}
	}
	return Case{}
}

func parseGrid(lines []string) [][]Case {
	var grid [][]Case

	for y, line := range lines {
		var lineGrid []Case
		for x, char := range line {
			lineGrid = append(lineGrid, Case{value: char, pos: math.NewVector2(x, y)})
		}
		grid = append(grid, lineGrid)
	}
	return grid
}
