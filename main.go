package main

import (
	"marching/drawer"

	"github.com/gonutz/prototype/draw"
)

func main() {
	draw.RunWindow("Title", 1024, 768, drawer.Update(nil))
}
