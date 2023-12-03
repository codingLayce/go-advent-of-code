package day3

import (
	"fmt"
	"strconv"
	"strings"
)

type Day struct {
	Input string
	Dir   string
}

func New() Day {
	return Day{
		Input: "2023/day3/input.txt",
		Dir:   "2023/day3/",
	}
}

func (d Day) ProcessPuzzle1(lines []string) (string, error) {
	sum := 0
	for col, line := range lines {
		numbers := getNumbers(line)

		cache := make(map[string]int)
		for _, number := range numbers {
			idx := findIndex(line, number, 0, cache[number])

			if isPart(idx, col, len(number), lines) {
				nb, err := strconv.Atoi(number)
				if err != nil {
					return "", err
				}

				sum += nb
			}

			if _, ok := cache[number]; ok {
				cache[number] += 1
			}
			cache[number] = 1
		}
	}

	return fmt.Sprintf("%d", sum), nil
}

func (d Day) ProcessPuzzle2(lines []string) (string, error) {
	return "not implemented", nil
}

// isPart determines if the given coordinates are surrounded by a symbol.
func isPart(startX, startY, length int, lines []string) bool {
	for x := startX - 1; x <= startX+length+1; x++ {
		if checkCoordinates(x, startY+1, lines) || checkCoordinates(x, startY, lines) || checkCoordinates(x, startY-1, lines) {
			return true
		}
	}

	return false
}

func checkCoordinates(x, y int, lines []string) bool {
	if y < 0 || y >= len(lines) || x < 0 || x >= len(lines[0]) {
		return false
	}
	ch := lines[y][x]
	return !((ch >= '0' && ch <= '9') || ch == '.')
}

func getNumbers(line string) []string {
	var all []string
	var builder strings.Builder
	for _, ch := range line {
		if ch >= '0' && ch <= '9' {
			builder.WriteByte(byte(ch))
			continue
		}
		if builder.String() != "" {
			all = append(all, builder.String())
			builder.Reset()
		}
	}
	if builder.String() != "" {
		all = append(all, builder.String())
	}

	return all
}

func findIndex(line, searching string, current, skip int) int {
	idx := strings.Index(line, searching)
	if skip == 0 {
		return current + idx
	}
	_, remaining, _ := strings.Cut(line, searching)
	return findIndex(remaining, searching, current+idx+len(searching), skip-1)
}
