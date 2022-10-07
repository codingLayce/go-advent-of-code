package maps

func ValuesInt[T comparable](m map[T]int) []int {
	var values []int

	for _, value := range m {
		values = append(values, value)
	}

	return values
}
