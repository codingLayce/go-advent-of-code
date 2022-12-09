package math

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDistanceManathan(t *testing.T) {
	for name, tt := range map[string]struct {
		a                Vector2
		b                Vector2
		expectedDistance int
	}{
		"1 right": {
			a:                NewVector2(0, 0),
			b:                NewVector2(1, 0),
			expectedDistance: 1,
		},
		"2 right": {
			a:                NewVector2(0, 0),
			b:                NewVector2(2, 0),
			expectedDistance: 2,
		},
		"1 left": {
			a:                NewVector2(0, 0),
			b:                NewVector2(-1, 0),
			expectedDistance: 1,
		},
		"2 left": {
			a:                NewVector2(0, 0),
			b:                NewVector2(-2, 0),
			expectedDistance: 2,
		},
		"1 top right": {
			a:                NewVector2(0, 0),
			b:                NewVector2(1, 1),
			expectedDistance: 2,
		},
	} {
		tc := tt // Use local variable to avoid closure-caused race condition
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expectedDistance, tc.a.DistanceManathan(tc.b))
		})
	}
}

func TestDistance(t *testing.T) {
	for name, tt := range map[string]struct {
		a                Vector2
		b                Vector2
		expectedDistance int
	}{
		"1 right": {
			a:                NewVector2(0, 0),
			b:                NewVector2(1, 0),
			expectedDistance: 1,
		},
		"2 right": {
			a:                NewVector2(0, 0),
			b:                NewVector2(2, 0),
			expectedDistance: 2,
		},
		"1 left": {
			a:                NewVector2(0, 0),
			b:                NewVector2(-1, 0),
			expectedDistance: 1,
		},
		"2 left": {
			a:                NewVector2(0, 0),
			b:                NewVector2(-2, 0),
			expectedDistance: 2,
		},
		"1 top right": {
			a:                NewVector2(0, 0),
			b:                NewVector2(1, 1),
			expectedDistance: 1,
		},
		"2 top right": {
			a:                NewVector2(0, 0),
			b:                NewVector2(1, 2),
			expectedDistance: 2,
		},
	} {
		tc := tt // Use local variable to avoid closure-caused race condition
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expectedDistance, tc.a.DistanceManathan(tc.b))
		})
	}
}
