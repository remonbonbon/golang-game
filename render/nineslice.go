package render

import (
	"image"
	"math"

	"myapp/geometry"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawNineSlice(screen *ebiten.Image, nineSliceImage *ebiten.Image, Box geometry.Box) {
	// 9スライス画像を3等分する
	w3 := nineSliceImage.Bounds().Dx()
	h3 := nineSliceImage.Bounds().Dy()
	w1 := int(math.Round(float64(w3) / 3))
	h1 := int(math.Round(float64(h3) / 3))
	w2 := int(math.Round(float64(w3) * 2 / 3))
	h2 := int(math.Round(float64(h3) * 2 / 3))

	const (
		None = iota // ストレッチ無し
		Fit         // ストレッチあり
	)
	slices := [][][]int{
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

	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	ox := float64(Box.X)
	oy := float64(Box.Y)
	for _, sliceConf := range slices {
		src := sliceConf[0]
		img := nineSliceImage.SubImage(image.Rect(src[0], src[1], src[2], src[3])).(*ebiten.Image)

		// 角以外の部分は伸縮させる
		op.GeoM.Reset()
		scaleX := float64(1)
		scaleY := float64(1)
		stretch := sliceConf[1]
		// 伸びる部分の倍率 = 左右の角を除いた長さ / 1スライスの長さ
		if stretch[0] == Fit {
			scaleX = float64(Box.Width-w2) / float64(w1)
		}
		if stretch[1] == Fit {
			scaleY = float64(Box.Height-h2) / float64(h1)
		}
		op.GeoM.Scale(scaleX, scaleY)

		dest := sliceConf[2]
		dx := float64(dest[0])
		dy := float64(dest[1])
		// マイナスの場合は逆側の端から
		if dx < 0 {
			dx = float64(Box.Width + dest[0])
		}
		if dy < 0 {
			dy = float64(Box.Height + dest[1])
		}
		op.GeoM.Translate(ox+dx, oy+dy)
		screen.DrawImage(img, op)
	}
}
