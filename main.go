package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"myapp/geometry"
	"myapp/render"
	"myapp/ui"
)

var (
	imageLoader *render.ImageLoader
)

func init() {
	imageLoader = render.NewImageLoader()
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	{
		tileImage := imageLoader.Load("resource/images/bg.png")

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

	windowImage := imageLoader.Load("resource/images/window.png")

	{
		window := ui.Window{WindowImage: windowImage, Box: geometry.Box{Width: 500, Height: 500, X: 50, Y: 50}}
		window.Draw(screen)
	}

	{
		window := ui.Window{WindowImage: windowImage, Box: geometry.Box{Width: 240, Height: 320, X: 100, Y: 300}}
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
