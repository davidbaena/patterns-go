package domain

type Pizza struct {
	ID      string
	OrderId string
	Name    string
	Status  Status
}

type Status string

const (
	OrderReceived         Status = "OrderReceived"
	CollectingIngredients Status = "CollectingIngredients"
	OvenBaking            Status = "OvenBaking"
	Prepared              Status = "Prepared"
)

func (p *Pizza) nextState() {
	switch p.Status {
	case OrderReceived:
		p.Status = CollectingIngredients
	case CollectingIngredients:
		p.Status = OvenBaking
	case OvenBaking:
		p.Status = Prepared
	}
}
