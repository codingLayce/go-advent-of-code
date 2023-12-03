package slices

import "math"

func Reverse[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func Contains[T comparable](slice []T, element T) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}

	return false
}

func HasDuplicates[T comparable](slice []T) bool {
	cache := make(map[T]struct{})
	for _, value := range slice {
		if _, ok := cache[value]; ok {
			return true
		}
		cache[value] = struct{}{}
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

func MinInt(slice []int) (int, int) {
	idx := 0
	min := math.MaxInt
	for i, e := range slice {
		if e < min {
			min = e
			idx = i
		}
	}
	return idx, min
}

func MinUint(slice []uint64) (int, uint64) {
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
