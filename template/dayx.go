package dayx

type Dayx struct {
	Input string
	Dir   string
}

func New() Dayx {
	return Dayx{
		Input: "202x/dayx/input.txt",
		Dir:   "202x/dayx/",
	}
}

func (d Dayx) ProcessPuzzle1(lines []string) (string, error) {
	return "not implemented", nil
}

func (d Dayx) ProcessPuzzle2(lines []string) (string, error) {
	return "not implemented", nil
}
