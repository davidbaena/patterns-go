package main

import "fmt"

type GunType string

const (
	NormalGunType GunType = "NormalGun"
	Ak47Type      GunType = "AK47"
	MusketType    GunType = "MUSKET"
)

type GunsFactory struct {
}

func (GunsFactory) ProvideGun(gunType GunType) (IGun, error) {

	switch gunType {
	case NormalGunType:
		return NewNormalGun(), nil
	case Ak47Type:
		return NewAk47(), nil
	case MusketType:
		return NewMusket(), nil
	}

	return nil, fmt.Errorf("ErrUnknownGunType")
}
