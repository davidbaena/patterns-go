package main

import (
	"fmt"
	"time"
)

type Drone interface {
	ControlSystem
	clone() Drone
	describe()
}

type Quadcopter struct {
	model     string
	battery   int
	numRotors int
	gps       bool
}

func NewQuadcopter() *Quadcopter {
	fmt.Printf("Creating Quadcopter...\n")
	time.Sleep(1 * time.Second)
	return &Quadcopter{
		model:     "Phantom",
		battery:   5000,
		numRotors: 4,
		gps:       true,
	}
}

func (q Quadcopter) clone() Drone {
	return Quadcopter{
		model:     q.model,
		battery:   q.battery,
		numRotors: q.numRotors,
		gps:       q.gps,
	}
}

func (q Quadcopter) Flight() string {
	return "Run " + q.model
}

func (q Quadcopter) describe() {
	fmt.Println("-----------------")
	fmt.Printf("Model: %s\n", q.model)
	fmt.Printf("Battery: %d\n", q.battery)
	fmt.Printf("Number of rotors: %d\n", q.numRotors)
	fmt.Printf("GPS: %t\n", q.gps)
	fmt.Println("-----------------")

}

type Hexacopter struct {
	model     string
	battery   int
	numRotors int
	camera    string
	gps       bool
}

func NewHexacopter() Hexacopter {
	fmt.Printf("Creating Hexacopter...\n")

	time.Sleep(1 * time.Second)
	return Hexacopter{
		model:     "HexaX",
		battery:   8000,
		numRotors: 6,
		camera:    "4k",
		gps:       true,
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
	return Hexacopter{
		model:     h.model,
		battery:   h.battery,
		numRotors: h.numRotors,
		camera:    h.camera,
		gps:       h.gps,
	}
}

func (h Hexacopter) Flight() string {
	return "Run " + h.model
}

func NewFixedWing() FixedWing {
	fmt.Printf("Creating FixedWing...\n")
	time.Sleep(1 * time.Second)

	return FixedWing{
		model:   "SkyHawk",
		battery: 10000,
		wheels:  true,
		color:   "white",
	}
}

type FixedWing struct {
	model   string
	battery int
	wheels  bool
	color   string
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
	return FixedWing{
		model:   f.model,
		battery: f.battery,
		wheels:  f.wheels,
		color:   f.color,
	}
}

func (h FixedWing) Flight() string {
	return "Run " + h.model
}

type DroneType = string

const (
	quadcopter     DroneType = "quadcopter"
	quadcopterPlus DroneType = "quadcopterPlus"
	hexacopter     DroneType = "hexacopter"
	hexacopterPlus DroneType = "hexacopterPlus"
	fixedWing      DroneType = "fixedWing"
	fixedWingPlus  DroneType = "fixedWingPlus"
)

type DroneRegistry struct {
	drones map[DroneType]Drone
}

func NewDroneRegistry() DroneRegistry {
	return DroneRegistry{
		drones: make(map[string]Drone),
	}
}

func (dr *DroneRegistry) Register(name string, drone Drone) {
	dr.drones[name] = drone
}

func (dr *DroneRegistry) GetByName(name string) Drone {
	drone, ok := dr.drones[name]
	if !ok {
		fmt.Println("Drone not found")
		return nil
	}
	return drone

}

func (dr *DroneRegistry) CreateDrone(name string) Drone {
	drone, ok := dr.drones[name]
	if !ok {
		fmt.Println("Drone not found")
		return nil
	}
	return drone.clone()
}

func SetupRegistry() DroneRegistry {

	quadcopterModel := NewQuadcopter()
	hexacopterModel := NewHexacopter()
	fixedWingModel := NewFixedWing()

	droneRegistry := NewDroneRegistry()

	droneRegistry.Register(quadcopter, quadcopterModel)
	droneRegistry.Register(hexacopter, hexacopterModel)
	droneRegistry.Register(fixedWing, fixedWingModel)

	quadcopterClone := droneRegistry.CreateDrone(quadcopter)
	quadcopterModelPlus := quadcopterClone.(Quadcopter)
	quadcopterModelPlus.battery = 6000
	droneRegistry.Register(quadcopterPlus, quadcopterClone)

	hexacopterClone := droneRegistry.CreateDrone(hexacopter)
	hexacopterModelPlus := hexacopterClone.(Hexacopter)
	hexacopterModelPlus.camera = "8k"
	droneRegistry.Register(hexacopterPlus, hexacopterClone)

	fixedWingClone := droneRegistry.CreateDrone(fixedWing)
	fixedWingModelPlus := fixedWingClone.(FixedWing)
	fixedWingModelPlus.color = "black"
	droneRegistry.Register(fixedWingPlus, fixedWingClone)

	return droneRegistry

}
