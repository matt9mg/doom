package engine

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/matt9mg/doom/images"
	"github.com/matt9mg/doom/music"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"math"
	"math/rand"
	"path/filepath"
	"runtime"
)

type Mode int

const (
	screenWidth  = 1429
	screenHeight = 893
	sampleRate   = 22050
	fontSize     = 72
	screenScale  = 1.0
	texSize      = 256
	dpi          = 72
)

var (
	audioContext = audio.NewContext(sampleRate)
)

type Game struct {
	player       *audio.Player
	audioContext *audio.Context
	mode         Mode

	leve1Map1 *audio.Player

	floor *ebiten.Image

	//--viewport and width / height--//
	view   *ebiten.Image
	width  int
	height int

	//--define camera--//
	camera *Camera

	mouseMode      MouseMode
	mouseX, mouseY int

	//--test texture--//

	//--array of levels, levels refer to "floors" of the world--//
	mapObj     *Map
	levels     []*Level
	spriteLvls []*Level
	floorLvl   *HorLevel

	tex *TextureHandler

	slices []*image.Rectangle

	spriteBatch *SpriteBatch

	level1 *audio.Player

	menu *MainMenu

	weapons *Weapons
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func (g *Game) init() {
	SetGameGlobals()
	g.menu = NewMainMenu()

	// use scale to keep the desired window width and height
	g.width = int(math.Floor(float64(screenWidth) / screenScale))
	g.height = int(math.Floor(float64(screenHeight) / screenScale))

	g.tex = NewTextureHandler(texSize)

	//--init texture slices--//
	g.slices = g.tex.GetSlices()

	// load map
	g.mapObj = NewMap(g.tex)

	//--inits the levels--//
	g.levels, g.floorLvl = g.createLevels(4)

	// load content once when first run
	g.loadContent()

	// init the sprites
	g.mapObj.LoadSprites()
	g.spriteLvls = g.createSpriteLevels()

	// give sprite a sample velocity for movement
	s := g.mapObj.GetSprite(0)
	s.Vx = -0.02
	// give sprite custom bounds for collision instead of using image bounds
	s.W = int(s.Scale * 85)
	s.H = int(s.Scale * 126)

	// init mouse movement mode
	ebiten.SetCursorMode(ebiten.CursorModeCaptured)
	g.mouseMode = MouseModeMove
	g.mouseX, g.mouseY = math.MinInt32, math.MinInt32

	//--init camera--//
	g.camera = NewCamera(g.width, g.height, texSize, g.mapObj, g.slices, g.levels, g.floorLvl, g.spriteLvls, g.tex)

	level1map1, err := mp3.Decode(audioContext, bytes.NewReader(music.Music_Level1_Map1))

	if err != nil {
		panic(err)
	}

	g.level1, err = audioContext.NewPlayer(level1map1)
	if err != nil {
		panic(err)
	}

	g.weapons = LoadWeapons()
}

func (g *Game) Update() error {
	if level == 1 {
		err := g.menu.IntroMusic.Close()
		g.level1.Play()

		if err != nil {
			panic(err)
		}
		g.updateSprites()

		// handle input
		g.handleInput()

		if inpututil.IsKeyJustPressed(ebiten.Key1) == true {
			g.weapons.ChangeWeapon(1)
		}

		if inpututil.IsKeyJustPressed(ebiten.Key2) == true {
			g.weapons.ChangeWeapon(2)
		}

		if inpututil.IsKeyJustPressed(ebiten.Key3) == true {
			g.weapons.ChangeWeapon(3)
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyF) == true {
			g.weapons.History = g.weapons.CurrentWeapon
			g.weapons.ChangeWeapon(4)
		}

		g.weapons.CurrentWeapon.Update()

		if inpututil.IsKeyJustReleased(ebiten.KeyF) == true {
			g.weapons.CurrentWeapon = g.weapons.History
			g.weapons.History = nil
		}
	}

	if level == 0 {
		g.menu.HandleKeyPress()
		//level = 1
	}

	return nil
}

var test = false

func (g *Game) Draw(screen *ebiten.Image) {
	g.view = screen
	g.view.Clear()

	if level == 1 {
		g.camera.Update()

		//--draw basic sky and floor--//
		texRect := image.Rect(0, 0, texSize, texSize)
		whiteRGBA := &color.RGBA{255, 255, 255, 255}

		floorRect := image.Rect(0, int(float64(g.height)*0.5)+g.camera.GetPitch(),
			g.width, 2*int(float64(g.height)*0.5)-g.camera.GetPitch())
		g.spriteBatch.draw(g.floor, &floorRect, &texRect, whiteRGBA)

		//--draw walls--//
		for x := 0; x < g.width; x++ {
			for i := cap(g.levels) - 1; i >= 0; i-- {
				g.spriteBatch.draw(g.levels[i].CurrTex[x], g.levels[i].Sv[x], g.levels[i].Cts[x], g.levels[i].St[x])
			}
		}

		// draw textured floor
		floorImg := ebiten.NewImageFromImage(g.floorLvl.HorBuffer)
		if floorImg == nil {
			log.Fatal("floorImg is nil")
		} else {
			op := &ebiten.DrawImageOptions{}
			op.Filter = ebiten.FilterLinear
			g.view.DrawImage(floorImg, op)
		}

		// draw sprites
		for x := 0; x < g.width; x++ {
			for i := 0; i < cap(g.spriteLvls); i++ {
				spriteLvl := g.spriteLvls[i]
				if spriteLvl == nil {
					continue
				}

				texture := spriteLvl.CurrTex[x]
				if texture != nil {
					g.spriteBatch.draw(texture, spriteLvl.Sv[x], spriteLvl.Cts[x], spriteLvl.St[x])
				}
			}
		}

		g.weapons.CurrentWeapon.RenderCurrentFrame(screen)
	} else {
		// render the menu
		g.menu.Render(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func Play() {
	// need to set some runtime vars
	numCPU := runtime.NumCPU()
	maxProcs := runtime.GOMAXPROCS(numCPU)
	runtime.GOMAXPROCS(maxProcs)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Doom MT")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

func (s *SpriteBatch) draw(texture *ebiten.Image, destinationRectangle *image.Rectangle, sourceRectangle *image.Rectangle, color *color.RGBA) {
	if texture == nil || destinationRectangle == nil || sourceRectangle == nil {
		return
	}

	if sourceRectangle.Min.X == 0 {
		// fixes subImage from clipping at edges of textures which can cause gaps
		sourceRectangle.Min.X++
		sourceRectangle.Max.X++
	}

	// if destinationRectangle is not the same size as sourceRectangle, scale to fit
	var scaleX, scaleY float64 = 1.0, 1.0
	if !destinationRectangle.Eq(*sourceRectangle) {
		sSize := sourceRectangle.Size()
		dSize := destinationRectangle.Size()

		scaleX = float64(dSize.X) / float64(sSize.X)
		scaleY = float64(dSize.Y) / float64(sSize.Y)
	}

	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear

	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(float64(destinationRectangle.Min.X), float64(destinationRectangle.Min.Y))

	destTexture := texture.SubImage(*sourceRectangle).(*ebiten.Image)

	if color != nil {
		// color channel modulation/tinting
		op.ColorM.Scale(float64(color.R)/255, float64(color.G)/255, float64(color.B)/255, float64(color.A)/255)
	}

	view := s.g.view
	view.DrawImage(destTexture, op)
}

func (g *Game) createSpriteLevels() []*Level {
	// create empty "level" for all sprites to render using similar slice methods as walls
	numSprites := g.mapObj.GetNumSprites()

	spriteArr := make([]*Level, numSprites)

	return spriteArr
}

type SpriteBatch struct {
	g *Game
}

// loadContent will be called once per game and is the place to load
// all of your content.
func (g *Game) loadContent() {
	// Create a new SpriteBatch, which can be used to draw textures.
	g.spriteBatch = &SpriteBatch{g: g}

	// TODO: use loadContent to load your game content here
	g.tex.Textures = make([]*ebiten.Image, 16)

	g.spriteBatch = &SpriteBatch{g: g}

	// TODO: use loadContent to load your game content here
	g.tex.Textures = make([]*ebiten.Image, 16)

	g.tex.Textures[0] = getTextureFromFile("wall.png")
	g.tex.Textures[1] = getTextureFromFile("left_bot_house.png")
	g.tex.Textures[2] = getTextureFromFile("right_bot_house.png")
	g.tex.Textures[3] = getTextureFromFile("left_top_house.png")
	g.tex.Textures[4] = getTextureFromFile("right_top_house.png")
	g.tex.Textures[5] = getTextureFromFile("door.png")

	// separating sprites out a bit from wall textures
	g.tex.Textures[9] = getSpriteFromFile("tree_09.png")
	g.tex.Textures[10] = getSpriteFromFile("tree_10.png")
	g.tex.Textures[14] = getSpriteFromFile("tree_14.png")

	g.tex.Textures[15] = getSpriteFromFile("mt.png")

	// just setting the grass texture apart from the rest since it gets special handling
	g.floorLvl.TexRGBA = make([]*image.NRGBA, 1)
	g.floorLvl.TexRGBA[0] = getNRGBAFromFile("floor2.png")

	floor, _, err := image.Decode(bytes.NewReader(images.Images_Floor))
	if err != nil {
		panic(err)
	}

	g.floor = ebiten.NewImageFromImage(floor)
}

func getNRGBAFromFile(texFile string) *image.NRGBA {
	var rgba *image.NRGBA
	resourcePath := filepath.Join("engine", "content", "textures")
	_, tex, err := ebitenutil.NewImageFromFile(filepath.Join(resourcePath, texFile))
	if err != nil {
		log.Fatal(err)
	}
	if tex != nil {
		rgba = image.NewNRGBA(image.Rect(0, 0, texSize, texSize))
		// convert into NRGBA format
		for x := 0; x < texSize; x++ {
			for y := 0; y < texSize; y++ {
				clr := tex.At(x, y).(color.NRGBA)
				rgba.SetNRGBA(x, y, clr)
			}
		}
	}

	return rgba
}

func getTextureFromFile(texFile string) *ebiten.Image {
	resourcePath := filepath.Join("engine", "content", "textures")
	eImg, _, err := ebitenutil.NewImageFromFile(filepath.Join(resourcePath, texFile))
	if err != nil {
		log.Fatal(err)
	}
	return eImg
}

func getSpriteFromFile(sFile string) *ebiten.Image {
	resourcePath := filepath.Join("engine", "content", "sprites")
	eImg, _, err := ebitenutil.NewImageFromFile(filepath.Join(resourcePath, sFile))
	if err != nil {
		log.Fatal(err)
	}
	return eImg
}

//returns initialised Level structs
func (g *Game) createLevels(numLevels int) ([]*Level, *HorLevel) {
	levelArr := make([]*Level, numLevels)

	for i := 0; i < numLevels; i++ {
		levelArr[i] = new(Level)
		levelArr[i].Sv = SliceView(g.width, g.height)
		levelArr[i].Cts = make([]*image.Rectangle, g.width)
		levelArr[i].St = make([]*color.RGBA, g.width)
		levelArr[i].CurrTex = make([]*ebiten.Image, g.width)
	}

	horizontalLevel := new(HorLevel)
	horizontalLevel.Clear(g.width, g.height)

	return levelArr, horizontalLevel
}

func (g *Game) handleInput() {
	forward := false
	backward := false
	rotLeft := false
	rotRight := false

	moveModifier := 1.0
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		moveModifier = 2.0
	}

	switch {
	case ebiten.IsKeyPressed(ebiten.KeyControl):
		if g.mouseMode != MouseModeCursor {
			ebiten.SetCursorMode(ebiten.CursorModeVisible)
			g.mouseMode = MouseModeCursor
		}

	case ebiten.IsKeyPressed(ebiten.KeyAlt):
		if g.mouseMode != MouseModeMove {
			ebiten.SetCursorMode(ebiten.CursorModeCaptured)
			g.mouseMode = MouseModeMove
			g.mouseX, g.mouseY = math.MinInt32, math.MinInt32
		}

	case g.mouseMode != MouseModeLook:
		ebiten.SetCursorMode(ebiten.CursorModeCaptured)
		g.mouseMode = MouseModeLook
		g.mouseX, g.mouseY = math.MinInt32, math.MinInt32
	}

	switch g.mouseMode {
	case MouseModeCursor:
		g.mouseX, g.mouseY = ebiten.CursorPosition()
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			fmt.Printf("mouse left clicked: (%v, %v)\n", g.mouseX, g.mouseY)

			// using left click for debugging graphical issues
			/*if g.DebugX == -1 && g.DebugY == -1 {
				// only allow setting once between clears to debounce
				g.DebugX = g.mouseX
				g.DebugY = g.mouseY
				g.DebugOnce = true
			}*/
		}

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
			fmt.Printf("mouse right clicked: (%v, %v)\n", g.mouseX, g.mouseY)

			// using right click to clear the debugging flag
			/*g.DebugX = -1
			g.DebugY = -1
			g.DebugOnce = false*/
		}

	case MouseModeMove:
		x, y := ebiten.CursorPosition()
		switch {
		case g.mouseX == math.MinInt32 && g.mouseY == math.MinInt32:
			// initialize first position to establish delta
			if x != 0 && y != 0 {
				g.mouseX, g.mouseY = x, y
			}

		default:
			dx, dy := g.mouseX-x, g.mouseY-y
			g.mouseX, g.mouseY = x, y

			if dx != 0 {
				g.camera.Rotate(0.005 * float64(dx) * moveModifier)
			}

			if dy != 0 {
				g.camera.Move(0.01 * float64(dy) * moveModifier)
			}
		}
	case MouseModeLook:
		x, y := ebiten.CursorPosition()
		switch {
		case g.mouseX == math.MinInt32 && g.mouseY == math.MinInt32:
			// initialize first position to establish delta
			if x != 0 && y != 0 {
				g.mouseX, g.mouseY = x, y
			}

		default:
			dx, dy := g.mouseX-x, g.mouseY-y
			g.mouseX, g.mouseY = x, y

			if dx != 0 {
				g.camera.Rotate(0.005 * float64(dx) * moveModifier)
			}

			if dy != 0 {
				g.camera.Pitch(dy)
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		rotLeft = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		rotRight = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		forward = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		backward = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyC) {
		g.camera.Crouch()
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.camera.Jump()
	} else {
		g.camera.Stand()
	}

	if forward {
		g.camera.Move(0.06 * moveModifier)
	} else if backward {
		g.camera.Move(-0.06 * moveModifier)
	}

	if g.mouseMode == MouseModeLook || g.mouseMode == MouseModeMove {
		// strafe instead of rotate
		if rotLeft {
			g.camera.Strafe(-0.05 * moveModifier)
		} else if rotRight {
			g.camera.Strafe(0.05 * moveModifier)
		}
	} else {
		if rotLeft {
			g.camera.Rotate(0.03 * moveModifier)
		} else if rotRight {
			g.camera.Rotate(-0.03 * moveModifier)
		}
	}
}

func (g *Game) updateSprites() {
	// Testing animated sprite movement
	sprites := g.mapObj.GetSprites()

	for _, s := range sprites {
		if s.Vx != 0 || s.Vy != 0 {
			// TODO: use ebiten.CurrentTPS() to determine actual velicity amount to move sprite per tick

			horBounds := float64(s.W/2) / float64(texSize)

			xCheck := int(s.X)
			yCheck := int(s.Y)
			if s.Vx > 0 {
				xCheck = int(s.X + s.Vx + horBounds)
			} else if s.Vx < 0 {
				xCheck = int(s.X + s.Vx - horBounds)
			}

			if s.Vy > 0 {
				yCheck = int(s.Y + s.Vy + horBounds)
			} else if s.Vy < 0 {
				yCheck = int(s.Y + s.Vy - horBounds)
			}

			if g.mapObj.GetAt(xCheck, yCheck) == 0 {
				// simple collision check to prevent phasing through walls
				s.X += s.Vx
				s.Y += s.Vy
			} else {
				// for testing purposes, letting the sample sprite ping pong off walls in somewhat random direction
				s.Vx = randFloat(-0.03, 0.03)
				s.Vy = randFloat(-0.03, 0.03)
			}
		}
		s.Update()
	}
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
