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
	assert.Equal(t, "55312", answer)

	answer, err = puzzle.Solve("testdata/input.txt", SolvePart1)
	require.NoError(t, err)
	assert.Equal(t, "203953", answer)
}

func TestSolvePart2(t *testing.T) {
	answer, err := puzzle.Solve("testdata/input.txt", SolvePart2)
	require.NoError(t, err)
	assert.Equal(t, "242090118578155", answer)
}
