package main

type ControlSystem interface {
	Flight() string
}

type BasicControlSystem struct{}

func (bcs BasicControlSystem) Flight() string {
	return "Initiating basic flight"
}

type ObstacleAvoidance struct {
	wrappee ControlSystem
}

func NewObstacleAvoidance(wrappee ControlSystem) ControlSystem {
	return ObstacleAvoidance{wrappee: wrappee}
}

func (oa ObstacleAvoidance) Flight() string {
	return oa.wrappee.Flight() + " with obstacle avoidance"
}

type NightVision struct {
	wrappee ControlSystem
}

func NewNightVision(wrappee ControlSystem) ControlSystem {
	return NightVision{wrappee: wrappee}
}

func (nv NightVision) Flight() string {
	return nv.wrappee.Flight() + " with night vision"
}
