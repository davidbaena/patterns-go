package main

import "fmt"

type Observer interface {
	Notify(string)
}
type Fan struct {
	Name string
}

func (f Fan) Notify(event string) {
	fmt.Printf("Fan %s received match result: %s\n", f.Name, event)
}

type FootballMatch struct {
	observers []Observer
}

func (f *FootballMatch) Subscribe(o Observer) {
	f.observers = append(f.observers, o)
}

func (f *FootballMatch) GoalScored(team string) {
	event := fmt.Sprintf("Goal scored by %s!", team)
	for _, o := range f.observers {
		o.Notify(event)
	}
}

func main() {
	match := FootballMatch{}
	fan1 := Fan{Name: "Alice"}
	fan2 := Fan{Name: "Bob"}

	match.Subscribe(fan1)
	match.Subscribe(fan2)

	match.GoalScored("Team 1")
	match.GoalScored("Team 2")
	match.GoalScored("Team 1")
}
