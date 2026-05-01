package service

import (
	"context"

	"github.com/agdaniel10/Go-BasicAPI/internal/model"
	"github.com/agdaniel10/Go-BasicAPI/internal/repository"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAll(ctx context.Context) ([]model.User, error) {
	return s.repo.FindAll(ctx)
}

func (s *UserService) GetByID(ctx context.Context, id string) (*model.User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repo.FindByEmail(ctx, email)
}

func (s *UserService) Create(ctx context.Context, user *model.User) error {
	return s.repo.Create(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error) {
	return s.repo.UpdateUser(ctx, id, user)
}
