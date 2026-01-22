package main

type ObserverClock interface {
	NotifyHour(int)
}

type Clock struct {
	observers []ObserverClock
}

func (c *Clock) Subscribe(o ObserverClock) {
	c.observers = append(c.observers, o)
}

func (c *Clock) TimeChange(hour int) {
	for _, o := range c.observers {
		o.NotifyHour(hour)
	}
}
