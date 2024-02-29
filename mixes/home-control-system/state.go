package main

type State interface {
	ChangeHour(int)
	ChangeNumPersons(int)
	Music() Music
	Light() Light
}
