package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	images "myapp/img"
)

var (
	windowImage *ebiten.Image
	tileImage   *ebiten.Image
)

func init() {
	{
		img, _, err := image.Decode(bytes.NewReader(images.Window))
		if err != nil {
			log.Fatal(err)
		}
		windowImage = ebiten.NewImageFromImage(img)
	}

	{
		img, _, err := image.Decode(bytes.NewReader(images.Background))
		if err != nil {
			log.Fatal(err)
		}
		tileImage = ebiten.NewImageFromImage(img)
	}
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	{
		screenX := screen.Bounds().Bounds().Dx()
		screenY := screen.Bounds().Bounds().Dy()
		w := tileImage.Bounds().Dx()
		h := tileImage.Bounds().Dy()

		op := &ebiten.DrawImageOptions{}
		for y := 0; y < screenY+h; y += h {
			for x := 0; x < screenX+w; x += w {
				screen.DrawImage(tileImage, op)
				op.GeoM.Translate(float64(w), 0)
			}
			op.GeoM.Reset()
			op.GeoM.Translate(0, float64(y))
		}
	}

	{
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(100, 100)
		screen.DrawImage(windowImage, op)
	}

	{
		windowWidth := 200
		windowHeight := 150

		// 9スライス画像を3等分する
		w3 := windowImage.Bounds().Dx()
		h3 := windowImage.Bounds().Dy()
		w1 := int(math.Round(float64(w3) / 3))
		h1 := int(math.Round(float64(h3) / 3))
		w2 := int(math.Round(float64(w3) * 2 / 3))
		h2 := int(math.Round(float64(h3) * 2 / 3))

		cornerLU := windowImage.SubImage(image.Rect(0, 0, w1, h1)).(*ebiten.Image)
		middleL := windowImage.SubImage(image.Rect(0, h1, w1, h2)).(*ebiten.Image)
		cornerLD := windowImage.SubImage(image.Rect(0, h2, w1, h3)).(*ebiten.Image)
		middleU := windowImage.SubImage(image.Rect(w1, 0, w2, h1)).(*ebiten.Image)
		cornerRU := windowImage.SubImage(image.Rect(w2, 0, w3, h1)).(*ebiten.Image)
		middleR := windowImage.SubImage(image.Rect(w2, h1, w3, h2)).(*ebiten.Image)
		cornerRD := windowImage.SubImage(image.Rect(w2, h2, w3, h3)).(*ebiten.Image)
		middleD := windowImage.SubImage(image.Rect(w1, h2, w2, h3)).(*ebiten.Image)
		middleC := windowImage.SubImage(image.Rect(w1, h1, w2, h2)).(*ebiten.Image)

		op := &ebiten.DrawImageOptions{}
		op.Filter = ebiten.FilterLinear
		ox := float64(100)
		oy := float64(400)

		// 左辺
		op.GeoM.Reset()
		op.GeoM.Scale(1, float64(windowHeight-h2)/float64(h1))
		op.GeoM.Translate(ox, oy+float64(h1))
		screen.DrawImage(middleL, op)

		// 上辺
		op.GeoM.Reset()
		op.GeoM.Scale(float64(windowWidth-w2)/float64(w1), 1)
		op.GeoM.Translate(ox+float64(w1), oy)
		screen.DrawImage(middleU, op)

		// 右辺
		op.GeoM.Reset()
		op.GeoM.Scale(1, float64(windowHeight-h2)/float64(h1))
		op.GeoM.Translate(ox+float64(windowWidth-w1), oy+float64(h1))
		screen.DrawImage(middleR, op)

		// 下辺
		op.GeoM.Reset()
		op.GeoM.Scale(float64(windowWidth-w2)/float64(w1), 1)
		op.GeoM.Translate(ox+float64(w1), oy+float64(windowHeight-h1))
		screen.DrawImage(middleD, op)

		// 左上
		op.GeoM.Reset()
		op.GeoM.Translate(ox, oy)
		screen.DrawImage(cornerLU, op)

		// 左下
		op.GeoM.Reset()
		op.GeoM.Translate(ox, oy+float64(windowHeight-h1))
		screen.DrawImage(cornerLD, op)

		// 右上
		op.GeoM.Reset()
		op.GeoM.Translate(ox+float64(windowWidth-w1), oy)
		screen.DrawImage(cornerRU, op)

		// 右下
		op.GeoM.Reset()
		op.GeoM.Translate(ox+float64(windowWidth-w1), oy+float64(windowHeight-h1))
		screen.DrawImage(cornerRD, op)

		// 中央
		op.GeoM.Reset()
		op.GeoM.Scale(float64(windowWidth-w2)/float64(w1), float64(windowHeight-h2)/float64(h1))
		op.GeoM.Translate(ox+float64(w1), oy+float64(h1))
		screen.DrawImage(middleC, op)
	}

	// ebitenutil.DebugPrint(screen, "Hello, World!")
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 960
}

func main() {
	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
