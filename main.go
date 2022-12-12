package main

import (
	"fmt"
	"time"

	"adc/2022/day12"
	"adc/common"
)

func main() {
	current := day12.New()
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
	exec1 := timeTrackStr(start)
	fmt.Printf("Puzzle 1: %v (executed in %v)\n", res1, exec1)

	start = time.Now()
	res2, err := current.ProcessPuzzle2(data)
	if err != nil {
		fmt.Printf("Puzzle 1: ERROR %v\n", err)
		return
	}
	exec2 := timeTrackStr(start)
	fmt.Printf("Puzzle 2: %v (executed in %v)\n", res2, exec2)

	_ = common.WriteResults(current.Dir, [2]string{res1, res2}, [2]string{exec1, exec2})
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
