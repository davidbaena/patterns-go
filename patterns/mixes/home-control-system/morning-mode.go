package main

type MorningModeState struct {
	*HomeControlSystem
}

func (m *MorningModeState) ChangeHour(hour int) {
	if hour >= 20 && hour < 6 {
		m.SetState(m.NightModeState)
	} else if hour >= 6 && hour < 13 {
		m.SetState(m.MorningModeState)
	} else if hour >= 13 && hour < 20 {
		m.SetState(m.AfternoonModeState)
	}
}

func (m *MorningModeState) ChangeNumPersons(numPersons int) {
	if numPersons > 4 {
		m.SetState(m.PartyModeState)
	} else if numPersons == 0 {
		m.SetState(m.EcoModeState)
	}
}

func (m *MorningModeState) Music() Music {
	return ClassicMusic
}

func (m *MorningModeState) Light() Light {
	return BrightLight
}
