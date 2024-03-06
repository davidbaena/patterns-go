package main

import (
	"fmt"
)

type Drone interface {
	controlSystem() ControlSystem // interface embedding
	setControlSystem(ControlSystem)

	clone() Drone
	describe()

	// Actions
	TakeOff() error
	Land() error
	RoutePlanner() error
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
	droneRegistry.Register(hexacopter, &hexacopterModel)
	droneRegistry.Register(fixedWing, &fixedWingModel)

	quadcopterClone := droneRegistry.CreateDrone(quadcopter)
	quadcopterModelPlus := quadcopterClone.(*Quadcopter)
	quadcopterModelPlus.battery = 6000
	droneRegistry.Register(quadcopterPlus, quadcopterClone)

	hexacopterClone := droneRegistry.CreateDrone(hexacopter)
	hexacopterModelPlus := hexacopterClone.(*Hexacopter)
	hexacopterModelPlus.camera = "8k"
	droneRegistry.Register(hexacopterPlus, hexacopterClone)

	fixedWingClone := droneRegistry.CreateDrone(fixedWing)
	fixedWingModelPlus := fixedWingClone.(*FixedWing)
	fixedWingModelPlus.color = "black"
	droneRegistry.Register(fixedWingPlus, fixedWingClone)

	return droneRegistry
}
