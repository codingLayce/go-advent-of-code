package day7

import (
	"adc/common/slices"
	"fmt"
	"strconv"
	"strings"
)

type Day7 struct {
	Input string
	Dir   string
}

func New() Day7 {
	return Day7{
		Input: "2022/day7/input.txt",
		Dir:   "2022/day7/",
	}
}

func newDir(parent *Dir, name string) *Dir {
	return &Dir{
		name:     name,
		parent:   parent,
		children: []*Dir{},
		size:     0,
	}
}

type Dir struct {
	name     string
	parent   *Dir
	children []*Dir
	size     uint64
}

func (d *Dir) String() string {
	parent := ""
	if d.parent != nil {
		parent = d.parent.name
	}
	return fmt.Sprintf("[%s] (%s) %d childs, size: %d", d.name, parent, len(d.children), d.size)
}

func (d *Dir) debug(level int) {
	header := strings.Repeat("   ", level)
	fmt.Printf("%s%s\n", header, d)
	for _, child := range d.children {
		child.debug(level + 1)
	}
}

func (d *Dir) addDir(name string) *Dir {
	next := newDir(d, name)
	d.children = append(d.children, next)
	return next
}

func (d *Dir) addFile(size uint64) {
	d.size += size
}

func (d *Dir) computeSize() {
	if len(d.children) == 0 {
		return
	}

	for _, child := range d.children {
		child.computeSize()
		d.size += child.size
	}
}

func (d Day7) ProcessPuzzle1(lines []string) (string, error) {
	root := parseFileSystem(lines)
	root.computeSize()

	res := sumSizesAtMost(root, 100000)

	return fmt.Sprintf("%d", res), nil
}

func (d Day7) ProcessPuzzle2(lines []string) (string, error) {
	root := parseFileSystem(lines)
	root.computeSize()

	needs := 30000000 - (70000000 - root.size)
	candidates := foundDeletableCandidates(root, needs)
	_, min := slices.Min(candidates)

	return fmt.Sprintf("%d", min), nil
}

func foundDeletableCandidates(dir Dir, needs uint64) []uint64 {
	var all []uint64
	if dir.size >= needs {
		all = append(all, dir.size)
	}
	for _, child := range dir.children {
		all = append(all, foundDeletableCandidates(*child, needs)...)
	}
	return all
}

func sumSizesAtMost(dir Dir, atMost uint64) uint64 {
	if len(dir.children) == 0 {
		if dir.size <= atMost {
			return dir.size
		}
		return 0
	}
	sum := uint64(0)
	if dir.size <= atMost {
		sum = dir.size
	}
	for _, child := range dir.children {
		sum += sumSizesAtMost(*child, atMost)
	}
	return sum
}

func parseFileSystem(lines []string) Dir {
	root := &Dir{
		name:     "/",
		parent:   nil,
		children: []*Dir{},
		size:     0,
	}
	currentDir := root

	for _, line := range lines[1:] {
		// $ <command>
		if line[0] == uint8('$') {
			// ls
			if strings.Contains(line, "ls") {
				continue
			}

			// cd ..
			if strings.Contains(line, "..") {
				currentDir = currentDir.parent
				continue
			}

			// cd <dir>
			idx := strings.LastIndex(line, " ") + 1
			currentDir = currentDir.addDir(line[idx:])
			continue
		}

		// dir <dir>
		if line[0] == uint8('d') {
			continue
		}

		// <size> <file>
		idx := strings.Index(line, " ")
		size, _ := strconv.ParseUint(line[:idx], 10, 64)
		currentDir.addFile(size)
	}
	return *root
}
