package day1

type Day struct {
	Input string
	Dir   string
}

func New() Day {
	return Day{
		Input: "2023/day1/input.txt",
		Dir:   "2023/day1/",
	}
}

func (d Day) ProcessPuzzle1(lines []string) (string, error) {
	return "not implemented", nil
}

func (d Day) ProcessPuzzle2(lines []string) (string, error) {
	return "not implemented", nil
}
