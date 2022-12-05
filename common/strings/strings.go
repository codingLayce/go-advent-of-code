package strings

import "strings"

func CountOccurrence(str string, sub string) int {
	count := 0

	for i := 0; i < len(str)-len(sub)+1; i++ {
		if str[i:i+len(sub)] == sub {
			count++
		}
	}

	return count
}

func CharAt(str string, index int) string {
	if index >= len(str) {
		return ""
	}

	return string(str[index])
}

func RemoveMultiple(str string, toRemove ...string) string {
	tmp := str
	for _, element := range toRemove {
		tmp = strings.ReplaceAll(tmp, element, "")
	}
	return tmp
}
