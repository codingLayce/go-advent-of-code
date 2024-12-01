package day1

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SolvePart1(reader *bufio.Scanner) (string, error) {
	sum := 0
	lineCount := 0
	for reader.Scan() {
		lineCount++

		digits := digitsOrdered(reader.Text())
		nbStr := fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1]) // Takes the left most and the right most digits
		nb, err := strconv.Atoi(nbStr)
		if err != nil {
			return "", fmt.Errorf("[l.%d] cannot convert %s into an int !", lineCount, nbStr)
		}
		sum += nb
	}

	return fmt.Sprintf("%d", sum), nil
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	sum := 0
	lineCount := 0
	for reader.Scan() {
		lineCount++

		digits := digitsOrdered2(reader.Text())
		nbStr := fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1]) // Takes the left most and the right most digits
		nb, err := strconv.Atoi(nbStr)
		if err != nil {
			return "", fmt.Errorf("[l.%d] cannot convert %s into an int !", lineCount, nbStr)
		}
		sum += nb
	}

	return fmt.Sprintf("%d", sum), nil
}

// digitsOrdered returns a slice of int ordered by index from the line string.
// It looks for digits in the string such as '1', '2', '3', ... '9'
func digitsOrdered(line string) []int {
	var numbers []int

	for _, ch := range line {
		if ch >= '1' && ch <= '9' {
			numbers = append(numbers, int(ch-'1'+1)) // ASCII manipulation: for a '2' it's like doing: 50 (2 in ASCII) - 49 (1 in ASCII) + 1 = 2 (actual decimal)
		}
	}
	return numbers
}

// digitsOrdered2 returns a slice of int ordered by index from the line string.
// It looks for digits in the string such as '1', '2', '3', ... '9' and also spelled numbers 'one', 'two', ... 'nine'.
func digitsOrdered2(line string) []int {
	var oredered []int
	lookup := strings.Builder{}

	for i := 0; i < len(line); i++ {
		lookup.Reset()
		for j := i; j < len(line); j++ {
			lookup.WriteByte(line[j])
			value, ok := numbers[lookup.String()]
			if ok {
				oredered = append(oredered, value)
				break
			}
		}
	}

	return oredered
}

var numbers = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
