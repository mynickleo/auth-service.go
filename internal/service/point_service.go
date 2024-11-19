package service

import (
	"auth-service/internal/interfaces"
	"auth-service/internal/models"
	"context"
)

type PointService struct {
	repo interfaces.UserPointsRepository
}

func NewPointService(repo interfaces.UserPointsRepository) *PointService {
	return &PointService{
		repo: repo,
	}
}

func (s *PointService) GetAll(ctx context.Context) ([]*models.GetUserPointsDto, error) {
	return s.repo.GetUserPoints(ctx)
}

func (s *PointService) Update(ctx context.Context, dto *models.UserPoints) error {
	return s.repo.Update(ctx, dto)
}
