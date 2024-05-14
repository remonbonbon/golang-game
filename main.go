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

type Game struct {
	imageLoader *render.ImageLoader
	window      *ui.Window
	list        *ui.List
}

func (g *Game) Update() error {
	if err := g.list.Update(); err != nil {
		return err
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	{
		tileImage := g.imageLoader.Load("resource/images/bg.png")

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

	g.window.Draw(screen)
	g.list.Draw(screen)

	msg := ""
	msg += fmt.Sprintf("TPS: %0.2f\n", ebiten.ActualTPS())
	msg += fmt.Sprintf("Scroll: %v\n", g.list.Scroll)
	msg += fmt.Sprintf("ScrollMax: %v\n", g.list.ScrollMax)
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 960
}

func main() {
	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("Game")

	game := &Game{}
	game.imageLoader = render.NewImageLoader()
	windowImage := game.imageLoader.Load("resource/images/window.png")
	game.window = &ui.Window{WindowImage: windowImage, Box: geometry.Box{Width: 400, Height: 600, X: 100, Y: 100}}
	game.list = &ui.List{Parent: game.window, BackgroundImage: windowImage}
	game.list.Item = make([]string, 0)
	for i := 1; i < 50; i++ {
		game.list.Item = append(game.list.Item, fmt.Sprintf("アイテム%d", i))
	}
	game.list.Item = append(game.list.Item, "最後のアイテム")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
