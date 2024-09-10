package repository

import "github.com/renanmav/GoExpert-CleanArch/internal/entity"

type OrderRepositoryInterface interface {
	Save(order *entity.Order) error
	FindAll() ([]*entity.Order, error)
	FindById(id string) (*entity.Order, error)
}
