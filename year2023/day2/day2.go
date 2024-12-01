package day2

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SolvePart1(reader *bufio.Scanner) (string, error) {
	lineCount := 0

	limits := map[string]int{"blue": 14, "green": 13, "red": 12}
	sum := 0

	for reader.Scan() {
		lineCount++

		// Comme les rounds ne m'interessent pas vraiment, a chaque fois qu'on tire une couleur je veux verifier si elle ne depasse pas la limite.
		// J'enleve donc tout le bruit non necessaire dans la ligne.
		line := reader.Text()
		line = line[strings.Index(line, ":")+2:]
		line = strings.ReplaceAll(line, "; ", ",")
		line = strings.ReplaceAll(line, ", ", ",")
		// Ici la ligne ressemble a ca : XX color,XX color,XX color,XX color

		ok := true

		// Je parcours donc chaque couleur et je verifie que la limite n'est pas depassee.
		for _, bag := range strings.Split(line, ",") {
			separatorIdx := strings.Index(bag, " ")
			nbStr := bag[:separatorIdx]
			nb, err := strconv.Atoi(nbStr)
			if err != nil {
				return "", fmt.Errorf("[l.%d] convert %q to int", lineCount, nbStr)
			}

			if nb > limits[bag[separatorIdx+1:]] { // Jeu impossible, on stop.
				ok = false
				break
			}
		}

		if ok {
			sum += lineCount // l'id est le numero de la ligne dans le jeu de donnees
		}
	}

	return fmt.Sprintf("%d", sum), nil
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	lineCount := 0
	sum := 0

	for reader.Scan() {
		lineCount++

		// Je veux a chaque jeu recuperer le nombre de max par couleur qui a ete tire.
		// J'enleve donc tout le bruit non necessaire dans la ligne.
		line := reader.Text()
		line = line[strings.Index(line, ":")+2:]
		line = strings.ReplaceAll(line, "; ", ",")
		line = strings.ReplaceAll(line, ", ", ",")
		// Ici la ligne ressemble a ca : XX color,XX color,XX color,XX color

		// Je parcours donc chaque couleur et je garde le maximum pour chacune d'elles.
		res := map[string]int{"blue": 0, "green": 0, "red": 0}
		for _, bag := range strings.Split(line, ",") {
			separatorIdx := strings.Index(bag, " ")
			nbStr := bag[:separatorIdx]
			nb, err := strconv.Atoi(nbStr)
			if err != nil {
				return "", fmt.Errorf("[l.%d] convert %q to int", lineCount, nbStr)
			}

			if nb > res[bag[separatorIdx+1:]] {
				res[bag[separatorIdx+1:]] = nb
			}
		}

		sum += res["red"] * res["green"] * res["blue"]
	}

	return fmt.Sprintf("%d", sum), nil
}
