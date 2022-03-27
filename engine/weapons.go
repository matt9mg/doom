package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Weapons struct {
	Weapons       []WeaponInterface
	CurrentWeapon WeaponInterface
	History       WeaponInterface
	TheBird       WeaponInterface
}

type WeaponInterface interface {
	Update()
	PlaySound()
	RenderCurrentFrame(screen *ebiten.Image)
}

func LoadWeapons() *Weapons {
	pistol := NewPistol()

	fuckoff := NewFuckOff()

	return &Weapons{
		Weapons: []WeaponInterface{
			pistol,
		},
		CurrentWeapon: pistol,
		TheBird: fuckoff,
	}
}

func (w *Weapons) AddChainsaw() {
	chaisaw := NewChainsaw()
	w.Weapons = append(w.Weapons, chaisaw)
	w.CurrentWeapon = chaisaw
}

func (w *Weapons) AddShotGun()  {
	shotgun := NewShotgun()
	w.Weapons = append(w.Weapons, shotgun)
	w.CurrentWeapon = shotgun
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


