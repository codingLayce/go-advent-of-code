package day6

import (
	"adc/common/strings"
	"fmt"
)

type Day6 struct {
	Input string
	Dir   string
}

func New() Day6 {
	return Day6{
		Input: "2022/day6/input.txt",
		Dir:   "2022/day6/",
	}
}

func (d Day6) ProcessPuzzle1(lines []string) (string, error) {
	return foundDistinctChars(lines[0], 4), nil
}

func (d Day6) ProcessPuzzle2(lines []string) (string, error) {
	return foundDistinctChars(lines[0], 14), nil
}

func foundDistinctChars(value string, length int) string {
	for i := 0; i < len(value)-length; i++ {
		current := value[i : i+length]
		ok := true
		for _, char := range current {
			occurence := strings.CountOccurrence(current, string(char))
			if occurence != 1 {
				ok = false
				break
			}
		}
		if ok {
			return fmt.Sprintf("%d", i+length)
		}
	}
	return ""
}
