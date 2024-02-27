package main

// VendingMachine is the context
type VendingMachine struct {
	NoItemState        State
	HasItemState       State
	ItemRequestedState State
	HasMoneyState      State

	currentState State
	itemCount    int
	itemPrice    int
}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *VendingMachine) AddItem(i int) error {
	return v.currentState.AddItem(i)
}

func (v *VendingMachine) InsertMoney(i int) error {
	return v.currentState.InsertMoney(i)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func (v *VendingMachine) SetState(s State) {
	v.currentState = s
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {

	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	v.NoItemState = &NoItemState{v}
	v.HasItemState = &HasItemState{v}
	v.ItemRequestedState = &ItemRequestedState{v}
	v.HasMoneyState = &HasMoneyState{v}

	v.currentState = v.HasItemState

	return v
}

func (v *VendingMachine) IncrementItem(i int) {
	v.itemCount += i
}
