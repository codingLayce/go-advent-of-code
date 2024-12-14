package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPgcd(t *testing.T) {
	for name, tc := range map[string]struct {
		a    int
		b    int
		pgcd int
	}{
		"10 20":                    {a: 10, b: 20, pgcd: 10},
		"3 2":                      {a: 3, b: 2, pgcd: 1},
		"934859743 10000000000000": {a: 934859742, b: 10000000000000, pgcd: 2},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.pgcd, Pgcd(tc.a, tc.b))
		})
	}
}
