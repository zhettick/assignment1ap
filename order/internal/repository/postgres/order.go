package postgres

import (
	"database/sql"
	"order/internal/domain/entities"

	_ "github.com/lib/pq"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(o *entities.Order) error {
	query := `INSERT INTO orders (id, customer_id, item_name, amount, status, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, o.ID, o.CustomerID, o.ItemName, o.Amount, o.Status, o.CreatedAt)
	return err
}

func (r *OrderRepository) GetByID(id string) (*entities.Order, error) {
	o := &entities.Order{}
	query := `SELECT id, customer_id, item_name, amount, status, created_at FROM orders WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&o.ID, &o.CustomerID, &o.ItemName, &o.Amount, &o.Status, &o.CreatedAt)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (r *OrderRepository) Update(o *entities.Order) error {
	query := `UPDATE orders SET status = $1 WHERE id = $2`
	_, err := r.db.Exec(query, o.Status, o.ID)
	return err
}
