package main

import (
	"fmt"
)

func main() {
	factory := GunsFactory{}

	gun, err := factory.ProvideGun(GUN)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(gun.GetName())
}
