package engine

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/matt9mg/doom/fonts"
	"github.com/matt9mg/doom/music"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

var (
	GlobalFont      *sfnt.Font
	YourFuckedAudio *audio.Player
	OofSound        *audio.Player
	DoorSound       *audio.Player
	YourAPussySound *audio.Player
	Door2Sound      *audio.Player
	level           = 0
)

func SetGameGlobals() {
	font, err := opentype.Parse(fonts.Font_Doom)

	if err != nil {
		panic(err)
	}

	GlobalFont = font

	fucked, err := wav.Decode(audioContext, bytes.NewReader(music.Music_Menu_Fucked))

	if err != nil {
		panic(err)
	}

	YourFuckedAudio, err = audioContext.NewPlayer(fucked)
	if err != nil {
		panic(err)
	}

	oof, err := wav.Decode(audioContext, bytes.NewReader(music.Music_Oof))

	if err != nil {
		panic(err)
	}

	OofSound, err = audioContext.NewPlayer(oof)
	if err != nil {
		panic(err)
	}

	door, err := wav.Decode(audioContext, bytes.NewReader(music.Music_Switch_On))

	if err != nil {
		panic(err)
	}

	DoorSound, err = audioContext.NewPlayer(door)
	if err != nil {
		panic(err)
	}

	pussy, err := wav.Decode(audioContext, bytes.NewReader(music.Music_Menu_Pussy))

	if err != nil {
		panic(err)
	}

	YourAPussySound, err = audioContext.NewPlayer(pussy)
	if err != nil {
		panic(err)
	}

	door2, err := wav.Decode(audioContext, bytes.NewReader(music.Music_Door2))

	if err != nil {
		panic(err)
	}

	Door2Sound, err = audioContext.NewPlayer(door2)
	if err != nil {
		panic(err)
	}
}
