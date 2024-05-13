package ui

import (
	"github.com/hajimehoshi/ebiten/v2"

	"myapp/geometry"
	"myapp/render"
)

type Window struct {
	WindowImage *ebiten.Image
	Box         geometry.Box
}

func (win *Window) Draw(screen *ebiten.Image) {
	render.DrawNineSlice(screen, win.WindowImage, win.Box)
}
