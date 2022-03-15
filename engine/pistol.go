package engine

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	images "github.com/matt9mg/doom/images/weapons"
	"github.com/matt9mg/doom/music"
	"image"
	"log"
)

type Pistol struct {
	CurrenFramePosition int
	MousePressed        bool
	Frames              []*Frame
	FramesLength        int
	WeaponSoundFrame    int
	Sprite              *ebiten.Image
	Sound               *audio.Player
}

func NewPistol() *Pistol {
	img, _, err := image.Decode(bytes.NewReader(images.Doom))
	if err != nil {
		log.Fatal(err)
	}

	pistol, err := wav.Decode(audioContext, bytes.NewReader(music.Music_pistol))

	if err != nil {
		panic(err)
	}

	pistolSound, err := audioContext.NewPlayer(pistol)
	if err != nil {
		panic(err)
	}

	return &Pistol{
		Sprite:              ebiten.NewImageFromImage(img),
		CurrenFramePosition: 0,
		MousePressed:        false,
		Frames: []*Frame{
			{
				x0: 0,
				y0: 206,
				x1: 122,
				y1: 0,
			},
			{
				x0: 130,
				y0: 206,
				x1: 255,
				y1: 0,
			},
			{
				x0: 290,
				y0: 206,
				x1: 420,
				y1: 0,
			},
			{
				x0: 440,
				y0: 206,
				x1: 572,
				y1: 0,
			},
			{
				x0: 600,
				y0: 206,
				x1: 725,
				y1: 0,
			},
		},
		FramesLength:     5,
		WeaponSoundFrame: 28,
		Sound:            pistolSound,
	}
}

func (w *Pistol) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		w.CurrenFramePosition++
		w.MousePressed = true
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		w.CurrenFramePosition = 0
		w.MousePressed = false
	} else if w.MousePressed == true {
		w.CurrenFramePosition++
	} else {
		w.CurrenFramePosition = 0
	}
}

func (w *Pistol) PlaySound() {
	w.Sound.Rewind()
	w.Sound.Play()
}

func (w *Pistol) RenderCurrentFrame(screen *ebiten.Image)  {
	var x int
	if w.CurrenFramePosition >= 0 && w.CurrenFramePosition < 8 {
		x = 0
	}

	if w.CurrenFramePosition >= 8 && w.CurrenFramePosition < 16 {
		x = 1
	}

	if w.CurrenFramePosition >= 16 && w.CurrenFramePosition < 24 {
		x = 2
	}

	if w.CurrenFramePosition >= 24 && w.CurrenFramePosition < 32 {
		x = 3
	}

	if w.CurrenFramePosition >= 32 {
		x = 4
	}

	if  w.MousePressed == true && x == 2 && w.Sound.IsPlaying() == false {
		w.PlaySound()
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate((screenWidth/2)-70, screenHeight-140)

	screen.DrawImage(w.Sprite.SubImage(image.Rect(w.Frames[x].x0, w.Frames[x].y0, w.Frames[x].x1, w.Frames[x].y1)).(*ebiten.Image), op)

	if x == 4 {
		w.CurrenFramePosition = 0
	}
}
