package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddLenVal(t *testing.T) {
	w := NewWorld()

	w.Add(func(x, y float64) float64 {
		return x + y
	})
	w.Add(func(x, y float64) float64 {
		return x * y
	})
	w.Add(func(x, y float64) float64 {
		return x / y
	})

	assert.Equal(t, 3, w.Len())
	assert.Equal(t, 5.5, w.Val(1, 2))
}

func TestLenVal(t *testing.T) {
	w := NewWorld(
		func(x, y float64) float64 {
			return x + y
		},
		func(x, y float64) float64 {
			return x * y
		},
		func(x, y float64) float64 {
			return x / y
		},
	)

	assert.Equal(t, 3, w.Len())
	assert.Equal(t, 5.5, w.Val(1, 2))
}

func TestVals(t *testing.T) {
	w := NewWorld(
		func(x, y float64) float64 {
			return x + y
		},
		func(x, y float64) float64 {
			return x * y
		},
		func(x, y float64) float64 {
			return x / y
		},
	)

	vv := w.Vals(1, 2)
	assert.Equal(t, []float64{3.0, 2.0, 0.5}, vv)
}

func TestWorldTriangles(t *testing.T) {
	w := NewWorld(
		func(x, y float64) float64 {
			return x + y
		},
		func(x, y float64) float64 {
			return x * y
		},
		func(x, y float64) float64 {
			return x / y
		},
	)

	tt := w.Triangles(BB{P{-1, -1}, P{1, 1}}, 0)
	assert.Equal(t, []Triangle{{1, 2}, {0, 1, 2, 3}, {0, 1, 2, 3}}, tt)
}

func TestWorldGeometries(t *testing.T) {
	b := BB{P{-3, -3}, P{3, 3}}

	t.Run("line", func(t *testing.T) {
		w := NewWorld(
			func(x, y float64) float64 {
				return x + y
			},
		)
		g := w.Geometries(b, 3, 3, -1)
		e := []E{
			{
				V: []P{
					{1, -2},
					{0, -1},
					{2, -3},
					{-1, 0},
					{-2, 1},
					{-3, 2},
				},
				E: []I{
					{0, 1},
					{2, 0},
					{3, 4},
					{1, 3},
					{4, 5},
				},
			},
		}
		assert.Equal(t, e, g)
	})

	t.Run("circle", func(t *testing.T) {
		w := NewWorld(Circle(0, 0, 2))
		g := w.Geometries(b, 3, 3, 0)
		e := []E{
			{
				V: []P{
					{-1, -2},
					{-2, -1},
					{1, -2},
					{2, -1},
					{-2, 1},
					{2, 1},
					{-1, 2},
					{1, 2},
				},
				E: []I{
					{0, 1},
					{2, 0},
					{3, 2},
					{1, 4},
					{3, 5},
					{4, 6},
					{7, 6},
					{5, 7},
				},
			},
		}
		assert.Equal(t, e, g)
	})
}

func BenchWorldTrianglesN(tot int, b *testing.B) {
	w := NewWorld()
	for t := 0; t < tot; t++ {
		w.Add(func(x, y float64) float64 {
			return x + y
		})
	}

	for n := 0; n < b.N; n++ {
		w.Triangles(BB{P{-1, -1}, P{1, 1}}, 0)
	}
}

func BenchmarkWorldTriangles1(b *testing.B)       { BenchWorldTrianglesN(1, b) }
func BenchmarkWorldTriangles10(b *testing.B)      { BenchWorldTrianglesN(10, b) }
func BenchmarkWorldTriangles100(b *testing.B)     { BenchWorldTrianglesN(100, b) }
func BenchmarkWorldTriangles1000(b *testing.B)    { BenchWorldTrianglesN(1000, b) }
func BenchmarkWorldTriangles10000(b *testing.B)   { BenchWorldTrianglesN(10000, b) }
func BenchmarkWorldTriangles100000(b *testing.B)  { BenchWorldTrianglesN(100000, b) }
func BenchmarkWorldTriangles1000000(b *testing.B) { BenchWorldTrianglesN(1000000, b) }
