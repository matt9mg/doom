package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Weapons struct {
	Weapons       []WeaponInterface
	CurrentWeapon WeaponInterface
	History       WeaponInterface
}

type WeaponInterface interface {
	Update()
	PlaySound()
	RenderCurrentFrame(screen *ebiten.Image)
}

func LoadWeapons() *Weapons {
	pistol := NewPistol()
	shotgun := NewShotgun()
	chainsaw := NewChainsaw()
	fuckoff := NewFuckOff()

	return &Weapons{
		Weapons: []WeaponInterface{
			pistol,
			shotgun,
			chainsaw,
			fuckoff,
		},
		CurrentWeapon: pistol,
	}
}

func (w *Weapons) ChangeWeapon(index int) {
	w.CurrentWeapon = w.Weapons[(index - 1)]
}

type Frame struct {
	x0 int
	y0 int
	x1 int
	y1 int
}
