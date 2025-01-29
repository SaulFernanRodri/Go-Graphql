package user

import (
	"context"
	"go-graphql/graph/model"
)

type UserResolver struct {
	Service UserServiceInterface
}

func NewUserResolver(service UserServiceInterface) *UserResolver {
	return &UserResolver{
		Service: service,
	}
}

func (r *UserResolver) CreateUser(ctx context.Context, input *model.NewUserInput) (*model.User, error) {
	return r.Service.CreateUser(ctx, input)
}

func (r *UserResolver) All(ctx context.Context) ([]*model.User, error) {
	return r.Service.GetAllUsers(ctx)
}

func (r *UserResolver) Find(ctx context.Context, id string) (*model.User, error) {
	return r.Service.GetUserByID(ctx, id)
}

func (r *UserResolver) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	return r.Service.GetUserByEmail(ctx, email)
}

func (r *UserResolver) UpdateUser(ctx context.Context, id string, name string, email string) (*model.User, error) {
	return r.Service.UpdateUser(ctx, id, name, email)
}

func (r *UserResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	return r.Service.DeleteUser(ctx, id)
}

func (r *UserResolver) UserCreated(ctx context.Context) (<-chan *model.User, error) {
	return r.Service.getUserCretaed(), nil
}
