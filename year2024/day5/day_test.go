package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"advent/lib/puzzle"
)

func TestSolvePart1(t *testing.T) {
	answer, err := puzzle.Solve("testdata/example1.txt", SolvePart1)
	require.NoError(t, err)
	assert.Equal(t, "143", answer)

	answer, err = puzzle.Solve("testdata/input.txt", SolvePart1)
	require.NoError(t, err)
	assert.Equal(t, "5747", answer)
}

func TestSolvePart2(t *testing.T) {
	answer, err := puzzle.Solve("testdata/example1.txt", SolvePart2)
	require.NoError(t, err)
	assert.Equal(t, "123", answer)

	answer, err = puzzle.Solve("testdata/input.txt", SolvePart2)
	require.NoError(t, err)
	assert.Equal(t, "5502", answer)
}
