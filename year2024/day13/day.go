package day

import (
	"advent/lib/vector"
	"bufio"
	"fmt"
)

type Game struct {
	A     vector.Vec2
	B     vector.Vec2
	Prize vector.Vec2
}

func SolvePart1(reader *bufio.Scanner) (string, error) {
	games := getInput(reader)

	sum := 0
	for _, game := range games {
		cheapest := -1

		// Test all the possibilities of button press (Because it cannot be pressed more than 100 times each)
		for aPress := 0; aPress < 100; aPress++ {
			for bPress := 0; bPress < 100; bPress++ {
				if game.A.Mul(aPress).Add(game.B.Mul(bPress)) != game.Prize {
					continue
				}
				cost := aPress*3 + bPress*1
				if cheapest == -1 || cheapest > cost {
					cheapest = cost
				}
			}
		}
		if cheapest != -1 {
			sum += cheapest
		}
	}

	return fmt.Sprintf("%d", sum), nil
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	games := getInput(reader)

	sum := 0
	for _, game := range games {
		game.Prize = game.Prize.Add(vector.NewVec2(10000000000000, 10000000000000))

	}

	return fmt.Sprintf("%d", sum), nil
}

func getInput(reader *bufio.Scanner) []Game {
	var games []Game

	for reader.Scan() {
		game := Game{}
		_, _ = fmt.Sscanf(reader.Text(), "Button A: X+%d, Y+%d", &game.A.X, &game.A.Y)
		reader.Scan()
		_, _ = fmt.Sscanf(reader.Text(), "Button B: X+%d, Y+%d", &game.B.X, &game.B.Y)
		reader.Scan()
		_, _ = fmt.Sscanf(reader.Text(), "Prize: X=%d, Y=%d", &game.Prize.X, &game.Prize.Y)
		reader.Scan()
		games = append(games, game)
	}

	return games
}
