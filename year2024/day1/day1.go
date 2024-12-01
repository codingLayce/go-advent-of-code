package day1

import (
	"advent/lib/slices"
	"bufio"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func SolvePart1(reader *bufio.Scanner) (string, error) {
	left, right, err := readLists(reader)
	if err != nil {
		return "", err
	}

	// Sorts ascending.
	sort.Ints(left)
	sort.Ints(right)

	// I can now compare each index because lists are sorted.
	diff := 0
	for idx, leftValue := range left {
		diff += int(math.Abs(float64(leftValue) - float64(right[idx])))
	}

	return fmt.Sprintf("%d", diff), nil

}
func SolvePart2(reader *bufio.Scanner) (string, error) {
	left, right, err := readLists(reader)
	if err != nil {
		return "", err
	}

	sum := 0
	for _, leftValue := range left {
		occurrences := slices.Occurrences(right, leftValue)
		sum += leftValue * occurrences
	}

	return fmt.Sprintf("%d", sum), nil
}

func readLists(reader *bufio.Scanner) ([]int, []int, error) {
	var left, right []int
	for reader.Scan() {
		arr := strings.Split(reader.Text(), " ")
		value, err := strconv.Atoi(arr[0])
		if err != nil {
			return left, right, err
		}
		left = append(left, value)
		value, err = strconv.Atoi(arr[len(arr)-1])
		if err != nil {
			return left, right, err
		}
		right = append(right, value)
	}
	return left, right, nil
}
