package repository

import (
	"database/sql"
	"time"

	"github.com/renanmav/GoExpert-CleanArch/internal/entity"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepositoryInterface {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)", order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) FindAll() ([]*entity.Order, error) {
	rows, err := r.db.Query("SELECT id, price, tax, final_price, created_at FROM orders ORDER BY created_at ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []*entity.Order{}
	for rows.Next() {
		var id string
		var price, tax, finalPrice float64
		var createdAt time.Time
		err = rows.Scan(&id, &price, &tax, &finalPrice, &createdAt)
		if err != nil {
			return nil, err
		}
		order := entity.Order{
			ID:         id,
			Price:      price,
			Tax:        tax,
			FinalPrice: finalPrice,
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

func (r *OrderRepository) FindById(id string) (*entity.Order, error) {
	row := r.db.QueryRow("SELECT id, price, tax, final_price FROM orders WHERE id = ?", id)
	var price, tax, finalPrice float64
	err := row.Scan(&id, &price, &tax, &finalPrice)
	if err != nil {
		return nil, err
	}
	order := entity.Order{
		ID:         id,
		Price:      price,
		Tax:        tax,
		FinalPrice: finalPrice,
	}
	return &order, nil
}
