package world

import (
	log "github.com/sirupsen/logrus"
)

// World describes the entities in the current world
type World struct {
	funcs []Entity
}

// NewWorld creates a new World instance
func NewWorld(f ...Entity) *World {
	return &World{f}
}

// Add adds an new entity to the world
func (w *World) Add(f Entity) {
	w.funcs = append(w.funcs, f)
}

// Len returns the number of entities in the world
func (w *World) Len() int {
	return len(w.funcs)
}

// Val returns the value of the World at the provided coordinates
func (w *World) Val(x, y float64) float64 {
	o := 0.0
	for _, f := range w.funcs {
		o += f(x, y)
	}
	return o
}

// Vals returns an array with the values of all the entries in the world
func (w *World) Vals(x, y float64) []float64 {
	o := []float64{}

	for _, f := range w.funcs {
		o = append(o, f(x, y))
	}

	return o
}

// Triangles returns the triangulation for the current BB
func (w *World) Triangles(b BB, t float64) []Triangle {
	o := []Triangle{}

	for _, f := range w.funcs {
		o = append(o, Triangulations[f.Number(b, t)])
	}

	return o
}

// Geometries returns the list of vertexes and edges for each entity of World in the BB calculated using an HxV grid
func (w *World) Geometries(bb BB, h int, v int, t float64) []E {
	o := make([]E, len(w.funcs))

	s := bb.Size()
	log.Infof("S: %#v", s)

	for j := 0; j < v; j++ {
		for i := 0; i < h; i++ {
			b := BB{
				P{
					bb.P1.X + float64(i)*s.X/float64(h),
					bb.P1.Y + float64(j)*s.Y/float64(v),
				},
				P{
					bb.P1.X + float64(i+1)*s.X/float64(h),
					bb.P1.Y + float64(j+1)*s.Y/float64(v),
				},
			}

			trs := w.Triangles(b, t)
			for k, tr := range trs {
				ps := tr.ToP(b)
				if len(ps) > 0 { // TODO handle inside/outside cases
					for ind, p := range ps {
						if ind%2 == 0 {
							o[k].AddE(p, ps[ind+1])
						}
					}
				}
			}
		}
	}
	return o
}
