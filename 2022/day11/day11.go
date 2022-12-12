package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day11 struct {
	Input string
	Dir   string
}

func New() Day11 {
	return Day11{
		Input: "2022/day11/example.txt",
		Dir:   "2022/day11/",
	}
}

type Monkey struct {
	items     []int
	inspected int
	op        func(current int) int
	test      func(current int) int
}

func (m *Monkey) String() string {
	return fmt.Sprintf("<items:%v | inspected:%d>", m.items, m.inspected)
}

// return newValue, index
func (m *Monkey) inspect() (int, int) {
	item := m.items[0]
	m.items = m.items[1:]
	m.inspected++

	worryLevel := m.op(item)
	worryLevel /= 3

	return worryLevel, m.test(worryLevel)
}

func (d Day11) ProcessPuzzle1(lines []string) (string, error) {
	monkeys := parseMonkeys(lines)

	for i := 0; i < 20; i++ {
		for idx, _ := range monkeys {
			for len(monkeys[idx].items) != 0 {
				item, toThrow := monkeys[idx].inspect()
				monkeys[toThrow].items = append(monkeys[toThrow].items, item)
			}
		}

		for _, monkey := range monkeys {
			fmt.Printf("%s\n", monkey.String())
		}
		fmt.Println()
	}

	var inspected []int
	for _, monkey := range monkeys {
		inspected = append(inspected, monkey.inspected)
	}

	sort.Ints(inspected)

	return fmt.Sprintf("%d", inspected[len(inspected)-1]*inspected[len(inspected)-2]), nil
}

func (d Day11) ProcessPuzzle2(lines []string) (string, error) {
	return "not implemented", nil
}

func parseMonkeys(lines []string) []Monkey {
	var monkeys []Monkey

	for i := 0; i <= len(lines)/7; i++ {
		idx := i * 7
		monkey := Monkey{
			items: parseStartingItems(lines[idx+1]),
			op:    parseOperation(lines[idx+2]),
			test:  parseTest(lines[idx+3], lines[idx+4], lines[idx+5]),
		}
		monkeys = append(monkeys, monkey)
	}
	return monkeys
}

func parseTest(testL, trueL, falseL string) func(current int) int {
	divisible, _ := strconv.Atoi(testL[strings.LastIndex(testL, " ")+1:])
	trueMonkey, _ := strconv.Atoi(trueL[strings.LastIndex(trueL, " ")+1:])
	falseMonkey, _ := strconv.Atoi(falseL[strings.LastIndex(falseL, " ")+1:])

	return func(current int) int {
		if current%divisible == 0 {
			return trueMonkey
		}
		return falseMonkey
	}
}

func parseOperation(line string) func(current int) int {
	i := strings.Index(line, "old ") + 4
	operation := string(line[i])
	value, err := strconv.Atoi(line[i+2:])
	if err != nil {
		value = -1
	}

	return func(current int) int {
		if value == -1 {
			value = current
		}

		switch operation {
		case "+":
			return current + value
		case "-":
			return current - value
		case "*":
			return current * value
		}

		return current
	}
}

func parseStartingItems(line string) []int {
	i := strings.Index(line, ":") + 2
	arr := strings.Split(line[i:], ", ")
	var items []int
	for _, item := range arr {
		value, _ := strconv.Atoi(item)
		items = append(items, value)
	}
	return items
}
