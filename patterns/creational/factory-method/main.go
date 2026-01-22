package main

import (
	"fmt"
)

func main() {
	factory := GunsFactory{}

	normalGun, err := factory.ProvideGun(NormalGunType)
	if err != nil {
		fmt.Print(err)
		return
	}

	musket, err := factory.ProvideGun(MusketType)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(normalGun.GetName())
	fmt.Println(musket.GetName())
}
