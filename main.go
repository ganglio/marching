package main

import (
	"marching/drawer"
	"marching/world"

	"github.com/gonutz/prototype/draw"
)

func main() {
	w := world.NewWorld(world.Circle(0, 0, 2))

	draw.RunWindow("Title", 1024, 768, drawer.Update(w))
}
