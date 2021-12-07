package world

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

// Triangles returns the triangulation for the current cell
func (w *World) Triangles(b BB, t float64) []Triangle {
	o := []Triangle{}

	for _, f := range w.funcs {
		o = append(o, Triangulations[f.Number(b, t)])
	}

	return o
}
