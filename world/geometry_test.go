package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeoMid(t *testing.T) {
	assert.Equal(t, P{0, 0}, BB{P{-1, -1}, P{1, 1}}.Mid())
	assert.Equal(t, P{0.5, 0.5}, BB{P{0, 0}, P{1, 1}}.Mid())
}

func TestPDiff(t *testing.T) {
	p := P{0, 1}.Diff(P{1, 0})
	assert.Equal(t, P{-1, 1}, p)
}

func TestPLen(t *testing.T) {
	l := P{3, 4}.Len()

	assert.Equal(t, 5.0, l)
}

func TestBBSize(t *testing.T) {
	b := BB{P{-3, 4}, P{5, 7}}
	assert.Equal(t, P{8, 3}, b.Size())
}

func TestBBScale(t *testing.T) {
	b := BB{P{2, 3}, P{3, 4}}
	assert.Equal(t, BB{P{3, 4.5}, P{4.5, 6}}, b.Scale(1.5))
}

func TestEAddV(t *testing.T) {
	e := E{}

	t.Run("Adds one P", func(t *testing.T) {
		i1 := e.AddV(P{1, 2})
		assert.Equal(t, 0, i1)
	})

	t.Run("Adds it again", func(t *testing.T) {
		i2 := e.AddV(P{1, 2})
		assert.Equal(t, 0, i2)
	})

	t.Run("Adds a second one", func(t *testing.T) {
		i3 := e.AddV(P{3, 2})
		assert.Equal(t, 1, i3)
	})
}

func TestEAddE(t *testing.T) {
	e := E{}

	t.Run("Adds one edge", func(t *testing.T) {
		i1 := e.AddE(P{1, 2}, P{2, 1})
		assert.Equal(t, 0, i1)
	})

	t.Run("Adds it again", func(t *testing.T) {
		i2 := e.AddE(P{1, 2}, P{2, 1})
		assert.Equal(t, 0, i2)
	})

	t.Run("Adds a second one", func(t *testing.T) {
		i3 := e.AddE(P{3, 2}, P{3, 2})
		assert.Equal(t, 1, i3)
	})
}

func TestEProject(t *testing.T) {
	e := E{}
	e.AddE(P{-1, -1}, P{0, 1})
	e.AddE(P{0, 1}, P{1, -1})

	r := E{}
	r.AddE(P{-3, -3}, P{0, 3})
	r.AddE(P{0, 3}, P{3, -3})

	p := e.Project(BB{P{-1, -1}, P{1, 1}}, BB{P{-3, -3}, P{3, 3}})

	assert.Equal(t, r, p)
}
