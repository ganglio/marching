package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {
	f := Entity(func(x, y float64) float64 {
		return x + y
	})

	n := f.Number(BB{P{-1, -1}, P{1, 1}}, 0.0)

	assert.Equal(t, 4, n)
}

func TestMid(t *testing.T) {
	f := Entity(func(x, y float64) float64 {
		return x + y
	})

	m := f.Mid(BB{P{0, 0}, P{1, 1}})

	assert.Equal(t, 1.0, m)
}

func TestTriangles(t *testing.T) {
	f := Entity(func(x, y float64) float64 {
		return x + y
	})

	tr := f.Triangles(BB{P{0, 0}, P{1, 1}}, 0)

	assert.Equal(t, Triangle{0, 3}, tr)
}

func TestCircle(t *testing.T) {
	c := Circle(0, 0, 1)

	assert.Equal(t, 0.0, c(0, 1))
	assert.Equal(t, 0.0, c(0, -1))
	assert.Equal(t, 0.0, c(1, 0))
	assert.Equal(t, 0.0, c(-1, 0))
	assert.Equal(t, 1.0, c(0, 0))
}
