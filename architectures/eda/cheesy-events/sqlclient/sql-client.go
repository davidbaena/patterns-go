package sqlclient

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Pizza struct {
	ID      int    `json:"id"`
	OrderID string `json:"order_id"`
	Pizza   string `json:"pizza"`
	Status  string `json:"status"`
}

type SQLClient struct {
	db *sql.DB
}

func NewSqClient() *SQLClient {
	dataSourceName := "postgres://user:password@localhost:5432/pizzadb?sslmode=disable"
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	createTable := `CREATE TABLE IF NOT EXISTS pizzas (
		id SERIAL PRIMARY KEY,
		order_id TEXT,
		pizza TEXT,
		status TEXT
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
	return &SQLClient{db: db}
}

// CreatePizza inserts a new pizza record into the database
func (s *SQLClient) CreatePizza(pizza Pizza) (int64, error) {
	result, err := s.db.Exec("INSERT INTO pizzas (order_id, pizza, status) VALUES ($1, $2, $3)", pizza.OrderID, pizza.Pizza, pizza.Status)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// GetPizza retrieves a pizza record by ID
func (s *SQLClient) GetPizza(id int) (*Pizza, error) {
	row := s.db.QueryRow("SELECT id, order_id, pizza, status FROM pizzas WHERE id = $1", id)
	var pizza Pizza
	err := row.Scan(&pizza.ID, &pizza.OrderID, &pizza.Pizza, &pizza.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &pizza, nil
}

// UpdatePizza updates an existing pizza record
func (s *SQLClient) UpdatePizza(pizza Pizza) error {
	_, err := s.db.Exec("UPDATE pizzas SET order_id = $1, pizza = $2, status = $3 WHERE id = $4", pizza.OrderID, pizza.Pizza, pizza.Status, pizza.ID)
	return err
}

// DeletePizza deletes a pizza record by ID
func (s *SQLClient) DeletePizza(id int) error {
	_, err := s.db.Exec("DELETE FROM pizzas WHERE id = $1", id)
	return err
}
