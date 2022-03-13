package engine

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/matt9mg/doom/images"
	"github.com/matt9mg/doom/music"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"time"
)

type MainMenu struct {
	Background *ebiten.Image
	Skull      *ebiten.Image
	IntroMusic *audio.Player
	Font       font.Face
	SkullPos   float64
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

	font, err := opentype.NewFace(GlobalFont, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}

	return &MainMenu{
		Background: GetEbitenImage(images.Main_menu),
		Skull:      GetEbitenImage(images.Images_Skull),
		IntroMusic: player,
		Font:       font,
		SkullPos:   1,
	}
}

func (m *MainMenu) Render(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	screen.DrawImage(m.Background, op)

	texts := []string{"Please Don't Hurt Me!!!", "", "Bring it on!!!", "", "Hardcode", "", "Physco!!!"}
	for i, l := range texts {
		x := screenWidth / 8
		text.Draw(screen, l, m.Font, x, (i+3)*fontSize, color.RGBA{
			R: 255,
			G: 50,
			B: 34,
			A: 255,
		})
	}

	skullOp := &ebiten.DrawImageOptions{}
	skullOp.GeoM.Scale(0.1, 0.1)
	skullOp.GeoM.Translate(100, ((2*m.SkullPos)*fontSize)+10)
	screen.DrawImage(m.Skull, skullOp)
}

func (m *MainMenu) HandleKeyPress()  {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) && m.SkullPos > 1 {
		m.SkullPos -= 1
		log.Println("key Pressed up redraw")
		OofSound.Rewind()
		OofSound.Play()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) && m.SkullPos < 4 {
		m.SkullPos += 1
		log.Println("key Pressed down redraw")
		OofSound.Rewind()
		OofSound.Play()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		log.Println("Play game")
		DoorSound.Rewind()
		DoorSound.Play()

		// add a little sleep time between playing sounds
		time.Sleep(time.Second / 2)

		if m.SkullPos == 1 {
			YourAPussySound.SetVolume(5.0)
			YourAPussySound.Play()
			time.Sleep(time.Second * 4)
		}

		if m.SkullPos == 4 {
			YourFuckedAudio.SetVolume(5.0)
			YourFuckedAudio.Play()
			time.Sleep(time.Second * 2)
		}

		level = 1
	}
}