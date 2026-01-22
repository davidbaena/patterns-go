package main

func main() {
	registry := SetupRegistry()
	//Create a quadcopter
	quadcopter101 := registry.GetByName(quadcopter)
	quadcopter101.describe()

	//Add features to the quadcopter
	avoidance := NewObstacleAvoidance(quadcopter101.controlSystem())
	nightVision := NewNightVision(avoidance)
	quadcopter101.setControlSystem(nightVision)

	println(quadcopter101.controlSystem().Flight())

	remoteControl := &RemoteControl{}
	remoteControl.Submit(&TakeOffCommand{Drone: quadcopter101})
	remoteControl.Submit(&RoutePlannerCommand{Drone: quadcopter101})
	remoteControl.Submit(&LandCommand{Drone: quadcopter101})
	remoteControl.Process()
}
