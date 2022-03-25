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

type Chainsaw struct {
	CurrenFramePosition int
	MousePressed        bool
	Frames              []*Frame
	FramesLength        int
	WeaponSoundFrame    int
	Sprite              *ebiten.Image
	Sound               *audio.Player
	IdleSound           *audio.Player
	Idle                int
}

func NewChainsaw() *Chainsaw {
	img, _, err := image.Decode(bytes.NewReader(images.Image_chainsaw))
	if err != nil {
		log.Fatal(err)
	}

	idle, err := wav.Decode(audioContext, bytes.NewReader(music.Music_chainsaw_idle))

	if err != nil {
		panic(err)
	}

	idleSound, err := audioContext.NewPlayer(idle)
	if err != nil {
		panic(err)
	}

	active, err := wav.Decode(audioContext, bytes.NewReader(music.Music_chainsaw_full))

	if err != nil {
		panic(err)
	}

	Sound, err := audioContext.NewPlayer(active)
	if err != nil {
		panic(err)
	}

	return &Chainsaw{
		Sprite:              ebiten.NewImageFromImage(img),
		CurrenFramePosition: 0,
		MousePressed:        false,
		Frames: []*Frame{
			{
				x0: 0,
				y0: 178,
				x1: 306,
				y1: 0,
			},
			{
				x0: 306,
				y0: 178,
				x1: 612,
				y1: 0,
			},
			{
				x0: 612,
				y0: 178,
				x1: 918,
				y1: 0,
			},
			{
				x0: 900,
				y0: 178,
				x1: 1300,
				y1: 0,
			},
		},
		FramesLength:     5,
		WeaponSoundFrame: 28,
		IdleSound:        idleSound,
		Sound:            Sound,
	}
}

func (w *Chainsaw) Update() {
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

func (w *Chainsaw) PlaySound() {
	w.Sound.Rewind()
	w.Sound.Play()
}

func (w *Chainsaw) RenderCurrentFrame(screen *ebiten.Image) {
	var x int

	minus := 0.0

	if w.MousePressed == false {
		x = 0

		if w.Idle%15 == 0 {
			x = 1
		}

		if w.Sound.IsPlaying() == true {
			w.Sound.Pause()
		}

		if w.IdleSound.IsPlaying() == false {
			w.IdleSound.Rewind()
			w.IdleSound.Play()
		}

		w.Idle++
	}

	if w.MousePressed == true && w.CurrenFramePosition >= 0 && w.CurrenFramePosition < 2 {
		x = 2
		minus = 100.00
	}

	if w.CurrenFramePosition >= 2 {
		x = 3
		minus = 100.00
	}

	if w.MousePressed == true && w.Sound.IsPlaying() == false {
		if w.IdleSound.IsPlaying() == true {
			w.IdleSound.Pause()
		}

		w.Sound.Rewind()
		w.Sound.Play()
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate((screenWidth/2)-150, (screenHeight-305) + minus)

	screen.DrawImage(w.Sprite.SubImage(image.Rect(w.Frames[x].x0, w.Frames[x].y0, w.Frames[x].x1, w.Frames[x].y1)).(*ebiten.Image), op)
}
