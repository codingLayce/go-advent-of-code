package day

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"advent/lib/puzzle"
)

func TestSolvePart1(t *testing.T) {
	answer, err := puzzle.Solve("testdata/example1.txt", SolvePart1)
	require.NoError(t, err)
	assert.Equal(t, "140", answer)

	answer, err = puzzle.Solve("testdata/example2.txt", SolvePart1)
	require.NoError(t, err)
	assert.Equal(t, "1930", answer)

	answer, err = puzzle.Solve("testdata/input.txt", SolvePart1)
	require.NoError(t, err)
	assert.Equal(t, "1483212", answer)
}

func TestSolvePart2(t *testing.T) {
	answer, err := puzzle.Solve("testdata/example1.txt", SolvePart2)
	require.NoError(t, err)
	assert.Equal(t, "80", answer)

	answer, err = puzzle.Solve("testdata/example3.txt", SolvePart2)
	require.NoError(t, err)
	assert.Equal(t, "436", answer)

	answer, err = puzzle.Solve("testdata/example4.txt", SolvePart2)
	require.NoError(t, err)
	assert.Equal(t, "236", answer)

	answer, err = puzzle.Solve("testdata/example5.txt", SolvePart2)
	require.NoError(t, err)
	assert.Equal(t, "368", answer)

	answer, err = puzzle.Solve("testdata/example2.txt", SolvePart2)
	require.NoError(t, err)
	assert.Equal(t, "1206", answer)

	answer, err = puzzle.Solve("testdata/input.txt", SolvePart2)
	require.NoError(t, err)
	assert.Equal(t, "897062", answer)
}
