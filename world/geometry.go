package world

import (
	"math"
)

const eps = 1e-5

// P is a 2D point
type P struct {
	X float64
	Y float64
}

// Diff calculates the difference between two P
func (a P) Diff(b P) P {
	return P{a.X - b.X, a.Y - b.Y}
}

// Len calculates the length of a P
func (a P) Len() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y)
}

// BB is a bounding box (a rectangle too...)
type BB struct {
	P1 P
	P2 P
}

// Size returns the width and height of a BB
func (b BB) Size() P {
	return P{b.P2.X - b.P1.X, b.P2.Y - b.P1.Y}
}

// Mid returns the mid point of a BB
func (b BB) Mid() P {
	return P{(b.P1.X + b.P2.X) / 2.0, (b.P1.Y + b.P2.Y) / 2.0}
}

// I is the starting and ending indexes of an edge
type I struct {
	S int
	E int
}

// E is a set of unique edges. Each edge is comprised of the two indexes from the enclosed V object.
type E struct {
	V []P // A set of unique vertexes
	E []I // A set of unique edges defined as starting and ending vertexes using the indexes of those vertexes
}

// AddV adds a new vertex and returns the index. If the vertex is already there it just returns the index.
func (e *E) AddV(p P) int {
	for k, v := range e.V {
		if v.Diff(p).Len() < eps {
			return k
		}
	}
	e.V = append(e.V, p)
	return len(e.V) - 1
}

// AddE adds a new edge and returns the index. It first adds start and end vertexes to the vertex list then adds the edge and returns the index.
func (e *E) AddE(a P, b P) int {
	i1 := e.AddV(a)
	i2 := e.AddV(b)

	for k, v := range e.E {
		if (v.S == i1 && v.E == i2) || (v.S == i2 && v.E == i1) {
			return k
		}
	}
	e.E = append(e.E, I{i1, i2})
	return len(e.E) - 1
}

// Project projects E from one BB to another
func (e E) Project(a BB, b BB) E {
	t := make([]P, len(e.V))

	// Normalize first
	s := a.Size()
	for k, p := range e.V {
		t[k] = P{(p.X - a.P1.X) / s.X, (p.Y - a.P1.Y) / s.Y}
	}

	s = b.Size()
	for k, p := range t {
		t[k] = P{b.P1.X + s.X*p.X, b.P1.Y + s.Y*p.Y}
	}

	// Copy, Assign and Return
	o := e
	o.V = t

	return o
}
