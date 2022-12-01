package main

import (
	"adc/2022/day1"
	"adc/common"
	"fmt"
	"time"
)

func main() {
	current := day1.New()
	data, err := common.ReadInput(current.Input)
	if err != nil {
		fmt.Printf("Error while reading input: %v\n", err)
		return
	}

	start := time.Now()
	res1, err := current.ProcessPuzzle1(data)
	if err != nil {
		fmt.Printf("Puzzle 1: ERROR %v\n", err)
		return
	}
	fmt.Printf("Puzzle 1: %v (execution in %s)\n", res1, timeTrackStr(start))

	start = time.Now()
	res2, err := current.ProcessPuzzle2(data)
	if err != nil {
		fmt.Printf("Puzzle 1: ERROR %v\n", err)
		return
	}
	fmt.Printf("Puzzle 2: %v (execution in %s)\n", res2, timeTrackStr(start))
}

func timeTrackStr(start time.Time) string {
	value := time.Since(start)

	if value.Microseconds() < 10 {
		return fmt.Sprintf("%d ns", value.Nanoseconds())

	} else if value.Milliseconds() < 10 {
		return fmt.Sprintf("%d µs", value.Microseconds())
	} else if value.Seconds() < 1 {
		return fmt.Sprintf("%d ms", value.Milliseconds())
	}

	return fmt.Sprintf("%d s", value.Seconds())
}
