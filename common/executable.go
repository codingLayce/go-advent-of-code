package common

type Executable interface {
	ProcessPuzzle1(lines []string) (string, error)
	ProcessPuzzle2(lines []string) (string, error)
}
