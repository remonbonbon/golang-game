package ui

import (
	"myapp/geometry"
	"myapp/render"
)

type UIComponent interface {
	render.Drawable
	geometry.Boxable
	Move(x, y int) UIComponent
}
