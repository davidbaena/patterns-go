package main

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (n *NoItemState) RequestItem() error {
	return fmt.Errorf("item out of stock")
}

func (n *NoItemState) AddItem(i int) error {
	n.vendingMachine.IncrementItem(i)
	n.vendingMachine.currentState = &HasItemState{}
	return nil
}

func (n *NoItemState) InsertMoney(i int) error {
	return fmt.Errorf("item out of stock")
}

func (n *NoItemState) DispenseItem() error {
	return fmt.Errorf("item out of stock")
}
