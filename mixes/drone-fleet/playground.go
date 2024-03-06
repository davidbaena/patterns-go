package main

func main() {
	registry := SetupRegistry()
	//Create a quadcopter
	quadcopter101 := registry.GetByName(quadcopter)
	quadcopter101.describe()

	//Add features to the quadcopter
	avoidance := NewObstacleAvoidance(quadcopter101)
	nightVision := NewNightVision(avoidance)

	flight := nightVision.Flight()
	println(flight)
}
