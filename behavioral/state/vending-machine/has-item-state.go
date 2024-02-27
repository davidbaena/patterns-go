package main

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (h *HasItemState) RequestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.currentState = h.vendingMachine.NoItemState
		return fmt.Errorf("item out of stock")
	}
	fmt.Printf("Item requestd\n")
	h.vendingMachine.currentState = h.vendingMachine.HasMoneyState
	return nil
}

func (h *HasItemState) AddItem(i int) error {
	fmt.Printf("Adding %d items\n", i)
	h.vendingMachine.IncrementItem(i)
	return nil
}

func (h *HasItemState) InsertMoney(i int) error {
	return fmt.Errorf("please select item first")

}

func (h *HasItemState) DispenseItem() error {
	return fmt.Errorf("please select item first")
}
