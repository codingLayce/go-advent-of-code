package day2

import (
	"adc/common/maps"
	"adc/common/slices"
	"adc/common/strings"
	"math"
)

type Day2 struct {
	Input string
}

func New() Day2 {
	return Day2{
		Input: "2018/day2/input.txt",
	}
}

func (d Day2) ProcessPuzzle1(lines []string) (interface{}, error) {
	containingTwo := 0
	containingThree := 0
	for _, line := range lines {
		cacheLn := letterOccurrence(line)
		values := maps.ValuesInt(cacheLn)
		if slices.Contains(values, 2) {
			containingTwo++
		}
		if slices.Contains(values, 3) {
			containingThree++
		}
	}

	return containingTwo * containingThree, nil
}

func (d Day2) ProcessPuzzle2(lines []string) (interface{}, error) {
	var a, b string
	min := math.MaxInt

	for _, line := range lines {
		for _, nextLine := range lines {
			if line == nextLine {
				continue
			}

			count := countCharsDiffer(line, nextLine)
			if count < min {
				min = count
				a = line
				b = nextLine
			}
		}
	}

	build := ""
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			build += string(a[i])
		}
	}

	return build, nil
}

func letterOccurrence(line string) map[string]int {
	cacheLn := make(map[string]int)

	for i := 0; i < len(line); i++ {
		char := strings.CharAt(line, i)

		if _, ok := cacheLn[char]; !ok {
			cacheLn[char] = 1
		} else {
			cacheLn[char]++
		}
	}

	return cacheLn
}

func countCharsDiffer(a, b string) int {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count
}
