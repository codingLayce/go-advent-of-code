package day

import (
	"advent/lib/vector"
	"bufio"
	"fmt"
	"slices"
)

type Region struct {
	plant      byte
	area       int
	perimeters int
}

func SolvePart1(reader *bufio.Scanner) (string, error) {
	grid := getInput(reader)

	var regions []Region
	visited := make(map[vector.Vec2]struct{})

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if _, ok := visited[vector.NewVec2(row, col)]; ok {
				continue
			}
			var region Region
			region, _ = findRegion(grid, visited, vector.NewVec2(row, col), Region{plant: grid[row][col]})
			regions = append(regions, region)
		}
	}

	sum := 0
	for _, region := range regions {
		sum += region.area * region.perimeters
	}

	return fmt.Sprintf("%d", sum), nil
}

func findRegion(grid [][]byte, visited map[vector.Vec2]struct{}, pos vector.Vec2, region Region) (Region, bool) {
	if pos.X < 0 || pos.X >= len(grid) || pos.Y < 0 || pos.Y >= len(grid[pos.X]) {
		return region, false
	}
	if grid[pos.X][pos.Y] != region.plant {
		return region, false
	}
	if _, ok := visited[pos]; ok {
		return region, grid[pos.X][pos.Y] == region.plant
	}

	visited[pos] = struct{}{}
	region.area++

	samePlant := false
	if region, samePlant = findRegion(grid, visited, pos.Add(vector.RightDir), region); !samePlant { // right
		region.perimeters++
	}
	if region, samePlant = findRegion(grid, visited, pos.Add(vector.BottomDir), region); !samePlant { // down
		region.perimeters++
	}
	if region, samePlant = findRegion(grid, visited, pos.Add(vector.LeftDir), region); !samePlant { // left
		region.perimeters++
	}
	if region, samePlant = findRegion(grid, visited, pos.Add(vector.TopDir), region); !samePlant { // top
		region.perimeters++
	}

	return region, true
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	grid := getInput(reader)

	var regions []Region2
	visited := make(map[vector.Vec2]struct{})

	// Read the grid and extract the regions.
	// Each region is composed of nodes and each node has a barrier on each side by default (4 sides).
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if _, ok := visited[vector.NewVec2(row, col)]; ok {
				continue
			}
			regions = append(regions, findRegion2(grid, visited, vector.NewVec2(row, col), Region2{plant: grid[row][col]}))
		}
	}

	// For each region compute the number of sides and compute the price.
	sum := 0
	for _, region := range regions {
		sides := region.Sides()
		sum += sides * len(region.nodes)
	}

	return fmt.Sprintf("%d", sum), nil
}

type Region2 struct {
	plant byte
	nodes []*Node
}

func (r Region2) Barriers() {
	// Remove unnecessary barriers (When a node of the same region is next to it).
	// Each node will know if a barrier is present at one of each side.
	for _, node := range r.nodes {
		if r.FindNextNode(node.pos.Add(vector.BottomDir)) != nil {
			node.UnsetBarrier(bottomBarrier)
		}
		if r.FindNextNode(node.pos.Add(vector.TopDir)) != nil {
			node.UnsetBarrier(topBarrier)
		}
		if r.FindNextNode(node.pos.Add(vector.RightDir)) != nil {
			node.UnsetBarrier(rightBarrier)
		}
		if r.FindNextNode(node.pos.Add(vector.LeftDir)) != nil {
			node.UnsetBarrier(leftBarrier)
		}
	}
}

