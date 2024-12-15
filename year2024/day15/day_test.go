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
	assert.Equal(t, "2028", answer)

	answer, err = puzzle.Solve("testdata/example2.txt", SolvePart1)
	require.NoError(t, err)
	assert.Equal(t, "10092", answer)

	answer, err = puzzle.Solve("testdata/input.txt", SolvePart1)
	require.NoError(t, err)
	assert.Equal(t, "1426855", answer)
}

func TestSolvePart2(t *testing.T) {
	answer, err := puzzle.Solve("testdata/example2.txt", SolvePart2)
	require.NoError(t, err)
	assert.Equal(t, "9021", answer)

	/*answer, err = puzzle.Solve("testdata/input.txt", SolvePart2)
	require.NoError(t, err)
	assert.Equal(t, "", answer)*/
}
