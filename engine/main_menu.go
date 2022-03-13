package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/matt9mg/doom/images"
)

type MainMenu struct {
	background *ebiten.Image
}

func NewMainMenu() *MainMenu  {
	return &MainMenu{}
}

func (m *MainMenu) loadAssets()  {
	m.background = GetEbitenImage(images.Main_menu)
}