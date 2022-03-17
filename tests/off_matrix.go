package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/matt9mg/doom/images"
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
var x0, y0, x1, y1 int

func (g *Game) Draw(screen *ebiten.Image) {
	if g.count >= 0 && g.count < 12 {
		x0 = 0
		y0 = 150
		x1 = 154
		y1 = 0
	}

	if g.count >= 12 {
		x0 = 208
		y0 = 150
		x1 = 302
		y1 = 0
	}

	if mousePressed == true {
		if pistolSound.IsPlaying() == false {
			pistolSound.Rewind()
			pistolSound.Play()
		}
	}

	log.Println(g.count)
	screen.Clear()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate((screenWidth/2)-70, screenHeight-175)

	screen.DrawImage(runnerImage.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) init() {
	pistol, err := mp3.Decode(audioContext, bytes.NewReader(music.Music_middle_finger))

	if err != nil {
		log.Fatal(err)
	}

	pistolSound, err = audioContext.NewPlayer(pistol)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(images.Image_middle_finger))
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
