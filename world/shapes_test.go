package world

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircle(t *testing.T) {
	c := Circle(0, 0, 1)

	assert.Equal(t, 1.0, c(0, 1))
	assert.Equal(t, 1.0, c(0, -1))
	assert.Equal(t, 1.0, c(1, 0))
	assert.Equal(t, 1.0, c(-1, 0))
	assert.Equal(t, math.Inf(1), c(0, 0))
}

func TestBox(t *testing.T) {
	b := Box(0, 0, 1, 1, 30)

	assert.Greater(t, eps, math.Abs(b(-0.5, -1)-1.0))
	assert.Equal(t, 1.0, b(0, -1))
	assert.Greater(t, eps, math.Abs(b(0.5, -1)-1.0))

	assert.Greater(t, eps, math.Abs(b(-0.5, 1)-1.0))
	assert.Equal(t, 1.0, b(0, 1))
	assert.Greater(t, eps, math.Abs(b(0.5, 1)-1.0))

	assert.Greater(t, eps, math.Abs(b(-1, -0.5)-1.0))
	assert.Equal(t, 1.0, b(-1, 0))
	assert.Greater(t, eps, math.Abs(b(-1, 0.5)-1.0))

	assert.Greater(t, eps, math.Abs(b(1, -0.5)-1.0))
	assert.Equal(t, 1.0, b(1, 0))
	assert.Greater(t, eps, math.Abs(b(1, 0.5)-1.0))
}
