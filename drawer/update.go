package drawer

import (
	"marching/world"

	"github.com/gonutz/prototype/draw"
)

// Update is used to render a new frame
func Update(w *world.World) func(window draw.Window) {
	return func(window draw.Window) {
		mouseX, mouseY := window.MousePosition()
		window.FillEllipse(mouseX-20, mouseY-20, 40, 40, draw.DarkRed)
	}
}
