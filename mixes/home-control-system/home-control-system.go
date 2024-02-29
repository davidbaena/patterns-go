package main

import "log"

type HomeControlSystem struct {
	MorningModeState   State
	AfternoonModeState State
	NightModeState     State
	EcoModeState       State
	PartyModeState     State

	currentState State
	hour         int
	numPersons   int
	temperature  int
}

func NewClimateControlSystem() *HomeControlSystem {
	cs := &HomeControlSystem{
		temperature: 22,
		numPersons:  2,
		hour:        7,
	}
	cs.MorningModeState = &MorningModeState{cs}
	cs.AfternoonModeState = &AfternoonModeState{cs}
	cs.NightModeState = &NightModeState{cs}
	cs.EcoModeState = &EcoModeState{cs}
	cs.PartyModeState = &PartyModeState{cs}

	cs.currentState = cs.MorningModeState

	return cs
}

type Music string

const (
	NoMusic      Music = "none"
	ClassicMusic Music = "classic"
	RelaxMusic   Music = "soft"
	DanceMusic   Music = "dance"
	PartyMusic   Music = "upbeat"
)

type Light string

const (
	NoLight      Light = "none"
	BrightLight  Light = "bright"
	MinimalLight Light = "minimal"
	PartyLight   Light = "colorful"
)

func (c *HomeControlSystem) SetState(s State) {
	c.currentState = s
	c.CurrentStatus()
}

func (c *HomeControlSystem) ChangeHour(hour int) {
	c.hour = hour
	c.currentState.ChangeHour(hour)
}

func (c *HomeControlSystem) ChangeNumPersons(numPersons int) {
	c.numPersons = numPersons
	c.currentState.ChangeNumPersons(numPersons)
}

func (c *HomeControlSystem) Music() Music {
	return c.currentState.Music()
}

func (c *HomeControlSystem) Light() Light {
	return c.currentState.Light()
}

func (c *HomeControlSystem) CurrentStatus() State {
	log.Printf("-----------------------------")
	log.Printf("Current state: %T", c.currentState)
	log.Printf("        hour: %d", c.hour)
	log.Printf("        number of persons: %d", c.numPersons)
	log.Printf("        temperature: %d", c.temperature)
	log.Printf("        music: %s", c.Music())
	log.Printf("        light: %s", c.Light())
	log.Printf("-----------------------------")

	return c.currentState
}

func (c *HomeControlSystem) NotifyHour(hour int) {
	c.ChangeHour(hour)
}

func (c *HomeControlSystem) NotifyPresence(num int) {
	c.ChangeNumPersons(num)
}
