package user

import (
	"context"
	"go-graphql/graph/model"
	"go-graphql/pkg/security"
)

type UserService struct {
	Repo        UserRepositoryInterface
	UserCreated chan *model.User
}

type UserServiceInterface interface {
	CreateUser(ctx context.Context, input *model.NewUserInput) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	UpdateUser(ctx context.Context, id string, name string, email string) (*model.User, error)
	DeleteUser(ctx context.Context, id string) (*model.User, error)
	getUserCretaed() <-chan *model.User
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{
		Repo:        repo,
		UserCreated: make(chan *model.User, 1),
	}
}

func (s *UserService) CreateUser(ctx context.Context, input *model.NewUserInput) (*model.User, error) {
	hashedPassword, err := security.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Activate: input.Activate,
	}

	err = s.Repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	go func() {
		s.UserCreated <- user
	}()

	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	return s.Repo.FindByID(ctx, id)
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.Repo.FindByEmail(ctx, email)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return s.Repo.All(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, id string, name string, email string) (*model.User, error) {
	return s.Repo.Update(ctx, id, name, email)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	return s.Repo.Delete(ctx, id)
}

func (s *UserService) getUserCretaed() <-chan *model.User {
	return s.UserCreated
}
