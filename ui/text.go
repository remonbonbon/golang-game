package ui

import (
	"log"
	"myapp/geometry"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/text/language"
)

type Text struct {
	UIComponent
	// fontFaceSource *text.GoTextFaceSource
	face        *text.GoTextFace
	lineSpacing int
	Message     string
	X           int
	Y           int
}

func NewText(Size int) *Text {
	t := &Text{}
	t.face = &text.GoTextFace{
		Source:   fontFaceSource,
		Size:     float64(Size),
		Language: language.Japanese,
	}
	t.lineSpacing = Size
	return t
}

func NewTextMessage(Size int, msg string) *Text {
	t := NewText(Size)
	t.Message = msg
	return t
}

func (t *Text) Draw(screen *ebiten.Image) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(t.X), float64(t.Y))
	op.LineSpacing = float64(t.lineSpacing)
	text.Draw(screen, t.Message, t.face, op)
}

func (t *Text) Box() *geometry.Box {
	w, h := text.Measure(t.Message, t.face, float64(t.lineSpacing))
	return &geometry.Box{X: t.X, Y: t.Y, Width: int(w), Height: int(h)}
}

func (t *Text) Move(x, y int) UIComponent {
	t.X = x
	t.Y = y
	return t
}

////////////////////////////////////////////////

var fontFaceSource *text.GoTextFaceSource

func init() {
	file, err := os.Open("./resource/fonts/rounded-mgenplus-1cp-regular.ttf")
	if err != nil {
		log.Fatal(err)
	}

	s, err := text.NewGoTextFaceSource(file)
	if err != nil {
		log.Fatal(err)
	}
	fontFaceSource = s
}

////////////////////////////////////////////////
