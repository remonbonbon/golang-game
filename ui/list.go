package ui

import (
	"image/color"
	"math"
	"myapp/geometry"
	"myapp/render"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type List struct {
	Item            []UIComponent
	Parent          *Window
	BackgroundImage *ebiten.Image

	Scroll       float64
	ScrollMax    float64
	ScrollBarBox geometry.Box

	PrevMouseY int
	Grabbed    bool
}

func (l *List) Update() error {
	_, dy := ebiten.Wheel()
	l.Scroll += -dy * 30

	mx, my := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if l.ScrollBarBox.IsInside(mx, my) {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				l.Grabbed = true
			}
		}
	} else {
		l.Grabbed = false
	}
	if l.Grabbed {
		l.Scroll += float64(my - l.PrevMouseY)
	}
	l.PrevMouseY = my

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

	box := l.Parent.Box().Padding(10)
	render.DrawNineSlice(screen, l.BackgroundImage, box)
	box2 := box.Padding(10)

	offsetY := 0

	clipImage := ebiten.NewImage(box2.Width, box2.Height)

	l.ScrollMax = 0
	// c := color.RGBA{0x00, 0x00, 0xff, 0xff}
	for i, comp := range l.Item {
		bb := comp.Box()
		x, y := 0, offsetY*i-int(l.Scroll)
		offsetY = bb.Height
		l.ScrollMax += float64(offsetY)

		// vector.StrokeRect(clipImage, float32(x), float32(y), float32(bb.Width), float32(bb.Height), 1, c, false)

		comp.Move(x, y).Draw(clipImage)
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

		l.ScrollBarBox = geometry.Box{
			X:      (box2.X + box2.Width - barW),
			Y:      box2.Y + int(l.Scroll/l.ScrollMax*(float64(box2.Height)-barH)),
			Width:  barW,
			Height: int(barH),
		}

		c2 := color.RGBA{0x00, 0x00, 0xff, 0xff}
		mx, my := ebiten.CursorPosition()
		if l.ScrollBarBox.IsInside(mx, my) {
			c2 = color.RGBA{0x00, 0xff, 0xff, 0xff}
		}

		vector.DrawFilledRect(screen,
			float32(l.ScrollBarBox.X), float32(l.ScrollBarBox.Y),
			float32(l.ScrollBarBox.Width), float32(l.ScrollBarBox.Height),
			c2, false)
	}
}
