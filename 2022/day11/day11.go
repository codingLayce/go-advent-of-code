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
		Input: "2022/day11/input.txt",
		Dir:   "2022/day11/",
	}
}

type Monkey struct {
	items        []uint64
	inspected    int
	op           func(current uint64) uint64
	worriedLevel func(current uint64) uint64
	test         func(current uint64) int
	modulo       int
}

// return newValue, index
func (m *Monkey) inspect(lambda func(uint64) uint64) (uint64, int) {
	item := m.items[0]
	m.items = m.items[1:]
	m.inspected++

	worryLevel := m.op(item)
	worryLevel = lambda(worryLevel)

	return worryLevel, m.test(worryLevel)
}

func (d Day11) ProcessPuzzle1(lines []string) (string, error) {
	monkeys := parseMonkeys(lines)

	inspected := process(monkeys, 20, func(u uint64) uint64 {
		return u / 3
	})

	sort.Ints(inspected)

	return fmt.Sprintf("%d", inspected[len(inspected)-1]*inspected[len(inspected)-2]), nil
}

func (d Day11) ProcessPuzzle2(lines []string) (string, error) {
	monkeys := parseMonkeys(lines)

	// Pour la seconde partie, le problème qui va se poser est que les nombres seront beaucoup trop grands pour être stocké
	// Il faut donc trouver un moyen de réduire l'"item" sans compromettre le test (divisible).
	// Il faut donc trouver le plus petit diviseur common des tests.
	// Or ceux-ci sont tous des nombres premier, du coup le plus petit diviseur common est le produit de ceux-ci.
	lcm := monkeys[0].modulo
	for i := 1; i < len(monkeys); i++ {
		lcm *= monkeys[i].modulo
	}

	inspected := process(monkeys, 10000, func(u uint64) uint64 {
		return u % uint64(lcm)
	})
	sort.Ints(inspected)

	return fmt.Sprintf("%d", inspected[len(inspected)-1]*inspected[len(inspected)-2]), nil
}

func process(monkeys []Monkey, rounds int, lambda func(uint64) uint64) []int {
	for i := 0; i < rounds; i++ {
		for idx := range monkeys {
			for len(monkeys[idx].items) != 0 {
				item, toThrow := monkeys[idx].inspect(lambda)
				monkeys[toThrow].items = append(monkeys[toThrow].items, item)
			}
		}
	}

	var inspected []int
	for _, monkey := range monkeys {
		inspected = append(inspected, monkey.inspected)
	}

	return inspected
}

func parseMonkeys(lines []string) []Monkey {
	var monkeys []Monkey

	for i := 0; i <= len(lines)/7; i++ {
		idx := i * 7
		tstFunc, modulo := parseTest(lines[idx+3], lines[idx+4], lines[idx+5])
		monkey := Monkey{
			items:  parseStartingItems(lines[idx+1]),
			op:     parseOperation(lines[idx+2]),
			test:   tstFunc,
			modulo: modulo,
		}
		monkeys = append(monkeys, monkey)
	}
	return monkeys
}

func parseTest(testL, trueL, falseL string) (func(uint64) int, int) {
	divisible, _ := strconv.Atoi(testL[strings.LastIndex(testL, " ")+1:])
	trueMonkey, _ := strconv.Atoi(trueL[strings.LastIndex(trueL, " ")+1:])
	falseMonkey, _ := strconv.Atoi(falseL[strings.LastIndex(falseL, " ")+1:])

	return func(current uint64) int {
		if current%uint64(divisible) == 0 {
			return trueMonkey
		}
		return falseMonkey
	}, divisible
}

func parseOperation(line string) func(current uint64) uint64 {
	i := strings.Index(line, "old ") + 4
	operation := string(line[i])
	value, err := strconv.ParseUint(line[i+2:], 10, 64)
	if err != nil {
		return func(current uint64) uint64 {
			switch operation {
			case "+":
				return current + current
			case "-":
				return current - current
			case "*":
				return current * current
			}

			return current
		}
	}

	return func(current uint64) uint64 {
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

func parseStartingItems(line string) []uint64 {
	i := strings.Index(line, ":") + 2
	arr := strings.Split(line[i:], ", ")
	var items []uint64
	for _, item := range arr {
		value, _ := strconv.ParseUint(item, 10, 64)
		items = append(items, value)
	}
	return items
}
