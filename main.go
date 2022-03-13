package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/matt9mg/doom/fonts"
	"github.com/matt9mg/doom/images"
	"github.com/matt9mg/doom/music"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"time"
)

type Mode int

const (
	screenWidth  = 1429
	screenHeight = 893
	sampleRate   = 22050
	fontSize     = 72
)

type Game struct {
	player       *audio.Player
	audioContext *audio.Context
	mode         Mode
	oof          *audio.Player
	door         *audio.Player
	fucked       *audio.Player
	pussy        *audio.Player
}

var (
	backgroundImage *ebiten.Image
	skull           *ebiten.Image

	arcadeFont font.Face
)

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func (g *Game) init() {
	img, _, err := image.Decode(bytes.NewReader(images.Main_menu))
	if err != nil {
		panic(err)
	}

	backgroundImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.Images_Skull))
	if err != nil {
		panic(err)
	}

	skull = ebiten.NewImageFromImage(img)

	tt, err := opentype.Parse(fonts.Font_Doom)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72

	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	if g.audioContext == nil {
		g.audioContext = audio.NewContext(sampleRate)
	}

	intro, err := mp3.Decode(g.audioContext, bytes.NewReader(music.Doom_intro))

	if err != nil {
		panic(err)
	}

	g.player, err = g.audioContext.NewPlayer(intro)
	if err != nil {
		panic(err)
	}

	g.player.Play()

	oof, err := wav.Decode(g.audioContext, bytes.NewReader(music.Music_Oof))

	if err != nil {
		panic(err)
	}

	g.oof, err = g.audioContext.NewPlayer(oof)
	if err != nil {
		panic(err)
	}

	door, err := wav.Decode(g.audioContext, bytes.NewReader(music.Music_Switch_On))

	if err != nil {
		panic(err)
	}

	g.door, err = g.audioContext.NewPlayer(door)
	if err != nil {
		panic(err)
	}

	fucked, err := wav.Decode(g.audioContext, bytes.NewReader(music.Music_Menu_Fucked))

	if err != nil {
		panic(err)
	}

	g.fucked, err = g.audioContext.NewPlayer(fucked)
	if err != nil {
		panic(err)
	}

	pussy, err := wav.Decode(g.audioContext, bytes.NewReader(music.Music_Menu_Pussy))

	if err != nil {
		panic(err)
	}

	g.pussy, err = g.audioContext.NewPlayer(pussy)
	if err != nil {
		panic(err)
	}

	g.player.Play()
}

func (g *Game) isKeyPresses() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		return true
	}

	return false
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) && skullPos > 1 {
		skullPos -= 1
		log.Println("key Pressed up redraw")
		g.oof.Rewind()
		g.oof.Play()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) && skullPos < 4 {
		skullPos += 1
		log.Println("key Pressed down redraw")
		g.oof.Rewind()
		g.oof.Play()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		log.Println("Play game")
		g.door.Play()

		// add a little sleep time between playing sounds
		time.Sleep(time.Second / 2)

		if skullPos == 1 {
			g.pussy.SetVolume(10.0)
			g.pussy.Play()
		}

		if skullPos == 4 {
			g.fucked.SetVolume(10.0)
			g.fucked.Play()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draws Background Image
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	screen.DrawImage(backgroundImage, op)

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

	g.addSkull(screen)
}

var skullPos = 1.0

func (g *Game) addSkull(screen *ebiten.Image) {

	skullOp := &ebiten.DrawImageOptions{}
	skullOp.GeoM.Scale(0.1, 0.1)
	skullOp.GeoM.Translate(100, ((2*skullPos)*fontSize)+10)
	screen.DrawImage(skull, skullOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Doom MT")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
