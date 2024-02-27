package main

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (h *HasMoneyState) RequestItem() error {
	return fmt.Errorf("item dispense in progress")

}

func (h *HasMoneyState) AddItem(i int) error {
	return fmt.Errorf("item dispense in progress")
}

func (h *HasMoneyState) InsertMoney(i int) error {
	return fmt.Errorf("item dispense in progress")
}

func (h *HasMoneyState) DispenseItem() error {
	fmt.Printf("Dispensing item\n")
	h.vendingMachine.itemCount--
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.currentState = h.vendingMachine.NoItemState
	} else {
		h.vendingMachine.currentState = h.vendingMachine.HasItemState
	}
	return nil
}
