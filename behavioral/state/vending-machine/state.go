package main

type State interface {
	RequestItem() error
	AddItem(int) error
	InsertMoney(int) error
	DispenseItem() error
}
