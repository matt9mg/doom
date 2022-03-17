package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/matt9mg/doom/images/weapons"
	"github.com/matt9mg/doom/music"
	"image"
	_ "image/png"
	"log"
)

const (
	screenWidth  = 500
	screenHeight = 500
	sampleRate   = 22050
)

var (
	runnerImage  *ebiten.Image
	audioContext = audio.NewContext(sampleRate)
	pistolSound  *audio.Player
	activeSound  *audio.Player
)

type Game struct {
	count int
}

var mousePressed = false

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.count++
		mousePressed = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		g.count = 0
		mousePressed = false
	} else if mousePressed == true {
		g.count++
	} else {
		g.count = 0
	}

	return nil
}

var x0 = 0
var y0 = 206
var x1 = 122
var y1 = 0
var idle int
var minus float64

func (g *Game) Draw(screen *ebiten.Image) {
	if g.count >= 0 && g.count < 12 {
		minus = 0.00
		x0 = 0
		y0 = 178
		x1 = 306
		y1 = 0

		//idle
		if pistolSound.IsPlaying() == false {
			if activeSound.IsPlaying() == true {
				activeSound.Pause()
			}
			pistolSound.Rewind()
			pistolSound.Play()
		}

		if idle%15 == 0 {
			x0 = 306
			y0 = 178
			x1 = 612
			y1 = 0
		}

		idle++
	}

	if g.count >= 12 && g.count < 14 {
		minus = 100.00
		x0 = 612
		y0 = 178
		x1 = 918
		y1 = 0
	}

	if g.count >= 14 {
		minus = 100.00
		x0 = 900
		y0 = 178
		x1 = 1300
		y1 = 0
	}

	if mousePressed == true {
		if pistolSound.IsPlaying() == true {
			pistolSound.Pause()
		}

		if activeSound.IsPlaying() == false {
			activeSound.Rewind()
			activeSound.Play()
		}
	}

	log.Println(g.count)
	screen.Clear()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate((screenWidth/2)-70, (screenHeight-175) + minus)

	screen.DrawImage(runnerImage.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) init() {
	pistol, err := wav.Decode(audioContext, bytes.NewReader(music.Music_chainsaw_idle))

	if err != nil {
		log.Fatal(err)
	}

	pistolSound, err = audioContext.NewPlayer(pistol)
	if err != nil {
		log.Fatal(err)
	}

	sound, err := wav.Decode(audioContext, bytes.NewReader(music.Music_chainsaw_full))

	if err != nil {
		log.Fatal(err)
	}

	activeSound, err = audioContext.NewPlayer(sound)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(images.Image_chainsaw))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Test Playground for Features")

	game := &Game{}
	game.init()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
