package day7

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SolvePart1(reader *bufio.Scanner) (string, error) {
	input, err := getInput(reader)
	if err != nil {
		return "", err
	}

	sum := 0
	for target, values := range input {
		if bruteForceOperations(target, values[0], 1, values...) {
			sum += target
		}
	}

	return fmt.Sprintf("%d", sum), nil
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	input, err := getInput(reader)
	if err != nil {
		return "", err
	}

	sum := 0
	for target, values := range input {
		if bruteForceOperations2(target, values[0], 1, values...) {
			sum += target
		}
	}

	return fmt.Sprintf("%d", sum), nil
}

func bruteForceOperations(target, cur, index int, values ...int) bool {
	if index == len(values) {
		return cur == target
	}

	if bruteForceOperations(target, cur+values[index], index+1, values...) {
		return true
	}
	return bruteForceOperations(target, cur*values[index], index+1, values...)
}

func bruteForceOperations2(target, cur, index int, values ...int) bool {
	if index == len(values) {
		return cur == target
	}

	if bruteForceOperations2(target, cur+values[index], index+1, values...) {
		return true
	}
	if bruteForceOperations2(target, cur*values[index], index+1, values...) {
		return true
	}
	concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", cur, values[index]))
	return bruteForceOperations2(target, concat, index+1, values...)
}

func getInput(reader *bufio.Scanner) (map[int][]int, error) {
	values := make(map[int][]int)

	for reader.Scan() {
		line := strings.Split(reader.Text(), ":")
		key, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, err
		}

		line = strings.Split(strings.TrimSpace(line[1]), " ")
		values[key] = make([]int, len(line))
		for idx, str := range line {
			nb, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			values[key][idx] = nb
		}
	}

	return values, nil
}
