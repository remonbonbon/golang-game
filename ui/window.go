package ui

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Window struct {
	WindowImage   *ebiten.Image
	Width, Height int
	X, Y          int
}

func (win *Window) Draw(screen *ebiten.Image) {
	// 9スライス画像を3等分する
	w3 := win.WindowImage.Bounds().Dx()
	h3 := win.WindowImage.Bounds().Dy()
	w1 := int(math.Round(float64(w3) / 3))
	h1 := int(math.Round(float64(h3) / 3))
	w2 := int(math.Round(float64(w3) * 2 / 3))
	h2 := int(math.Round(float64(h3) * 2 / 3))

	cornerLU := win.WindowImage.SubImage(image.Rect(0, 0, w1, h1)).(*ebiten.Image)
	middleL := win.WindowImage.SubImage(image.Rect(0, h1, w1, h2)).(*ebiten.Image)
	cornerLD := win.WindowImage.SubImage(image.Rect(0, h2, w1, h3)).(*ebiten.Image)
	middleU := win.WindowImage.SubImage(image.Rect(w1, 0, w2, h1)).(*ebiten.Image)
	cornerRU := win.WindowImage.SubImage(image.Rect(w2, 0, w3, h1)).(*ebiten.Image)
	middleR := win.WindowImage.SubImage(image.Rect(w2, h1, w3, h2)).(*ebiten.Image)
	cornerRD := win.WindowImage.SubImage(image.Rect(w2, h2, w3, h3)).(*ebiten.Image)
	middleD := win.WindowImage.SubImage(image.Rect(w1, h2, w2, h3)).(*ebiten.Image)
	middleC := win.WindowImage.SubImage(image.Rect(w1, h1, w2, h2)).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	ox := float64(win.X)
	oy := float64(win.Y)

	// 左辺
	op.GeoM.Reset()
	op.GeoM.Scale(1, float64(win.Height-h2)/float64(h1))
	op.GeoM.Translate(ox, oy+float64(h1))
	screen.DrawImage(middleL, op)

	// 上辺
	op.GeoM.Reset()
	op.GeoM.Scale(float64(win.Width-w2)/float64(w1), 1)
	op.GeoM.Translate(ox+float64(w1), oy)
	screen.DrawImage(middleU, op)

	// 右辺
	op.GeoM.Reset()
	op.GeoM.Scale(1, float64(win.Height-h2)/float64(h1))
	op.GeoM.Translate(ox+float64(win.Width-w1), oy+float64(h1))
	screen.DrawImage(middleR, op)

	// 下辺
	op.GeoM.Reset()
	op.GeoM.Scale(float64(win.Width-w2)/float64(w1), 1)
	op.GeoM.Translate(ox+float64(w1), oy+float64(win.Height-h1))
	screen.DrawImage(middleD, op)

	// 左上
	op.GeoM.Reset()
	op.GeoM.Translate(ox, oy)
	screen.DrawImage(cornerLU, op)

	// 左下
	op.GeoM.Reset()
	op.GeoM.Translate(ox, oy+float64(win.Height-h1))
	screen.DrawImage(cornerLD, op)

	// 右上
	op.GeoM.Reset()
	op.GeoM.Translate(ox+float64(win.Width-w1), oy)
	screen.DrawImage(cornerRU, op)

	// 右下
	op.GeoM.Reset()
	op.GeoM.Translate(ox+float64(win.Width-w1), oy+float64(win.Height-h1))
	screen.DrawImage(cornerRD, op)

	// 中央
	op.GeoM.Reset()
	op.GeoM.Scale(float64(win.Width-w2)/float64(w1), float64(win.Height-h2)/float64(h1))
	op.GeoM.Translate(ox+float64(w1), oy+float64(h1))
	screen.DrawImage(middleC, op)
}
