package dayx

type Day struct {
	Input string
	Dir   string
}

func New() Day {
	return Day{
		Input: "202x/dayx/input.txt",
		Dir:   "202x/dayx/",
	}
}

func (d Day) ProcessPuzzle1(lines []string) (string, error) {
	return "not implemented", nil
}

func (d Day) ProcessPuzzle2(lines []string) (string, error) {
	return "not implemented", nil
}
