package day9

import (
	"bufio"
	"container/list"
	"fmt"
)

func SolvePart1(reader *bufio.Scanner) (string, error) {
	reader.Scan()
	disk := buildList(reader.Text())

	compacted := list.New()
	for e := disk.Front(); e != nil; e = e.Next() {
		if e.Value != -1 {
			compacted.PushBack(e.Value)
		} else {
			back := disk.Back()
			for back.Value == -1 {
				disk.Remove(back)
				back = disk.Back()
			}
			compacted.PushBack(back.Value)
			disk.Remove(back)
		}
	}
	sum := checksum(compacted)

	return fmt.Sprintf("%d", sum), nil
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	reader.Scan()
	disk := buildList(reader.Text())

	// Look for each file ID starting from the max one
	for id := getMaxID(disk); id >= 0; id-- {
		length, index, start := getFile(disk, id)
		space, spaceIndex := getLeftMostFreeSpace(disk, length)
		if space == nil || spaceIndex > index {
			continue
		}
		swap(space, start, id, length)
	}

	sum := checksum(disk)

	return fmt.Sprintf("%d", sum), nil
}

func buildList(value string) *list.List {
	l := list.New()
	index := 0
	for i, ch := range value {
		nb := digits[byte(ch)]
		if i%2 == 0 {
			for j := 0; j < nb; j++ {
				l.PushBack(index)
			}
			index++
		} else {
			for j := 0; j < nb; j++ {
				l.PushBack(-1)
			}
		}
	}
	return l
}

func checksum(l *list.List) int {
	sum := 0
	idx := -1
	for e := l.Front(); e != nil; e = e.Next() {
		idx++
		if e.Value == -1 {
			continue
		}
		value := e.Value.(int)
		sum += value * idx
	}
	return sum
}

func getMaxID(l *list.List) int {
	el := l.Back()
	for ; el != nil && el.Value == -1; el = el.Prev() {
	}
	if el == nil {
		return -1
	}
	return el.Value.(int)
}

func getFile(l *list.List, id int) (int, int, *list.Element) {
	count := 0
	index := 0
	startIndex := 0
	var startEl *list.Element
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == id {
			if count == 0 {
				startIndex = index
				startEl = e
			}
			count++
		} else if count != 0 { // Stop early because the file are ordered by id
			return count, startIndex, startEl
		}
		index++
	}
	return count, startIndex, startEl
}

func getLeftMostFreeSpace(l *list.List, target int) (*list.Element, int) {
	count := 0
	index := 0
	var start *list.Element
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == -1 {
			if count == 0 {
				start = e
			}
			count++
			if count == target {
				return start, index
			}
		} else {
			count = 0
		}
		index++
	}
	return nil, -1
}

func swap(space, start *list.Element, id, count int) {
	i := 0
	for e := space; e != nil && i < count; e = e.Next() {
		e.Value = id
		i++
	}
	i = 0
	for e := start; e != nil && i < count; e = e.Next() {
		e.Value = -1
		i++
	}
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
