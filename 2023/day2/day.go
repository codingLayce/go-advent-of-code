package day2

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
		Input: "2023/day2/input.txt",
		Dir:   "2023/day2/",
	}
}

type color string

var rules = map[color]int{"red": 12, "green": 13, "blue": 14}

func (d Day) ProcessPuzzle1(lines []string) (string, error) {
	sum := 0
	for _, line := range lines {
		id, remaining, err := parseGameID(line)
		if err != nil {
			return "", fmt.Errorf("parse game id for line '%s': %w", line, err)
		}

		sets, err := parseRemaining(remaining)
		if err != nil {
			return "", fmt.Errorf("parse remaining for line '%s': %w", line, err)
		}

		if isPossible(sets) {
			sum += id
		}
	}

	return fmt.Sprintf("%d", sum), nil
}

func (d Day) ProcessPuzzle2(lines []string) (string, error) {
	sum := 0
	for _, line := range lines {
		_, remaining, err := parseGameID(line)
		if err != nil {
			return "", fmt.Errorf("parse game id for line '%s': %w", line, err)
		}

		sets, err := parseRemaining(remaining)
		if err != nil {
			return "", fmt.Errorf("parse remaining for line '%s': %w", line, err)
		}

		cur := map[color]int{"red": 0, "green": 0, "blue": 0}
		for _, set := range sets {
			cur = mergeMax(cur, set)
		}

		power := 1
		for _, value := range cur {
			power *= value
		}
		sum += power
	}

	return fmt.Sprintf("%d", sum), nil
}

func mergeMax(min, set map[color]int) map[color]int {
	for key, value := range set {
		if min[key] < value {
			min[key] = value
		}
	}
	return min
}

// isPossible determines if the given set is part of the rules.
func isPossible(sets []map[color]int) bool {
	for _, set := range sets {
		for col, nb := range set {
			if nb > rules[col] {
				return false
			}
		}
	}
	return true
}

// parseRemaining retrieves the array of sets.
func parseRemaining(remaining string) ([]map[color]int, error) {
	var sets []map[color]int

	for _, set := range strings.Split(remaining, ";") {
		m := make(map[color]int)
		for _, cube := range strings.Split(set, ",") {
			value, col, _ := strings.Cut(strings.TrimSpace(cube), " ")
			nb, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("parsing to int for %s, nb=%s, color=%s", cube, value, col)
			}
			m[color(col)] = nb
		}
		sets = append(sets, m)
	}

	return sets, nil
}

// parseGameID retrieves the game ID and the remaining.
func parseGameID(line string) (int, string, error) {
	arr := strings.Split(line, ":")
	id, err := strconv.Atoi(arr[0][5:]) // arr[0] = "Game XXX" where XXX is an undefined sized number.
	if err != nil {
		return 0, "", err
	}
	return id, arr[1], nil
}
