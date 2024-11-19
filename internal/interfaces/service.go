package interfaces

import (
	"auth-service/internal/models"
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	Create(ctx context.Context, user *models.CreateUserDto) error
	GetUsers(ctx context.Context) ([]*models.GetUserDto, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.GetUserDto, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type PointService interface {
	GetAll(ctx context.Context) ([]*models.GetUserPointsDto, error)
	Update(ctx context.Context, dto *models.UserPoints) error
}

type AuthService interface {
	Register(ctx context.Context, dto *models.CreateUserDto) (string, error)
	Login(ctx context.Context, dto *models.LoginDto) (string, error)
}
