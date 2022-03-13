package engine

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/matt9mg/doom/images"
	"github.com/matt9mg/doom/music"
	"image/color"
)

type MainMenu struct {
	Background *ebiten.Image
	Skull      *ebiten.Image
	IntroMusic *audio.Player
}

func NewMainMenu() *MainMenu {
	intro, err := mp3.Decode(audioContext, bytes.NewReader(music.Doom_intro))

	if err != nil {
		panic(err)
	}

	player, err := audioContext.NewPlayer(intro)

	if err != nil {
		panic(err)
	}

	player.Play()

	return &MainMenu{
		Background: GetEbitenImage(images.Main_menu),
		Skull:      GetEbitenImage(images.Images_Skull),
		IntroMusic: player,
	}
}

func (m *MainMenu) Render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	screen.DrawImage(m.Background, op)

	texts := []string{"Please Don't Hurt Me!!!", "", "Bring it on!!!", "", "Hardcode", "", "Physco!!!"}
	for i, l := range texts {
		x := screenWidth / 8
		text.Draw(screen, l, arcadeFont, x, (i+3)*fontSize, color.RGBA{
			R: 255,
			G: 50,
			B: 34,
			A: 255,
		})
	}

	skullOp := &ebiten.DrawImageOptions{}
	skullOp.GeoM.Scale(0.1, 0.1)
	skullOp.GeoM.Translate(100, ((2*skullPos)*fontSize)+10)
	screen.DrawImage(m.Skull, skullOp)
}
