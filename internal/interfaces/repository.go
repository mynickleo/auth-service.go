package interfaces

import (
	"auth-service/internal/models"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, dto *models.CreateUserDto) (*uuid.UUID, error)
	GetUsers(ctx context.Context) ([]*models.GetUserDto, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.GetUserDto, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, dto *models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserPointsRepository interface {
	Create(ctx context.Context, dto *models.UserPoints) error
	GetUserPoints(ctx context.Context) ([]*models.GetUserPointsDto, error)
	GetByUserID(ctx context.Context, id uuid.UUID) (*models.GetUserPointsDto, error)
	Update(ctx context.Context, dto *models.UserPoints) error
	Delete(ctx context.Context, id uuid.UUID) error
}
