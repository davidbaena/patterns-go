package main

import "fmt"

type GunType string

const (
	GUN    GunType = "GUN"
	AK47   GunType = "AK47"
	MUSKET GunType = "MUSKET"
)

type GunsFactory struct {
}

func (GunsFactory) ProvideGun(gunType GunType) (IGun, error) {

	switch gunType {
	case GUN:
		return NewNormalGun(), nil
	case AK47:
		return NewAk47(), nil
	case MUSKET:
		return NewMusket(), nil
	}

	return nil, fmt.Errorf("ErrUnknownGunType")
}
