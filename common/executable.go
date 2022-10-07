package common

type Executable interface {
	ProcessPuzzle1(lines []string) (interface{}, error)
	ProcessPuzzle2(lines []string) (interface{}, error)
}
