package service

import (
	"context"

	"newaccess/internal/dto"
	"newaccess/internal/repository"
)

type DeviceService interface {
	Create(ctx context.Context, device *dto.DeviceRequest) (int, error)
	List(ctx context.Context) ([]dto.DeviceResponse, error)
	FindByID(ctx context.Context, id int) (*dto.DeviceResponse, error)
	Update(ctx context.Context, user *dto.DeviceUpdateRequest) error
	Delete(ctx context.Context, id int) error
}

type deviceService struct {
	repo repository.DeviceRepository
}

func NewDeviceService(repo repository.DeviceRepository) DeviceService {
	return &deviceService{repo: repo}
}

func (s *deviceService) Create(ctx context.Context, device *dto.DeviceRequest) (int, error) {
	return s.repo.Create(ctx, device)
}

func (s *deviceService) List(ctx context.Context) ([]dto.DeviceResponse, error) {
	return s.repo.List(ctx)
}

func (s *deviceService) FindByID(ctx context.Context, id int) (*dto.DeviceResponse, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *deviceService) Update(ctx context.Context, device *dto.DeviceUpdateRequest) error {
	return s.repo.Update(ctx, device)
}

func (s *deviceService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
