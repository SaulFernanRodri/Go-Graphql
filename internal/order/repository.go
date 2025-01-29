package order

import (
	"context"
	"errors"
	"go-graphql/graph/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

type OrderRepositoryInterface interface {
	Create(ctx context.Context, order *model.Order) error
	FindByID(ctx context.Context, id string) (*model.Order, error)
	FindByUser(ctx context.Context, userID string) ([]*model.Order, error)
	All(ctx context.Context) ([]*model.Order, error)
	Delete(ctx context.Context, id string) (*model.Order, error)
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) Create(ctx context.Context, order *model.Order) error {
	return r.DB.Create(order).Error
}

func (r *OrderRepository) FindByID(ctx context.Context, id string) (*model.Order, error) {
	var order model.Order
	if err := r.DB.First(&order, "id = ?", id).Error; err != nil {
		return nil, errors.New("order not found")
	}
	return &order, nil
}

func (r *OrderRepository) FindByUser(ctx context.Context, userID string) ([]*model.Order, error) {
	var orders []*model.Order
	if err := r.DB.Find(&orders, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) All(ctx context.Context) ([]*model.Order, error) {
	var orders []*model.Order
	if err := r.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) Delete(ctx context.Context, id string) (*model.Order, error) {
	order, err := r.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := r.DB.Delete(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
