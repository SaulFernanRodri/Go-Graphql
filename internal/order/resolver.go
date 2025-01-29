package order

import (
	"context"
	"go-graphql/graph/model"
)

type OrderResolver struct {
	Service OrderServiceInterface
}

func NewOrderResolver(service OrderServiceInterface) *OrderResolver {
	return &OrderResolver{
		Service: service,
	}
}

func (r *OrderResolver) CreateOrder(ctx context.Context, input *model.NewOrderInput) (*model.Order, error) {
	return r.Service.CreateOrder(ctx, input)
}

func (r *OrderResolver) All(ctx context.Context) ([]*model.Order, error) {
	return r.Service.GetAllOrders(ctx)
}

func (r *OrderResolver) Find(ctx context.Context, id string) (*model.Order, error) {
	return r.Service.GetOrderByID(ctx, id)
}

func (r *OrderResolver) FindByUser(ctx context.Context, userID string) ([]*model.Order, error) {
	return r.Service.GetByUser(ctx, userID)
}


func (r *OrderResolver) DeleteOrder(ctx context.Context, id string) (*model.Order, error) {
	return r.Service.DeleteOrder(ctx, id)
}

func (r *OrderResolver) OrderCreated(ctx context.Context) (<-chan *model.Order, error) {
	return r.Service.getOrderCreated(), nil
}
