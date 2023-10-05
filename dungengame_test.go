package dungengame

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHolding(t *testing.T) {
	a := make([][]int, 3)
	assert.Equal(t, 1, calculateMinimumHP(a))
}
