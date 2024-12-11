package day

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SolvePart1(reader *bufio.Scanner) (string, error) {
	reader.Scan()

	// Naive approach using an array and making the whole simulation.
	sum := 0
	for _, stone := range strings.Split(reader.Text(), " ") {
		nb, _ := strconv.Atoi(stone)
		stones := []int{nb}

		for i := 0; i < 25; i++ {
			var newStones []int
			for _, st := range stones {
				if st == 0 {
					newStones = append(newStones, 1)
					continue
				}
				str := fmt.Sprintf("%d", st)
				if len(str)%2 == 0 {
					leftNb, _ := strconv.Atoi(str[:len(str)/2])
					newStones = append(newStones, leftNb)
					rightNb, _ := strconv.Atoi(str[len(str)/2:])
					newStones = append(newStones, rightNb)
					continue
				}
				newStones = append(newStones, st*2024)
			}
			stones = newStones
		}
		sum += len(stones)
	}

	return fmt.Sprintf("%d", sum), nil
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	reader.Scan()

	// Better solution using a map that keep track of the number of times a stone appears.
	// So if a stone appears 100 times, instead of computing his result 100 times and storing 100 (or 200 when splitting) individuals int
	// I only store 2 (the stone + the number of times it appears)
	stones := make(map[int]int)
	for _, stone := range strings.Split(reader.Text(), " ") {
		nb, _ := strconv.Atoi(stone)
		stones[nb] = 1 // There is no duplication in the input
	}

	for i := 0; i < 75; i++ {
		newStones := make(map[int]int)
		for stone, nb := range stones {
			if stone == 0 {
				newStones[1] += nb
				continue
			}
			str := fmt.Sprintf("%d", stone)
			if len(str)%2 == 0 {
				leftNb, _ := strconv.Atoi(str[:len(str)/2])
				newStones[leftNb] += nb
				rightNb, _ := strconv.Atoi(str[len(str)/2:])
				newStones[rightNb] += nb
				continue
			}
			newStones[stone*2024] += nb
		}
		stones = newStones
	}

	sum := 0
	for _, value := range stones {
		sum += value
	}

	return fmt.Sprintf("%d", sum), nil
}
