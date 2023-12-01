package day1

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day struct {
	Input string
	Dir   string
}

func New() Day {
	return Day{
		Input: "2023/day1/input.txt",
		Dir:   "2023/day1/",
	}
}

func (d Day) ProcessPuzzle1(lines []string) (string, error) {
	sum := 0
	for _, line := range lines {
		first := findFirstDigit(line)
		last := findLastDigit(line)
		calibration := fmt.Sprintf("%s%s", first, last)
		value, err := strconv.Atoi(calibration)
		if err != nil {
			return "", fmt.Errorf("atoi line '%s': %v", line, err)
		}
		sum += value
	}

	return fmt.Sprintf("%d", sum), nil
}

func findFirstDigit(value string) string {
	for i, ch := range value {
		if ch >= '0' && ch <= '9' {
			return value[i : i+1]
		}
	}
	return ""
}

func findLastDigit(value string) string {
	for i := len(value) - 1; i >= 0; i-- {
		ch := value[i]
		if ch >= '0' && ch <= '9' {
			return value[i : i+1]
		}
	}
	return ""
}

func (d Day) ProcessPuzzle2(lines []string) (string, error) {
	sum := 0
	for _, line := range lines {
		calibration := fmt.Sprintf("%s%s", findFirstDigits2(line), findLastDigits2(line))
		value, err := strconv.Atoi(calibration)
		if err != nil {
			return "", fmt.Errorf("atoi line '%s' -> '%s': %v", line, calibration, err)
		}
		sum += value
	}
	return fmt.Sprintf("%d", sum), nil
}

func findFirstDigits2(value string) string {
	firstIDX := math.MaxInt
	found := ""
	for toFind := range digits {
		idx := strings.Index(value, toFind)
		if idx != -1 && idx < firstIDX {
			firstIDX = idx
			found = toFind
		}
	}
	return digits[found]
}

func findLastDigits2(value string) string {
	lastIDX := 0
	found := ""
	for toFind := range digits {
		idx := strings.LastIndex(value, toFind)
		if idx != -1 && idx >= lastIDX {
			lastIDX = idx
			found = toFind
		}
	}
	return digits[found]
}

var digits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}
