package day

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"advent/lib/puzzle"
)

func TestSolvePart1(t *testing.T) {
	tilesWidth = 11
	tilesHeight = 7

	answer, err := puzzle.Solve("testdata/example1.txt", SolvePart1)
	require.NoError(t, err)
	assert.Equal(t, "12", answer)

	tilesWidth = 101
	tilesHeight = 103

	answer, err = puzzle.Solve("testdata/input.txt", SolvePart1)
	require.NoError(t, err)
	assert.Equal(t, "225648864", answer)
}

func TestSolvePart2(t *testing.T) {
	tilesWidth = 101
	tilesHeight = 103

	answer, err := puzzle.Solve("testdata/input.txt", SolvePart2)
	require.NoError(t, err)
	assert.Equal(t, "7847", answer)
}
