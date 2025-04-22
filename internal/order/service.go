package order

import (
	"context"
	"go-graphql/graph/model"
)

type OrderService struct {
	Repo         OrderRepositoryInterface
	OrderCreated chan *model.Order
}

type OrderServiceInterface interface {
	CreateOrder(ctx context.Context, input *model.NewOrderInput) (*model.Order, error)
	GetOrderByID(ctx context.Context, id string) (*model.Order, error)
	GetAllOrders(ctx context.Context) ([]*model.Order, error)
	GetByUser(ctx context.Context, userID string) ([]*model.Order, error)
	DeleteOrder(ctx context.Context, id string) (*model.Order, error)
	getOrderCreated() <-chan *model.Order
}

func NewOrderService(repo *OrderRepository) *OrderService {
	return &OrderService{
		Repo:         repo,
		OrderCreated: make(chan *model.Order, 1),
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, input *model.NewOrderInput) (*model.Order, error) {

	order := &model.Order{
		User:     input.User,
		Product:  input.Product,
		Quantity: input.Quantity,
		Total:    input.Total,
	}

	err := s.Repo.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	go func() {
		s.OrderCreated <- order
	}()

	return order, nil
}

func (s *OrderService) GetOrderByID(ctx context.Context, id string) (*model.Order, error) {
	return s.Repo.FindByID(ctx, id)
}

func (s *OrderService) GetAllOrders(ctx context.Context) ([]*model.Order, error) {
	return s.Repo.All(ctx)
}

func (s *OrderService) GetByUser(ctx context.Context, userID string) ([]*model.Order, error) {
	return s.Repo.FindByUser(ctx, userID)
}

func (s *OrderService) DeleteOrder(ctx context.Context, id string) (*model.Order, error) {
	return s.Repo.Delete(ctx, id)
}

func (s *OrderService) getOrderCreated() <-chan *model.Order {
	return s.OrderCreated
}
