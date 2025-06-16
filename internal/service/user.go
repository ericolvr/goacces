package service

import (
	"context"

	"newaccess/internal/dto"
	"newaccess/internal/repository"
)

type UserService interface {
	FindByID(ctx context.Context, id int) (*dto.UserResponse, error)
	Create(ctx context.Context, user *dto.UserRequest) (int, error)
	Update(ctx context.Context, user *dto.UserUpdateRequest) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]dto.UserResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(ctx context.Context, user *dto.UserRequest) (int, error) {
	return s.repo.Create(ctx, user)
}

func (s *userService) List(ctx context.Context) ([]dto.UserResponse, error) {
	return s.repo.List(ctx)
}

func (s *userService) FindByID(ctx context.Context, id int) (*dto.UserResponse, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *userService) Update(ctx context.Context, user *dto.UserUpdateRequest) error {
	return s.repo.Update(ctx, user)
}

func (s *userService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
