package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	images "myapp/img"
	"myapp/ui"
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
		window := ui.Window{}
		window.WindowImage = windowImage
		window.Height = 500
		window.Width = 500
		window.X = 50
		window.Y = 50
		window.Draw(screen)
	}

	{
		window := ui.Window{}
		window.WindowImage = windowImage
		window.Height = 320
		window.Width = 240
		window.X = 100
		window.Y = 300
		window.Draw(screen)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 960
}

func main() {
	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
