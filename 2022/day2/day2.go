package day2

import (
	"fmt"
	"strings"
)

type Day2 struct {
	Input string
	Dir   string
}

func New() Day2 {
	return Day2{
		Input: "2022/day2/input.txt",
		Dir:   "2022/day2/",
	}
}

var (
	scoreMapping = map[string]int{
		"X": 0, // Me Rock
		"Y": 1, // Me Paper
		"Z": 2, // Me Scissors

		"A": 0, // Op Rock
		"B": 1, // Op Paper
		"C": 2, // Op Scissors
	}
)

func (d Day2) ProcessPuzzle1(lines []string) (string, error) {
	//           Rock Paper Scissors
	// Rock        3    0      6
	// Paper       6    3      0
	// Scissors    0    6      3
	duelMatrix := [3][3]int{
		{3, 0, 6},
		{6, 3, 0},
		{0, 6, 3},
	}

	totalScore := 0
	for _, line := range lines {
		arr := strings.Split(line, " ")
		opponent := arr[0]
		me := arr[1]
		totalScore += scoreMapping[me] + 1 + duelMatrix[scoreMapping[me]][scoreMapping[opponent]]
	}

	return fmt.Sprintf("%d", totalScore), nil
}

func (d Day2) ProcessPuzzle2(lines []string) (string, error) {
	expectedEndMapping := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	scorePlayedMapping := []int{1, 2, 3}
	duelMatrix := [3][3]int{
		{3, 6, 0},
		{0, 3, 6},
		{6, 0, 3},
	}

	totalScore := 0
	for _, line := range lines {
		arr := strings.Split(line, " ")
		opponent := arr[0]
		me := arr[1]

		expectedEnd := expectedEndMapping[me]
		for i, result := range duelMatrix[scoreMapping[opponent]] {
			if result == expectedEnd {
				totalScore += expectedEnd + scorePlayedMapping[i]
				continue
			}
		}
	}

	return fmt.Sprintf("%d", totalScore), nil
}
