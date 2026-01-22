package main

type EcoModeState struct {
	*HomeControlSystem
}

func (e *EcoModeState) ChangeHour(hour int) {
	if hour > 6 {
		e.SetState(e.MorningModeState)
	} else if hour > 22 {
		e.SetState(e.AfternoonModeState)
	}
}

func (e *EcoModeState) ChangeNumPersons(numPersons int) {
	if numPersons > 4 {
		e.SetState(e.PartyModeState)
	}
}

func (e *EcoModeState) Music() Music {
	return NoMusic
}

func (e *EcoModeState) Light() Light {
	return NoLight
}
