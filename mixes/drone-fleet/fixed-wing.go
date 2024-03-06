package main

import (
	"fmt"
	"time"
)

func NewFixedWing() FixedWing {
	fmt.Printf("Creating FixedWing...\n")
	time.Sleep(1 * time.Second)

	return FixedWing{
		model:                 "SkyHawk",
		battery:               10000,
		wheels:                true,
		color:                 "white",
		internalControlSystem: BasicControlSystem{},
	}
}

type FixedWing struct {
	model                 string
	battery               int
	wheels                bool
	color                 string
	internalControlSystem ControlSystem
}

func (f *FixedWing) setControlSystem(system ControlSystem) {
	f.internalControlSystem = system
}

func (f FixedWing) controlSystem() ControlSystem {
	return f.internalControlSystem
}

func (f FixedWing) Land() error {
	fmt.Println("FixedWing landing...")
	return nil
}

func (f FixedWing) describe() {
	fmt.Println("-----------------")
	fmt.Printf("Model: %s\n", f.model)
	fmt.Printf("Battery: %d\n", f.battery)
	fmt.Printf("Wheels: %t\n", f.wheels)
	fmt.Printf("Color: %s\n", f.color)
	fmt.Println("-----------------")

}

func (f FixedWing) clone() Drone {
	return &FixedWing{
		model:                 f.model,
		battery:               f.battery,
		wheels:                f.wheels,
		color:                 f.color,
		internalControlSystem: f.internalControlSystem,
	}
}

func (f FixedWing) TakeOff() error {
	fmt.Println("FixedWing taking off...")
	return nil
}

func (h FixedWing) Flight() string {
	return "Run " + h.model
}

func (f FixedWing) RoutePlanner() error {
	fmt.Println("FixedWing route planner...")
	return nil
}
