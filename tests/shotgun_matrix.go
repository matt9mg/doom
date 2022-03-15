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

func (g *Game) Draw(screen *ebiten.Image) {
	if g.count >= 0 && g.count < 12 {
		x0 = 0
		y0 = 206
		x1 = 150
		y1 = 0
	}

	if g.count >= 12 && g.count < 24 {
		x0 = 200
		y0 = 206
		x1 = 402
		y1 = 0
		if pistolSound.IsPlaying() == false {
			pistolSound.Rewind()
			pistolSound.Play()
		}
	}

	if g.count >= 24 && g.count < 36 {
		x0 = 402
		y0 = 206
		x1 = 580
		y1 = 0
	}

	if g.count >= 36 && g.count < 48 {
		x0 = 580
		y0 = 206
		x1 = 795
		y1 = 0
	}

	if g.count >= 48 {
		x0 = 795
		y0 = 206
		x1 = 1209
		y1 = 0
	}

	if g.count > 55 {
		g.count = 0
	}


	log.Println(g.count)
	screen.Clear()
	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate((screenWidth/2)-70, screenHeight-500)

	screen.DrawImage(runnerImage.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) init() {
	pistol, err := wav.Decode(audioContext, bytes.NewReader(music.Music_shotgun))

	if err != nil {
		panic(err)
	}

	pistolSound, err = audioContext.NewPlayer(pistol)
	if err != nil {
		panic(err)
	}
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(images.Images_shotgun))
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
