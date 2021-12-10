package main

import (
	"math"

	"marching/drawer"
	"marching/world"

	"github.com/gonutz/prototype/draw"
)

func main() {
	w := world.NewWorld(
		world.Circle(0, 0, 1),
		world.Box(1, 1, 2, 1, 30).Rot(math.Pi/6).Scale(0.2),
		world.Entity(func(x, y float64) float64 {
			return x + y
		}).Rot(math.Pi/4),
	)

	draw.RunWindow("Title", 1024, 768, drawer.Update(w))
}
