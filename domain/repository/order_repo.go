package repository

import "github.com/gsvaldevieso/go-dream-architecture/domain/entity"

type OrderRepository interface {
	Store(entity.Order) error
}
