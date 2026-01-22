package main

import (
	"fmt"
	"time"
)

type Quadcopter struct {
	model                 string
	battery               int
	numRotors             int
	gps                   bool
	internalControlSystem ControlSystem
}

func NewQuadcopter() *Quadcopter {
	fmt.Printf("Creating Quadcopter...\n")
	time.Sleep(1 * time.Second)
	return &Quadcopter{
		model:                 "Phantom",
		battery:               5000,
		numRotors:             4,
		gps:                   true,
		internalControlSystem: BasicControlSystem{},
	}
}

func (q Quadcopter) clone() Drone {
	return &Quadcopter{
		model:                 q.model,
		battery:               q.battery,
		numRotors:             q.numRotors,
		gps:                   q.gps,
		internalControlSystem: q.internalControlSystem,
	}
}

func (q Quadcopter) Flight() string {
	return "Run " + q.model
}

func (q Quadcopter) TakeOff() error {
	fmt.Println("Quadcopter taking off...")
	return nil
}
func (q Quadcopter) Land() error {
	fmt.Println("Quadcopter landing...")
	return nil
}
func (q Quadcopter) RoutePlanner() error {
	fmt.Println("Quadcopter route planner...")
	return nil
}

func (q Quadcopter) describe() {
	fmt.Println("-----------------")
	fmt.Printf("Model: %s\n", q.model)
	fmt.Printf("Battery: %d\n", q.battery)
	fmt.Printf("Number of rotors: %d\n", q.numRotors)
	fmt.Printf("GPS: %t\n", q.gps)
	fmt.Println("-----------------")
}
func (q Quadcopter) controlSystem() ControlSystem {
	return q.internalControlSystem
}
func (q *Quadcopter) setControlSystem(system ControlSystem) {
	q.internalControlSystem = system
}
