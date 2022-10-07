package main

import (
	"adc/2018/day2"
	"adc/common"
	"fmt"
)

func main() {
	current := day2.New()
	data, err := common.ReadInput(current.Input)
	if err != nil {
		fmt.Printf("Error while reading input: %v\n", err)
		return
	}

	res1, err := current.ProcessPuzzle1(data)
	if err != nil {
		fmt.Printf("Puzzle 1: ERROR %v\n", err)
		return
	}
	fmt.Printf("Puzzle 1: %v\n", res1)

	res2, err := current.ProcessPuzzle2(data)
	if err != nil {
		fmt.Printf("Puzzle 1: ERROR %v\n", err)
		return
	}
	fmt.Printf("Puzzle 2: %v\n", res2)
}
