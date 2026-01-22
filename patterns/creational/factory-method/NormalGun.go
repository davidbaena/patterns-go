package main

type Gun struct {
	name  string
	power int
}

func NewNormalGun() Gun {
	return Gun{
		name:  "NORMAL_001",
		power: 2,
	}
}

func (g Gun) GetName() string {
	return g.name
}
func (g Gun) GetPower() int {
	return g.power
}
