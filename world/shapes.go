package world

import (
	"math"
)

// Circle - ditto
func Circle(cx, cy, r float64) Entity {
	return func(x, y float64) float64 {
		return r * r / ((x-cx)*(x-cx) + (y-cy)*(y-cy))
	}
}

// Box - ditto
func Box(cx, cy, w, h, exp float64) Entity {
	return func(x, y float64) float64 {
		return 1 / (math.Pow((x-cx)/w, exp) + math.Pow((y-cy)/h, exp))
	}
}
