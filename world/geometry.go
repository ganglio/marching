package world

// P is a 2D point
type P struct {
	X float64
	Y float64
}

// BB is a bounding box (a rectangle too...)
type BB struct {
	P1 P
	P2 P
}

// Mid returns the mid point of a BB
func (b BB) Mid() P {
	return P{(b.P1.X + b.P2.X) / 2.0, (b.P1.Y + b.P2.Y) / 2.0}
}
