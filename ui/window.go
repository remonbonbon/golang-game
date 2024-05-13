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

	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	ox := float64(win.X)
	oy := float64(win.Y)

	const (
		None = iota // ストレッチ無し
		Fit         // ストレッチあり
	)
	parts := [][][]int{
		// {
		//   { 9スライス画像の座標x,y,x,y },
		//   {横ストレッチ, 縦ストレッチ},
		//   {描画座標x,y マイナスの場合は逆側の端から}
		// },

		// 左辺
		{{0, h1, w1, h2}, {None, Fit}, {0, h1}},
		// 上辺
		{{w1, 0, w2, h1}, {Fit, None}, {w1, 0}},
		// 右辺
		{{w2, h1, w3, h2}, {None, Fit}, {-w1, h1}},
		// 下辺
		{{w1, h2, w2, h3}, {Fit, None}, {w1, -h1}},
		// 中央
		{{w1, h1, w2, h2}, {Fit, Fit}, {w1, h1}},
		// 左上
		{{0, 0, w1, h1}, {None, None}, {0, 0}},
		// 左下
		{{0, h2, w1, h3}, {None, None}, {0, -h1}},
		// 右上
		{{w2, 0, w3, h1}, {None, None}, {-w1, 0}},
		// 右下
		{{w2, h2, w3, h3}, {None, None}, {-w1, -h1}},
	}
	for _, p := range parts {
		img := win.WindowImage.SubImage(image.Rect(p[0][0], p[0][1], p[0][2], p[0][3])).(*ebiten.Image)
		op.GeoM.Reset()
		scaleX := float64(1)
		scaleY := float64(1)
		if p[1][0] == Fit {
			scaleX = float64(win.Width-w2) / float64(w1)
		}
		if p[1][1] == Fit {
			scaleY = float64(win.Height-h2) / float64(h1)
		}
		op.GeoM.Scale(scaleX, scaleY)

		sx := float64(p[2][0])
		sy := float64(p[2][1])
		if sx < 0 {
			sx = float64(win.Width + p[2][0])
		}
		if sy < 0 {
			sy = float64(win.Height + p[2][1])
		}
		op.GeoM.Translate(ox+sx, oy+sy)
		screen.DrawImage(img, op)
	}
}
