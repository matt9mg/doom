package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Weapons struct {
	Weapons       []WeaponInterface
	CurrentWeapon WeaponInterface
}

type WeaponInterface interface {
	Update()
	PlaySound()
	RenderCurrentFrame(screen *ebiten.Image)
}

func LoadWeapons() *Weapons {
	pistol := NewPistol()
	shotgun := NewShotgun()
	return &Weapons{
		Weapons: []WeaponInterface{
			pistol,
			shotgun,
		},
		CurrentWeapon: pistol,
	}
}

func (w *Weapons) ChangeWeapon(index int)  {
	w.CurrentWeapon = w.Weapons[(index - 1)]
}

type Frame struct {
	x0 int
	y0 int
	x1 int
	y1 int
}
