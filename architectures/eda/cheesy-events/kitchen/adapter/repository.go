package adapter

import (
	"cheesy-events/kitchen/domain"
	"cheesy-events/sqlclient"
)

// Repository implements the PizzaRepository interface using SQLClient
type Repository struct {
	sqlClient sqlclient.SQLClient
}

// NewPizzaRepository creates a new instance of Repository
func NewPizzaRepository(sqlClient sqlclient.SQLClient) Repository {
	return Repository{sqlClient: sqlClient}
}

// CreatePizza inserts a new pizza record into the database
func (r Repository) CreatePizza(pizza domain.Pizza) (int64, error) {
	return r.sqlClient.CreatePizza(sqlclient.Pizza{
		ID:        pizza.ID,
		OrderID:   pizza.OrderId,
		PizzaName: pizza.Name,
		Status:    string(pizza.Status),
	})
}

// GetPizza retrieves a pizza record by ID
func (r Repository) GetPizza(id int) (domain.Pizza, error) {
	pizza, err := r.sqlClient.GetPizza(id)
	if err != nil {
		return domain.Pizza{}, err
	}
	return domain.Pizza{
		ID:      pizza.ID,
		OrderId: pizza.OrderID,
		Name:    pizza.PizzaName,
		Status:  domain.Status(pizza.Status),
	}, nil
}

// UpdatePizza updates an existing pizza record
func (r Repository) UpdatePizza(pizza domain.Pizza) error {
	return r.sqlClient.UpdatePizza(sqlclient.Pizza{
		ID:        pizza.ID,
		OrderID:   pizza.OrderId,
		PizzaName: pizza.Name,
		Status:    string(pizza.Status),
	})
}

// DeletePizza deletes a pizza record by ID
func (r Repository) DeletePizza(id int) error {
	return r.sqlClient.DeletePizza(id)
}
