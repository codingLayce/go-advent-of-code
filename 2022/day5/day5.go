package day5

import (
	"adc/common/stack"
	commonstr "adc/common/strings"
	"strconv"
	"strings"
)

type Day5 struct {
	Input string
	Dir   string
}

func New() Day5 {
	return Day5{
		Input: "2022/day5/input.txt",
		Dir:   "2022/day5/",
	}
}

func (d Day5) ProcessPuzzle1(lines []string) (string, error) {
	stacks, movesIdx := parseStacks(lines)

	for _, line := range lines[movesIdx+1:] {
		nbToMove, from, to := parseMove(line)

		for i := 0; i < nbToMove; i++ {
			value := stacks[from].Pop()
			stacks[to].Push(value)
		}
	}

	return getResult(stacks), nil
}

func (d Day5) ProcessPuzzle2(lines []string) (string, error) {
	stacks, movesIdx := parseStacks(lines)

	for _, line := range lines[movesIdx+1:] {
		nbToMove, from, to := parseMove(line)
		value := stacks[from].PopRange(nbToMove)
		stacks[to].PushStack(value)
	}

	return getResult(stacks), nil
}

// getResult returns a string with all the first char in each stacks
func getResult(stacks []stack.Stack[string]) string {
	res := ""
	for _, stack := range stacks {
		res += stack.Pop()
	}
	return res
}

// parseStacks returns the parsed stacks and the index where it stopped
// Basically it translate : `[N] [D]    [A]` --> `ND_A`
// or `        [A]    [B]    ` --> __AB
// Which it's way easier to parse because the stack index corresponds to the index of the char in the string.
func parseStacks(lines []string) ([]stack.Stack[string], int) {
	stacks := make([]stack.Stack[string], 0)
	movesIdx := 0
	for idx, line := range lines {
		if line == "" {
			movesIdx = idx
			break
		}
		replaced := strings.ReplaceAll(line, "    ", "_")
		tweaked := commonstr.RemoveMultiple(replaced, "[", "]", " ")
		for i := 0; i < len(tweaked); i++ {
			if len(stacks) <= i {
				stacks = append(stacks, stack.Stack[string]{})
			}
			if tweaked[i] == uint8('_') {
				continue
			}
			stacks[i].PushBack(string(tweaked[i]))
		}
	}
	return stacks, movesIdx
}

// parseMove returns the nbToMove, from and to parsed from "move x from y to z"
func parseMove(line string) (int, int, int) {
	tweaked := strings.TrimSpace(commonstr.RemoveMultiple(line, "move", "from", "to"))
	arr := strings.Split(tweaked, "  ")
	nbToMove, _ := strconv.Atoi(arr[0])
	from, _ := strconv.Atoi(arr[1])
	to, _ := strconv.Atoi(arr[2])
	return nbToMove, from - 1, to - 1
}
