package engine

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/matt9mg/doom/images"
	"github.com/matt9mg/doom/music"
	"image"
	"log"
)

type FuckOff struct {
	CurrenFramePosition int
	MousePressed        bool
	Frames              []*Frame
	FramesLength        int
	WeaponSoundFrame    int
	Sprite              *ebiten.Image
	Sound               *audio.Player
	Played              bool
}

func NewFuckOff() *FuckOff {
	img, _, err := image.Decode(bytes.NewReader(images.Image_middle_finger))
	if err != nil {
		log.Fatal(err)
	}

	finger, err := mp3.Decode(audioContext, bytes.NewReader(music.Music_middle_finger))

	if err != nil {
		panic(err)
	}

	theFinger, err := audioContext.NewPlayer(finger)
	if err != nil {
		panic(err)
	}

	return &FuckOff{
		Sprite:              ebiten.NewImageFromImage(img),
		CurrenFramePosition: 0,
		MousePressed:        false,
		Frames: []*Frame{
			{
				x0: 0,
				y0: 150,
				x1: 154,
				y1: 0,
			},
			{
				x0: 208,
				y0: 150,
				x1: 302,
				y1: 0,
			},
		},
		FramesLength:     5,
		WeaponSoundFrame: 28,
		Sound:            theFinger,
	}
}

func (w *FuckOff) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		w.CurrenFramePosition++
		w.MousePressed = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyF) {
		w.CurrenFramePosition = 0
		w.MousePressed = false
		w.Played = false
	} else if w.MousePressed == true {
		w.CurrenFramePosition++
	} else {
		w.CurrenFramePosition = 0
	}

}

func (w *FuckOff) PlaySound() {
	w.Sound.Rewind()
	w.Sound.Play()
}

func (w *FuckOff) RenderCurrentFrame(screen *ebiten.Image) {
	var x int

	if w.Sound.IsPlaying() == false && w.Played == false {
		w.Played = true
		x = 0
		w.PlaySound()
	}

	if w.CurrenFramePosition > 100 {
		x = 1
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2,2)
	op.GeoM.Translate((screenWidth/2)-160, screenHeight-300)

	screen.DrawImage(w.Sprite.SubImage(image.Rect(w.Frames[x].x0, w.Frames[x].y0, w.Frames[x].x1, w.Frames[x].y1)).(*ebiten.Image), op)
}
