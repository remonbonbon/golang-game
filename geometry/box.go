package geometry

// 大きさと位置（左上）
type Box struct {
	Width, Height int // 大きさ
	X, Y          int // 左上の座標
}

func (b *Box) Padding(padding int) Box {
	return Box{
		Width:  b.Width - padding*2,
		Height: b.Height - padding*2,
		X:      b.X + padding,
		Y:      b.Y + padding}
}
