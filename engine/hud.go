package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/matt9mg/doom/images"
	"image"
)

type Hud struct {
	CurrenFramePosition int
	Hud                 *ebiten.Image
	HudFace             *ebiten.Image
	Frames              []*Frame
	ForceSmileCounter   int
}

func LoadHud() *Hud {
	return &Hud{
		Hud:     GetEbitenImage(images.Images_hud),
		HudFace: GetEbitenImage(images.Images_hud_face),
		Frames: []*Frame{
			{
				x0: 309,
				y0: 123,
				x1: 206,
				y1: 0,
			},
			{
				x0: 206,
				y0: 123,
				x1: 103,
				y1: 0,
			},
			{
				x0: 412,
				y0: 123,
				x1: 309,
				y1: 0,
			},
			{
				x0: 0,
				y0: 123,
				x1: 103,
				y1: 0,
			},
		},
	}
}

func (h *Hud) Update(forceSmile bool) {
	h.CurrenFramePosition++

	if forceSmile == true {
		h.ForceSmileCounter = 1
	}
}

func (h *Hud) RenderCurrentFrame(screen *ebiten.Image) {
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Scale(4.5, 4)

	ops.GeoM.Translate(0, (screenHeight/2)+319)
	screen.DrawImage(h.Hud, ops)

	x := 0

	if h.ForceSmileCounter > 0 && h.ForceSmileCounter < 100 {
		x = 3
		h.ForceSmileCounter++
		h.CurrenFramePosition = 0
	} else {
		h.ForceSmileCounter = 0

		if h.CurrenFramePosition >= 40 && h.CurrenFramePosition < 80 {
			x = 1
		}

		if h.CurrenFramePosition >= 80 && h.CurrenFramePosition < 120 {
			x = 0
		}

		if h.CurrenFramePosition > 120 && h.CurrenFramePosition < 160 {
			x = 2
		}

		if h.CurrenFramePosition > 160 {
			h.CurrenFramePosition = 0
		}
	}

	ops = &ebiten.DrawImageOptions{}
	ops.GeoM.Scale(0.8, 0.8)
	ops.GeoM.Translate((screenWidth/2)-30, (screenHeight/2)+340)
	screen.DrawImage(h.HudFace.SubImage(image.Rect(h.Frames[x].x0, h.Frames[x].y0, h.Frames[x].x1, h.Frames[x].y1)).(*ebiten.Image), ops)
}
