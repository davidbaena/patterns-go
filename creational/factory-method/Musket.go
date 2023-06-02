package main

type Musket struct {
	Gun
}

func NewMusket() Musket {
	return Musket{
		Gun: Gun{
			name:  "MUSKET_001",
			power: 15,
		},
	}
}

func (a Musket) GetName() string {
	return a.name
}
func (a Musket) GetPower() int {
	return a.power
}
