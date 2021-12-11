package main

import (
	"math"

	"marching/drawer"
	"marching/world"

	"github.com/gonutz/prototype/draw"
)

func main() {
	w := world.NewWorld(
		func(x, y float64) float64 {
			return world.Entity(func(x, y float64) float64 {
				return 3 * math.Sin(x) * math.Cos(y)
			}).Timed(func(x, y, t float64) (float64, float64) {
				return x + t, y
			})(x, y) * world.Box(0, 0, 2, 2, 10).Timed(func(x, y, t float64) (float64, float64) {
				u := x*math.Cos(t/2) - y*math.Sin(t/2)
				v := x*math.Sin(t/2) + y*math.Cos(t/2)
				return u, v
			})(x, y)
		},
	)

	draw.RunWindow("Title", 1024, 768, drawer.Update(w))
}
