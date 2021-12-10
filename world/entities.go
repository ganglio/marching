package world

import (
	"math"
)

// Entity an entity in the world.
type Entity func(x float64, y float64) float64

// Number returns the Triangulation Number for the given Entity, Bouding box and Threshold
func (e Entity) Number(b BB, t float64) int {
	var v0, v1, v2, v3 int
	if e(b.P1.X, b.P1.Y) > t {
		v0 = 1
	}
	if e(b.P2.X, b.P1.Y) > t {
		v1 = 2
	}
	if e(b.P2.X, b.P2.Y) > t {
		v2 = 4
	}
	if e(b.P1.X, b.P2.Y) > t {
		v3 = 8
	}
	return v0 + v1 + v2 + v3
	// TODO Handle Saddles
}

// Val returns the value of an entity in P
func (e Entity) Val(p P) float64 {
	return e(p.X, p.Y)
}

// Mid returns the value of the function in the middle point of the cell. Used to distinguish saddle points
func (e Entity) Mid(b BB) float64 {
	return e.Val(b.Mid())
}

// Triangles returns the triangulation for the current BB and threshold
func (e Entity) Triangles(b BB, t float64) Triangle {
	return Triangulations[e.Number(b, t)]
}

// Rot applies a teta radiant rotation to the entity
func (e Entity) Rot(teta float64) Entity {
	return func(x, y float64) float64 {
		u := x*math.Cos(teta) - y*math.Sin(teta)
		v := x*math.Sin(teta) + y*math.Cos(teta)
		return e(u, v)
	}
}

// Scale applies an alpha scaling factor to the entity
func (e Entity) Scale(alpha float64) Entity {
	return func(x, y float64) float64 {
		return e(x/alpha, y/alpha)
	}
}
