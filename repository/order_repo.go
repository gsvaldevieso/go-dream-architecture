package repository

import "github.com/gsvaldevieso/go-dream-architecture/domain/entity"

// OrderRepo
type OrderRepo struct {
	db Relational
}

// NewOrderRepo
func NewOrderRepo(db Relational) OrderRepo {
	return OrderRepo{db: db}
}

// Store
func (o OrderRepo) Store(order entity.Order) error {
	err := o.db.Exec("INSERT INTO orders (id) VALUES (?)", order.ID)
	if err != nil {
		return err
	}

	return nil
}
