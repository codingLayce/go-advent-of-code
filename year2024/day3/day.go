package day3

import (
	"bufio"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

func SolvePart1(reader *bufio.Scanner) (string, error) {
	// Regexp to capture following strings : `mul(X,Y)` where X and Y are {1-3} digits.
	r := regexp.MustCompile(`mul\((?P<a>\d{1,3}),(?P<b>\d{1,3})\)`)

	count := 0
	for reader.Scan() {
		res := r.FindAllStringSubmatch(reader.Text(), -1)

		for _, mul := range res {
			a, err := strconv.Atoi(mul[1])
			if err != nil {
				return "", err
			}
			b, err := strconv.Atoi(mul[2])
			if err != nil {
				return "", err
			}
			count += a * b
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	mulCapture := regexp.MustCompile(`mul\((?P<a>\d{1,3}),(?P<b>\d{1,3})\)`)
	doCapture := regexp.MustCompile(`do\(\)`)
	dontCapture := regexp.MustCompile(`don't\(\)`)

	program := ""

	// Merging all lines into one to easily track the on/off.
	for reader.Scan() {
		program += reader.Text()
	}

	count := 0
	donts := dontCapture.FindAllStringIndex(program, -1)
	dos := doCapture.FindAllStringIndex(program, -1)
	muls := mulCapture.FindAllStringIndex(program, -1)

	// Determine all indexes intervals that must be ignored.
	var intervalsToRemove [][]int
	for _, dont := range donts {
		found := false
		for _, do := range dos {
			if do[0] > dont[0] {
				found = true
				intervalsToRemove = append(intervalsToRemove, []int{dont[0], do[0]})
				break
			}
		}
		if !found { // DON'T forget if the program finishes with a don't and no do following.
			intervalsToRemove = append(intervalsToRemove, []int{dont[0], len(program)})
		}
	}

	// Filter the instructions that must be ignored.
	muls = slices.DeleteFunc(muls, func(ints []int) bool {
		for _, interval := range intervalsToRemove {
			if ints[0] >= interval[0] && ints[1] <= interval[1] {
				return true
			}
		}
		return false
	})

	// Simply retrieve the values from the index and multiply them.
	for _, mul := range muls {
		capture := mulCapture.FindStringSubmatch(program[mul[0]:mul[1]])
		a, err := strconv.Atoi(capture[1])
		if err != nil {
			return "", err
		}
		b, err := strconv.Atoi(capture[2])
		if err != nil {
			return "", err
		}
		count += a * b
	}

	return fmt.Sprintf("%d", count), nil
}
