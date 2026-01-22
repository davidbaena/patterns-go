package main

type AfternoonModeState struct {
	*HomeControlSystem
}

func (a *AfternoonModeState) ChangeHour(hour int) {
	if hour >= 20 && hour < 6 {
		a.SetState(a.NightModeState)
	} else if hour >= 6 && hour < 13 {
		a.SetState(a.MorningModeState)
	} else if hour >= 13 && hour < 20 {
		a.SetState(a.AfternoonModeState)
	}
}

func (a *AfternoonModeState) ChangeNumPersons(numPersons int) {
	if numPersons > 4 {
		a.SetState(a.PartyModeState)
	} else if numPersons == 0 {
		a.SetState(a.EcoModeState)
	}
}

func (a *AfternoonModeState) Music() Music {
	return DanceMusic
}

func (a *AfternoonModeState) Light() Light {
	return BrightLight
}
