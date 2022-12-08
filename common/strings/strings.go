package strings

import (
	"strconv"
	"strings"
)

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

func ToIntSlice(str string) []int {
	var arr []int

	for _, e := range str {
		value, _ := strconv.Atoi(string(e))
		arr = append(arr, value)
	}

	return arr
}