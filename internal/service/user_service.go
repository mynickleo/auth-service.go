package service

import (
	"auth-service/internal/interfaces"
	"auth-service/internal/models"
	"auth-service/internal/utils"
	"context"

	"github.com/google/uuid"
)

type UserService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, user *models.CreateUserDto) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	_, err = s.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUsers(ctx context.Context) ([]*models.GetUserDto, error) {
	return s.repo.GetUsers(ctx)
}

func (s *UserService) GetByID(ctx context.Context, id uuid.UUID) (*models.GetUserDto, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) Update(ctx context.Context, user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return s.repo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
