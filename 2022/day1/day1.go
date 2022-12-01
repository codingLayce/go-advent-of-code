package day1

import (
	"adc/common/slices"
	"fmt"
	"math"
	"strconv"
)

type Day1 struct {
	Dir   string
	Input string
}

func New() Day1 {
	return Day1{
		Input: "2022/day1/input.txt",
		Dir:   "2022/day1/",
	}
}

func (d Day1) ProcessPuzzle1(lines []string) (string, error) {
	var (
		maxCarrying     = 0
		currentCarrying = 0
	)
	for _, line := range lines {
		if line == "" {
			maxCarrying = int(math.Max(float64(currentCarrying), float64(maxCarrying)))
			currentCarrying = 0
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				return "", err
			}
			currentCarrying += value
		}
	}

	return fmt.Sprintf("%d", maxCarrying), nil
}

func (d Day1) ProcessPuzzle2(lines []string) (string, error) {
	elves := make([]int, 3)
	current := 0
	for _, line := range lines {
		if line == "" {
			elves = replaceMin(elves, current)
			current = 0
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				return "", err
			}
			current += value
		}
	}
	elves = replaceMin(elves, current)
	fmt.Printf("%v\n", elves)

	return fmt.Sprintf("%d", slices.Sum(elves)), nil
}

func replaceMin(slice []int, value int) []int {
	min := math.MaxInt
	idx := 0
	for i, e := range slice {
		if e < min {
			min = e
			idx = i
		}
	}
	if value > min {
		slice[idx] = value
	}
	return slice
}
