package slices

import "math"

func Contains[T comparable](slice []T, element T) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}

	return false
}

func Sum(slice []int) int {
	sum := 0
	for _, e := range slice {
		sum += e
	}
	return sum
}

func Max(slice []int) (int, int) {
	idx := 0
	max := 0
	for i, e := range slice {
		if e > max {
			max = e
			idx = i
		}
	}
	return idx, max
}

func Min(slice []uint64) (int, uint64) {
	idx := 0
	min := uint64(math.MaxUint64)
	for i, e := range slice {
		if e < min {
			min = e
			idx = i
		}
	}
	return idx, min
}
