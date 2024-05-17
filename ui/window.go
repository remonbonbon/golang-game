package ui

import (
	"github.com/hajimehoshi/ebiten/v2"

	"myapp/geometry"
	"myapp/render"
)

type Window struct {
	UIComponent
	WindowImage *ebiten.Image
	WindowBox   geometry.Box
}

func (win *Window) Draw(screen *ebiten.Image) {
	render.DrawNineSlice(screen, win.WindowImage, win.WindowBox)
}

func (win *Window) Box() *geometry.Box {
	return &win.WindowBox
}
