package day1

import (
	"bufio"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func SolvePart1(reader *bufio.Scanner) (string, error) {
	input, err := getInput(reader)
	if err != nil {
		return "", err
	}

	// Filtre les lignes qui ne sont ni en incrément ou décrément
	input = slices.DeleteFunc(input, func(ints []int) bool {
		return !isDecreasing(ints) && !isIncreasing(ints)
	})

	count := 0
	for _, line := range input {
		if checkSafety(line) {
			count++
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	input, err := getInput(reader)
	if err != nil {
		return "", err
	}

	count := 0
	for _, line := range input {
		if checkSafety(line) {
			count++
			continue
		}
		for i := 0; i < len(line); i++ {
			if checkSafety(slices.Delete(slices.Clone(line), i, i+1)) {
				count++
				break
			}
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func getInput(reader *bufio.Scanner) ([][]int, error) {
	var input [][]int
	for reader.Scan() {
		line := strings.Split(reader.Text(), " ")
		a := make([]int, len(line))
		for idx, nbStr := range line {
			nb, err := strconv.Atoi(nbStr)
			if err != nil {
				return nil, err
			}
			a[idx] = nb
		}
		input = append(input, a)
	}
	return input, nil
}

func checkSafety(slice []int) bool {
	if !isIncreasing(slice) && !isDecreasing(slice) {
		return false
	}
	for i := 0; i < len(slice)-1; i++ {
		diff := math.Abs(float64(slice[i]) - float64(slice[i+1]))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isIncreasing(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] >= slice[i+1] {
			return false
		}
	}
	return true
}
func isDecreasing(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] <= slice[i+1] {
			return false
		}
	}
	return true
}
