package main

import "fmt"

type Command interface {
	Execute() error
}

type TakeOffCommand struct {
	Drone Drone
}

func (c *TakeOffCommand) Execute() error {
	fmt.Println("Take off command...")
	return c.Drone.TakeOff()
}

type LandCommand struct {
	Drone Drone
}

func (c *LandCommand) Execute() error {
	fmt.Println("Land command...")
	return c.Drone.Land()
}

type RoutePlannerCommand struct {
	Drone Drone
}

func (c *RoutePlannerCommand) Execute() error {
	fmt.Println("Route planner command...")
	return c.Drone.RoutePlanner()
}

type RemoteControl struct {
	Command []Command
}

func (r *RemoteControl) Submit(c Command) {
	r.Command = append(r.Command, c)
}

func (r *RemoteControl) Process() {
	for _, c := range r.Command {
		fmt.Println("Processing command...")
		c.Execute()
	}
	r.Command = nil
}
