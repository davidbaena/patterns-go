package main

import "fmt"

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (it *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("item already requested")
}

func (it *ItemRequestedState) AddItem(i int) error {
	return fmt.Errorf("item Dispense in progress")
}

func (it *ItemRequestedState) InsertMoney(money int) error {
	if money < it.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less")
	}
	fmt.Printf("Money inserted: %d\n", money)
	it.vendingMachine.currentState = it.vendingMachine.HasMoneyState
	return nil
}

func (it *ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("please insert money first")
}
