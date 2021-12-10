package world

// Triangle defines the triangulation for a cell
type Triangle []int

// ToP convert the current triangle to a slice of Ps given the provided BB
func (t Triangle) ToP(b BB) []P {
	m := b.Mid()
	o := []P{}

	for _, tr := range t {
		switch tr {
		case 0:
			o = append(o, P{m.X, b.P1.Y})
		case 1:
			o = append(o, P{b.P2.X, m.Y})
		case 2:
			o = append(o, P{m.X, b.P2.Y})
		case 3:
			o = append(o, P{b.P1.X, m.Y})
		}
	}
	return o
}

// Triangulations contains all the possible cell triangulation given the cell Number
var Triangulations = []Triangle{
	{-100}, // Outside
	{0, 3},
	{0, 1},
	{1, 3},
	{1, 2},
	{0, 1, 2, 3}, // Saddle
	{0, 2},
	{2, 3},
	{2, 3},
	{0, 2},
	{1, 2, 0, 3}, // Saddle
	{1, 2},
	{1, 3},
	{0, 1},
	{0, 3},
	{100}, // Inside
}
