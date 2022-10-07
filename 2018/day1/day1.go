package day1

import (
	"adc/common/slices"
	"strconv"
)

type Day1 struct {
	Input string
}

func New() Day1 {
	return Day1{
		Input: "2018/day1/input.txt",
	}
}

func (d Day1) ProcessPuzzle1(lines []string) (interface{}, error) {
	current := 0
	for _, line := range lines {
		nb, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}

		switch line[0] {
		case '+':
			current += nb
		case '-':
			current -= nb
		default:
		}
	}

	return current, nil
}

func (d Day1) ProcessPuzzle2(lines []string) (interface{}, error) {
	var cache []int
	current := 0

	for true { // Loop until a cached value is reached
		for _, line := range lines {
			nb, err := strconv.Atoi(line[1:])
			if err != nil {
				return nil, err
			}

			switch line[0] {
			case '+':
				current += nb
			case '-':
				current -= nb
			default:
			}

			if slices.Contains(cache, current) {
				return current, nil
			}

			cache = append(cache, current)
		}
	}

	return "not found", nil
}
