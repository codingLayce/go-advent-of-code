package day4

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolvePart1(reader *bufio.Scanner) (string, error) {
	// Les regles sont la liste des valeurs devant se trouver apres la clef.
	rules := getRules(reader)

	// Extraction des lignes valides
	var valids [][]string
	for reader.Scan() {
		line := strings.Split(reader.Text(), ",")
		if checkLine(rules, line) {
			valids = append(valids, line)
		}
	}

	// Somme des valeurs du milieu de chaque ligne
	sum, err := sumMiddleValues(valids)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", sum), nil
}

func SolvePart2(reader *bufio.Scanner) (string, error) {
	// Les regles sont la liste des valeurs devant se trouver apres la clef.
	rules := getRules(reader)

	// Extraction des lignes invalides
	var invalids [][]string
	for reader.Scan() {
		line := strings.Split(reader.Text(), ",")
		if !checkLine(rules, line) {
			invalids = append(invalids, line)
		}
	}

	// Pour chaque ligne invalide, j'applique la function de sort de go en utilisant les regles recuperees.
	for _, invalid := range invalids {
		slices.SortFunc(invalid, func(a, b string) int {
			idx := slices.Index(rules[a], b)
			if idx != -1 { // b doit etre apres a
				return -1
			}
			idx = slices.Index(rules[b], a)
			if idx != -1 { // b doit etre avant a
				return 1
			}
			return 0 // pas de regle specifique
		})
	}

	// Somme des valeurs du milieu de chaque ligne
	sum, err := sumMiddleValues(invalids)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", sum), nil
}

func checkLine(rules map[string][]string, line []string) bool {
	for idx, value := range line {
		for _, rul := range rules[value] {
			rulIdx := slices.Index(line, rul)
			if rulIdx != -1 && rulIdx < idx {
				return false
			}
		}
	}
	return true
}

func getRules(reader *bufio.Scanner) map[string][]string {
	rules := make(map[string][]string)

	for reader.Scan() && reader.Text() != "" {
		arr := strings.Split(reader.Text(), "|")
		_, ok := rules[arr[0]]
		if ok {
			rules[arr[0]] = append(rules[arr[0]], arr[1])
		} else {
			rules[arr[0]] = []string{arr[1]}
		}
	}
	return rules
}

func sumMiddleValues(slice [][]string) (int, error) {
	count := 0
	for _, valid := range slice {
		v, err := strconv.Atoi(valid[len(valid)/2])
		if err != nil {
			return 0, err
		}
		count += v
	}
	return count, nil
}
