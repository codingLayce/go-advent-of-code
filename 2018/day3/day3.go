package day3

import (
	"adc/common/math"
	"strconv"
	"strings"
)

type Day3 struct {
	Input string
}

type square struct {
	id              string
	usedCoordinates []math.Vector2
}

func newSquareFromString(str string) (square, error) {
	s := square{}

	arr := strings.Split(str, " ")
	s.id = arr[0][1:]

	arrr := strings.Split(arr[2], ",")

	startX, err := strconv.Atoi(arrr[0])
	if err != nil {
		return square{}, err
	}
	startY, err := strconv.Atoi(arrr[1][:len(arrr[1])-1])
	if err != nil {
		return square{}, err
	}

	arrrr := strings.Split(arr[3], "x")
	width, err := strconv.Atoi(arrrr[0])
	if err != nil {
		return square{}, err
	}
	height, err := strconv.Atoi(arrrr[0])
	if err != nil {
		return square{}, err
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			s.usedCoordinates = append(s.usedCoordinates, math.NewVector2(startX+x, startY+y))
		}
	}

	return s, nil
}

func New() Day3 {
	return Day3{
		Input: "2018/day3/input.txt",
	}
}

func (d Day3) ProcessPuzzle1(lines []string) (interface{}, error) {
	var squares []square

	for _, line := range lines {
		s, err := newSquareFromString(line)
		if err != nil {
			return nil, err
		}
		squares = append(squares, s)
	}

	mapping := make(map[string]int)

	for _, s := range squares {
		for _, coordinate := range s.usedCoordinates {
			str := coordinate.ToString()
			if _, ok := mapping[str]; !ok {
				mapping[str] = 1
			} else {
				mapping[str]++
			}
		}
	}

	count := 0
	for _, v := range mapping {
		if v > 1 {
			count++
		}
	}

	return count, nil
}

func (d Day3) ProcessPuzzle2(lines []string) (interface{}, error) {
	return "not implemented", nil
}
