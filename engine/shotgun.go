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

type Shotgun struct {
	CurrenFramePosition int
	MousePressed        bool
	Frames              []*Frame
	FramesLength        int
	WeaponSoundFrame    int
	Sprite              *ebiten.Image
	Sound               *audio.Player
}

func NewShotgun() *Shotgun {
	img, _, err := image.Decode(bytes.NewReader(images.Images_shotgun))
	if err != nil {
		log.Fatal(err)
	}

	shotgun, err := wav.Decode(audioContext, bytes.NewReader(music.Music_shotgun))

	if err != nil {
		panic(err)
	}

	shotgunSounds, err := audioContext.NewPlayer(shotgun)
	if err != nil {
		panic(err)
	}

	return &Shotgun{
		Sprite:              ebiten.NewImageFromImage(img),
		CurrenFramePosition: 0,
		MousePressed:        false,
		Frames: []*Frame{
			{
				x0: 0,
				y0: 206,
				x1: 150,
				y1: 0,
			},
			{
				x0: 200,
				y0: 206,
				x1: 402,
				y1: 0,
			},
			{
				x0: 402,
				y0: 206,
				x1: 570,
				y1: 0,
			},
			{
				x0: 580,
				y0: 206,
				x1: 795,
				y1: 0,
			},
			{
				x0: 795,
				y0: 206,
				x1: 1209,
				y1: 0,
			},
		},
		FramesLength:     5,
		WeaponSoundFrame: 28,
		Sound:            shotgunSounds,
	}
}

func (w *Shotgun) Update() {
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

func (w *Shotgun) PlaySound() {
	w.Sound.Rewind()
	w.Sound.Play()
}

func (w *Shotgun) RenderCurrentFrame(screen *ebiten.Image)  {
	var x int

	minus := 0.0

	if w.CurrenFramePosition >= 0 && w.CurrenFramePosition < 12 {
		x = 0
	}

	if w.CurrenFramePosition >= 12 && w.CurrenFramePosition < 24 {
		x = 1
		minus = 95.0
	}

	if w.CurrenFramePosition >= 24 && w.CurrenFramePosition < 36 {
		x = 2
		minus = 95.0
	}

	if w.CurrenFramePosition >= 36 && w.CurrenFramePosition < 48 {
		x = 3
		minus = 95.0
	}

	if w.CurrenFramePosition >= 48 {
		x = 4
		minus = 95.0
	}

	if w.CurrenFramePosition > 55 {
		w.CurrenFramePosition = 0
	}

	if  w.MousePressed == true && x == 1 && w.Sound.IsPlaying() == false {
		w.PlaySound()
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate((screenWidth/2)-70, (screenHeight - 110) - minus)

	screen.DrawImage(w.Sprite.SubImage(image.Rect(w.Frames[x].x0, w.Frames[x].y0, w.Frames[x].x1, w.Frames[x].y1)).(*ebiten.Image), op)
}
