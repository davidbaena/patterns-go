package main

type PartyModeState struct {
	*HomeControlSystem
}

func (p *PartyModeState) ChangeHour(hour int) {
	if hour >= 20 && hour < 6 {
		p.SetState(p.NightModeState)
	} else if hour >= 6 && hour < 13 {
		p.SetState(p.MorningModeState)
	} else if hour >= 13 && hour < 20 {
		p.SetState(p.AfternoonModeState)
	}
}

func (p *PartyModeState) ChangeNumPersons(numPersons int) {
	if numPersons < 4 {
		p.SetState(p.EcoModeState)
	} else if numPersons == 0 {
		p.SetState(p.EcoModeState)
	}
}

func (p *PartyModeState) Music() Music {
	return PartyMusic
}

func (p *PartyModeState) Light() Light {
	return PartyLight
}
