package strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCountOccurrence(t *testing.T) {
	for name, tt := range map[string]struct {
		value     string
		searching string
		expected  int
	}{
		"Single char twice": {
			"abda",
			"a",
			2,
		},
		"Complex string twice": {
			"aklhdsqioijiqsdjioqshdqsdisquhzduhq hioqqsdshid zqiodhiqzhdsjqsdqhd zqudhqkjsdqsdqsd",
			"qsd",
			6,
		},
	} {
		tc := tt // Use local variable to avoid closure-caused race condition
		t.Run(name, func(t *testing.T) {
			count := CountOccurrence(tc.value, tc.searching)
			require.Equal(t, tc.expected, count)
		})
	}
}
