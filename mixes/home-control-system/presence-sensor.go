package main

type ObserverSensor interface {
	NotifyPresence(num int)
}

type PresenceSensor struct {
	observers []ObserverSensor
}

func (p *PresenceSensor) Subscribe(o ObserverSensor) {
	p.observers = append(p.observers, o)
}

func (p *PresenceSensor) DetectPresence(num int) {
	for _, o := range p.observers {
		o.NotifyPresence(num)
	}
}
