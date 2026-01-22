package main

import (
	"fmt"
	"time"
)

type Hexacopter struct {
	model                 string
	battery               int
	numRotors             int
	camera                string
	gps                   bool
	internalControlSystem ControlSystem
}

func NewHexacopter() Hexacopter {
	fmt.Printf("Creating Hexacopter...\n")

	time.Sleep(1 * time.Second)
	return Hexacopter{
		model:                 "HexaX",
		battery:               8000,
		numRotors:             6,
		camera:                "4k",
		gps:                   true,
		internalControlSystem: BasicControlSystem{},
	}
}

func (h Hexacopter) describe() {
	fmt.Println("-----------------")
	fmt.Printf("Model: %s\n", h.model)
	fmt.Printf("Battery: %d\n", h.battery)
	fmt.Printf("Number of rotors: %d\n", h.numRotors)
	fmt.Printf("Camera: %s\n", h.camera)
	fmt.Printf("GPS: %t\n", h.gps)
	fmt.Println("-----------------")
}

func (h Hexacopter) clone() Drone {
	return &Hexacopter{
		model:                 h.model,
		battery:               h.battery,
		numRotors:             h.numRotors,
		camera:                h.camera,
		gps:                   h.gps,
		internalControlSystem: h.internalControlSystem,
	}
}

func (h Hexacopter) Flight() string {
	return "Run " + h.model
}

func (h *Hexacopter) TakeOff() error {
	fmt.Println("Hexacopter taking off...")
	//Consumes battery
	h.battery -= 100
	return nil
}
func (h Hexacopter) RoutePlanner() error {
	h.controlSystem().Flight() // depends on the control system
	return nil
}

func (h Hexacopter) Land() error {
	fmt.Println("Hexacopter landing...")
	h.battery = 1000
	return nil
}

func (h Hexacopter) controlSystem() ControlSystem {
	return h.internalControlSystem
}
func (h *Hexacopter) setControlSystem(system ControlSystem) {
	h.internalControlSystem = system
}