func (r Region2) Sides() int {
	r.Barriers()

	// For each node, walk alongside the barriers and mark all the node's barrier it reaches.
	// By marking them I'm sure I'll never count them anymore, so I'm increasing the counter
	// only for the first barrier of each side is reached.
	count := 0
	for _, node := range r.nodes {
		if node.HasBarrier(topBarrier) && !node.HasBarrierVisited(topBarrier) { // The node has a barrier on top and hasn't been visited yet
			count++
			// Mark all the barriers of this side in order to not count them anymore
			r.walkAndMark(node, vector.RightDir, topBarrier)
			r.walkAndMark(node, vector.LeftDir, topBarrier)
		}
		if node.HasBarrier(bottomBarrier) && !node.HasBarrierVisited(bottomBarrier) { // The node has a barrier on bottom and hasn't been visited yet
			count++
			// Mark all the barriers of this side in order to not count them anymore
			r.walkAndMark(node, vector.RightDir, bottomBarrier)
			r.walkAndMark(node, vector.LeftDir, bottomBarrier)
		}
		if node.HasBarrier(rightBarrier) && !node.HasBarrierVisited(rightBarrier) { // The node has a barrier on right and hasn't been visited yet
			count++
			// Mark all the barriers of this side in order to not count them anymore
			r.walkAndMark(node, vector.TopDir, rightBarrier)
			r.walkAndMark(node, vector.BottomDir, rightBarrier)
		}
		if node.HasBarrier(leftBarrier) && !node.HasBarrierVisited(leftBarrier) { // The node has a barrier on right and hasn't been visited yet
			count++
			// Mark all the barriers of this side in order to not count them anymore
			r.walkAndMark(node, vector.TopDir, leftBarrier)
			r.walkAndMark(node, vector.BottomDir, leftBarrier)
		}
	}

	return count
}

func (r Region2) walkAndMark(node *Node, direction vector.Vec2, barrier int) {
	for node != nil && node.HasBarrier(barrier) {
		node.SetBarrierVisited(barrier)
		node = r.FindNextNode(node.pos.Add(direction))
	}
}

func (r Region2) FindNextNode(pos vector.Vec2) *Node {
	idx := slices.IndexFunc(r.nodes, func(n *Node) bool {
		return n.pos == pos
	})
	if idx != -1 {
		return r.nodes[idx]
	}
	return nil
}

const (
	topBarrier int = iota
	rightBarrier
	bottomBarrier
	leftBarrier
)

type Node struct {
	pos             vector.Vec2
	barriers        [4]bool
	barriersVisited [4]bool
}

func NewNode(pos vector.Vec2) *Node {
	return &Node{
		pos:             pos,
		barriers:        [4]bool{true, true, true, true},
		barriersVisited: [4]bool{false, false, false, false},
	}
}

func (n *Node) SetBarrier(barrier int) {
	n.barriers[barrier] = true
}

func (n *Node) UnsetBarrier(barrier int) {
	n.barriers[barrier] = false
}

func (n *Node) HasBarrier(barrier int) bool {
	return n.barriers[barrier]
}

func (n *Node) SetBarrierVisited(barrier int) {
	n.barriersVisited[barrier] = true
}

func (n *Node) HasBarrierVisited(barrier int) bool {
	return n.barriersVisited[barrier]
}

func findRegion2(grid [][]byte, visited map[vector.Vec2]struct{}, pos vector.Vec2, region Region2) Region2 {
	if pos.X < 0 || pos.X >= len(grid) || pos.Y < 0 || pos.Y >= len(grid[pos.X]) {
		return region
	}
	if grid[pos.X][pos.Y] != region.plant {
		return region
	}
	if _, ok := visited[pos]; ok {
		return region
	}

	visited[pos] = struct{}{}
	region.nodes = append(region.nodes, NewNode(pos))
	region = findRegion2(grid, visited, pos.Add(vector.RightDir), region)
	region = findRegion2(grid, visited, pos.Add(vector.BottomDir), region)
	region = findRegion2(grid, visited, pos.Add(vector.LeftDir), region)
	region = findRegion2(grid, visited, pos.Add(vector.TopDir), region)
	return region
}

func getInput(reader *bufio.Scanner) [][]byte {
	var res [][]byte
	for reader.Scan() {
		res = append(res, []byte(reader.Text()))
	}
	return res
}
