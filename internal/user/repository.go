package user

import (
	"context"
	"errors"
	"go-graphql/graph/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	All(ctx context.Context) ([]*model.User, error)
	Update(ctx context.Context, id string, name string, email string) (*model.User, error)
	Delete(ctx context.Context, id string) (*model.User, error)
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, errors.New("usuario no encontrado")
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, errors.New("usuario no encontrado")
	}
	return &user, nil
}

func (r *UserRepository) All(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, id string, name string, email string) (*model.User, error) {
	user, err := r.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user.Name = name
	user.Email = email
	if err := r.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) (*model.User, error) {
	user, err := r.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := r.DB.Delete(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
