package drawer

import (
	"fmt"
	"marching/world"

	"github.com/gonutz/prototype/draw"
)

var (
	camera    = world.BB{world.P{-4, -3}, world.P{4, 3}}
	threshold = 1.0
)

// Update is used to render a new frame
func Update(w *world.World) func(window draw.Window) {
	return func(window draw.Window) {

		window.DrawScaledText(fmt.Sprintf("%3.2f", threshold), 10, 10, 0.75, draw.Green)

		if window.MouseWheelY() > 0 {
			camera = camera.Scale(1.01)
		}
		if window.MouseWheelY() < 0 {
			camera = camera.Scale(0.99)
		}
		if window.IsKeyDown(draw.KeyRight) {
			threshold += 0.1
		}
		if window.IsKeyDown(draw.KeyLeft) {
			threshold -= 0.1
		}
		if window.IsKeyDown(draw.KeyR) {
			camera = world.BB{world.P{-4, -3}, world.P{4, 3}}
			threshold = 0
		}
		if window.IsKeyDown(draw.KeyQ) {
			window.Close()
		}

		g := w.Geometries(camera, 100, 100, threshold)

		w, h := window.Size()
		screen := world.BB{world.P{0, 0}, world.P{float64(w), float64(h)}}

		for _, e := range g {
			p := e.Project(camera, screen)
			for _, edge := range p.E {
				p0 := p.V[edge.S]
				p1 := p.V[edge.E]
				window.DrawLine(int(p0.X), int(p0.Y), int(p1.X), int(p1.Y), draw.Yellow)
			}
		}

		// Something to keep stuff moving
		mouseX, mouseY := window.MousePosition()
		window.FillEllipse(mouseX-10, mouseY-10, 20, 20, draw.DarkRed)

	}
}
