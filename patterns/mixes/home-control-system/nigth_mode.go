package main

type NightModeState struct {
	*HomeControlSystem
}

func (n *NightModeState) ChangeHour(hour int) {
	// [6, 13] -> MorningModeState
	// [13, 20] -> AfternoonModeState
	// [20, 6] -> NightModeState
	if hour >= 20 && hour < 6 {
		n.SetState(n.NightModeState)
	} else if hour >= 6 && hour < 13 {
		n.SetState(n.MorningModeState)
	} else if hour >= 13 && hour < 20 {
		n.SetState(n.AfternoonModeState)
	}
}

func (n *NightModeState) ChangeNumPersons(numPersons int) {
	if numPersons > 4 {
		n.SetState(n.PartyModeState)
	} else if numPersons == 0 {
		n.SetState(n.EcoModeState)
	}
}
func (n *NightModeState) Music() Music {
	return RelaxMusic
}

func (n *NightModeState) Light() Light {
	return MinimalLight
}
