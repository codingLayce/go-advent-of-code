package day9

import (
	"bufio"
	"fmt"
	"slices"
)

func SolvePart1(reader *bufio.Scanner) (string, error) {
	reader.Scan()

	root := buildTree(reader.Text(), 0, 1, nil)
	root.Walk()
	fmt.Println()
	root.compact()
	root.Walk()
	fmt.Println()

	return fmt.Sprintf("%d", 1), nil
}

func SolvePart2(reader *bufio.Scanner) (string, error) {

	return fmt.Sprintf("%d", 1), nil
}

type Node struct {
	id       int
	memory   []int
	next     *Node
	previous *Node
}

func (n *Node) compact() {
	index := slices.Index(n.memory, -1)
	for index != -1 {
		deepest := n.getDeepestIndex()
		if deepest == -1 {
			return
		}
		fmt.Println(deepest)
		n.memory[index] = deepest
		index = slices.Index(n.memory, -1)
	}
	if n.next == nil {
		return
	}
	n.next.compact()
}

func (n *Node) getDeepestIndex() int {
	if n.next == nil {
		return n.extractLeftMostIndex()
	}
	v := n.next.getDeepestIndex()
	if v != -1 {
		return v
	}
	return n.extractLeftMostIndex()
}

func (n *Node) extractLeftMostIndex() int {
	for i := len(n.memory) - 1; i >= 0; i-- {
		if n.memory[i] != -1 {
			tmp := n.memory[i]
			n.memory[i] = -1
			return tmp
		}
	}
	return -1
}

func (n *Node) Walk() {
	for _, idx := range n.memory {
		if idx == -1 {
			fmt.Printf(".")
		} else {
			fmt.Printf("%d", idx)
		}
	}

	if n.next == nil {
		return
	}
	n.next.Walk()
}

func buildTree(disk string, index, nodeType int, previous *Node) *Node {
	if len(disk) == 0 {
		return nil
	}
	nb := digits[disk[0]]
	n := &Node{previous: previous}
	if nodeType == -1 {
		mem := make([]int, nb)
		for i := range mem {
			mem[i] = -1
		}
		n.id = -1
		n.memory = mem
		n.next = buildTree(disk[1:], index, nodeType*-1, n)
	} else {
		mem := make([]int, nb)
		for i := range mem {
			mem[i] = index
		}
		n.id = index
		n.memory = mem
		n.next = buildTree(disk[1:], index+1, nodeType*-1, n)
	}
	return n
}

var digits = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}
