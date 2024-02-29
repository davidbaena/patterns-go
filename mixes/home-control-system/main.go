package main

import "time"

func main() {
	system := NewClimateControlSystem()

	clock := Clock{}
	clock.Subscribe(system)

	presenceSensor := PresenceSensor{}
	presenceSensor.Subscribe(system)

	for i := 0; i < 24; i += 2 {
		clock.TimeChange(i)
		time.Sleep(1 * time.Second)
	}

	presenceSensor.DetectPresence(6)
	time.Sleep(1 * time.Second)
	presenceSensor.DetectPresence(0)
	time.Sleep(1 * time.Second)
	presenceSensor.DetectPresence(2)
}
