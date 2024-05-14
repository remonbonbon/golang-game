package render

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type ImageLoader struct {
	// キー:ファイルパス, 値:画像
	imgMap map[string]*ebiten.Image
}

func NewImageLoader() *ImageLoader {
	il := &ImageLoader{}
	il.imgMap = make(map[string]*ebiten.Image)
	return il
}

// 画像を読み込んで返す。既に読み込んでいる場合はそれを返す
func (il *ImageLoader) Load(name string) *ebiten.Image {
	// 存在する場合はそのまま返す
	{
		img, ok := il.imgMap[name]
		if ok {
			return img
		}
	}

	// 存在しない場合は読み込んで返す
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	decoded, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	img := ebiten.NewImageFromImage(decoded)
	il.imgMap[name] = img

	return img
}
