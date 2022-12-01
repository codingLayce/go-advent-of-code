package common

import (
	"fmt"
	"os"
)

func WriteResults(dir string, res, exec [2]string) error {
	content := fmt.Sprintf("PUZZLE 1: %s (executed in %s)\nPUZZLE 2: %s (executed in %s)", res[0], exec[0], res[1], exec[1])

	f, err := os.OpenFile(fmt.Sprintf("%sresults.txt", dir), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		return err
	}

	return nil
}
