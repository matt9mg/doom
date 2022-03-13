package engine

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

func GetEbitenImage(imageBytes []byte) *ebiten.Image {
	object, _, err := image.Decode(bytes.NewReader(imageBytes))

	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(object)
}
