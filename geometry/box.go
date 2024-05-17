package geometry

type Boxable interface {
	Box() Box
}

// 大きさと位置（左上）
type Box struct {
	Width, Height int // 大きさ
	X, Y          int // 左上の座標
}

// 内側に縮小させたBoxを返す
func (b *Box) Padding(padding int) Box {
	return Box{
		Width:  b.Width - padding*2,
		Height: b.Height - padding*2,
		X:      b.X + padding,
		Y:      b.Y + padding}
}

// 座標x,yがBox上にある場合trueを返す
func (b *Box) IsInside(x, y int) bool {
	return b.X <= x && x <= b.X+b.Width && b.Y <= y && y <= b.Y+b.Height
}
