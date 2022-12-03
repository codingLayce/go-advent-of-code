package day3

import (
	"fmt"
	"strings"
)

type Day3 struct {
	Input string
	Dir   string
}

func New() Day3 {
	return Day3{
		Input: "2022/day3/input.txt",
		Dir:   "2022/day3/",
	}
}

func (d Day3) ProcessPuzzle1(lines []string) (string, error) {
	total := 0

	for _, line := range lines {
		compartmentL, compartmentR := splitCompartments(line)

		for _, item := range compartmentL {
			if strings.Contains(compartmentR, string(item)) {
				total += prioritize(int(item))
				break
			}
		}
	}

	return fmt.Sprintf("%d", total), nil
}

func (d Day3) ProcessPuzzle2(lines []string) (string, error) {
	total := 0

	for i := 0; i < len(lines)/3; i++ {
		group := lines[i*3 : i*3+3] // Group all 3 lines

		// Find the commons to the first and second
		commons := make(map[int32]struct{}, 0)
		for _, item := range group[0] {
			if strings.Contains(group[1], string(item)) {
				commons[item] = struct{}{}
			}
		}

		// Find the common from the previous commons with the third
		for item, _ := range commons {
			if strings.Contains(group[2], string(item)) {
				total += prioritize(int(item))
				break
			}
		}
	}

	return fmt.Sprintf("%d", total), nil
}

func prioritize(char int) int {
	if char >= 97 { // a >
		return char - 96
	} else { // A >
		return char - 64 + 26
	}
}

func splitCompartments(line string) (string, string) {
	middle := len(line) / 2
	return line[:middle], line[middle:]
}
