package puzzle

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Solver func(*bufio.Scanner) (string, error)

func Solve(filePath string, solver Solver) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewScanner(file)

	now := time.Now()
	answer, err := solver(reader)
	elapsed := time.Since(now)
	if err != nil {
		fmt.Printf("[ERROR] after %s --- %s\n", elapsed.String(), err)
	} else {
		fmt.Printf("[RESULT] after %s --- %s\n", elapsed.String(), answer)
	}

	return answer, err
}
