package repository

import (
	"auth-service/internal/models"
	"auth-service/pkg/sqlcqueries"
	"context"

	"github.com/google/uuid"
)

type UserPointsRepository struct {
	q *sqlcqueries.Queries
}

func NewUserPointsRepository(q *sqlcqueries.Queries) *UserPointsRepository {
	return &UserPointsRepository{
		q: q,
	}
}

func (r *UserPointsRepository) Create(ctx context.Context, dto *models.UserPoints) error {
	return r.q.CreateUserPoint(ctx, sqlcqueries.CreateUserPointParams{
		UserID: &dto.User_ID,
		Points: int32(dto.Points),
	})
}

func (r *UserPointsRepository) GetUserPoints(ctx context.Context) ([]*models.GetUserPointsDto, error) {
	pointsRow, err := r.q.GetUserPoints(ctx)
	if err != nil {
		return nil, err
	}

	points := make([]*models.GetUserPointsDto, len(pointsRow))
	for i, value := range pointsRow {
		point := &models.GetUserPointsDto{
			ID:            value.UserPointsID,
			User_ID:       *value.UserID,
			User_FullName: *value.UserFullName,
			Points:        int(value.UserPointsCount),
		}
		points[i] = point
	}

	return points, nil
}

func (r *UserPointsRepository) GetByUserID(ctx context.Context, id uuid.UUID) (*models.GetUserPointsDto, error) {
	pointRow, err := r.q.GetUserPointByUserID(ctx, sqlcqueries.GetUserPointByUserIDParams{UserID: &id})
	if err != nil {
		return nil, err
	}

	return &models.GetUserPointsDto{
		ID:            pointRow.UserPointsID,
		User_ID:       *pointRow.UserID,
		User_FullName: *pointRow.UserFullName,
		Points:        int(pointRow.UserPointsCount),
	}, nil
}

func (r *UserPointsRepository) Update(ctx context.Context, dto *models.UserPoints) error {
	err := r.q.UpdateUserPointsByUserId(ctx, sqlcqueries.UpdateUserPointsByUserIdParams{
		UserID: &dto.User_ID,
		Points: int32(dto.Points),
	})
	return err
}

func (r *UserPointsRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.q.DeleteUserPoints(ctx, sqlcqueries.DeleteUserPointsParams{ID: id})
	return err
}
