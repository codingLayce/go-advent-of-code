package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVec2_Magnitude(t *testing.T) {
	v := NewVec2(4, 3)
	assert.Equal(t, 5, v.Magnitude())
}
