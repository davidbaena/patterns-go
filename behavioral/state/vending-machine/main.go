package main

import "fmt"

func main() {
	vendingMachine := NewVendingMachine(1, 10)
	if err := vendingMachine.RequestItem(); err != nil {
		fmt.Printf("err: %s\n", err)
	}
	if err := vendingMachine.AddItem(2); err != nil {
		fmt.Printf("err: %s\n", err)
	}
	if err := vendingMachine.InsertMoney(10); err != nil {
		fmt.Printf("err: %s\n", err)
	}
	if err := vendingMachine.DispenseItem(); err != nil {
		fmt.Printf("err: %s\n", err)
	}
}
