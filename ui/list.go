package ui

import (
	"image/color"
	"log"
	"math"
	"myapp/render"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/text/language"
)

type List struct {
	Item            []string
	Parent          *Window
	BackgroundImage *ebiten.Image

	Scroll    float64
	ScrollMax float64
}

var japaneseFaceSource *text.GoTextFaceSource

func init() {
	file, err := os.Open("./resource/fonts/rounded-mgenplus-1cp-regular.ttf")
	if err != nil {
		log.Fatal(err)
	}

	s, err := text.NewGoTextFaceSource(file)
	if err != nil {
		log.Fatal(err)
	}
	japaneseFaceSource = s
}

func (l *List) Update() error {
	_, dy := ebiten.Wheel()
	l.Scroll += -dy * 30

	if l.ScrollMax < l.Scroll {
		l.Scroll = l.ScrollMax
	}
	if l.Scroll < 0 {
		l.Scroll = 0
	}

	return nil
}

func (l *List) Draw(screen *ebiten.Image) {
	if l.Parent == nil {
		return
	}
	if l.BackgroundImage == nil {
		return
	}

	box := l.Parent.Box.Padding(10)
	render.DrawNineSlice(screen, l.BackgroundImage, box)
	box2 := box.Padding(10)

	face := &text.GoTextFace{
		Source:   japaneseFaceSource,
		Size:     16,
		Language: language.Japanese,
	}
	offsetY := 0

	clipImage := ebiten.NewImage(box2.Width, box2.Height)
	// clipImage.Fill(color.Black)

	l.ScrollMax = 0
	color := color.RGBA{0x00, 0x00, 0xff, 0xff}
	for i, v := range l.Item {
		const lineSpacing = 0
		w, h := text.Measure(v, face, lineSpacing)
		x, y := 0, offsetY*i-int(l.Scroll)
		offsetY = int(math.Round(h))
		l.ScrollMax += float64(offsetY)

		vector.StrokeRect(clipImage, float32(x), float32(y), float32(w), float32(h), 1, color, false)

		op := &text.DrawOptions{}
		op.GeoM.Translate(float64(x), float64(y))
		op.LineSpacing = lineSpacing
		text.Draw(clipImage, v, face, op)
	}

	l.ScrollMax -= float64(box2.Height - offsetY)

	{
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(box2.X), float64(box2.Y))
		screen.DrawImage(clipImage, op)
	}

	if 0 < l.ScrollMax {
		barW := 20
		barH := math.Min(100, float64(box2.Height)/3)
		vector.DrawFilledRect(screen, float32(box2.X+box2.Width-barW), float32(float64(box2.Y)+l.Scroll/l.ScrollMax*(float64(box2.Height)-barH)), float32(barW), float32(barH), color, false)
	}
}
