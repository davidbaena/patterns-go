package main

type Ak47 struct {
	Gun
}

func NewAk47() Ak47 {
	return Ak47{
		Gun: Gun{
			name:  "AK47_001",
			power: 10,
		},
	}
}

func (a Ak47) GetName() string {
	return a.name
}
func (a Ak47) GetPower() int {
	return a.power
}
