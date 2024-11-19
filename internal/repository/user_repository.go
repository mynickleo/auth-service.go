package repository

import (
	"auth-service/internal/models"
	"auth-service/pkg/sqlcqueries"
	"context"

	"github.com/google/uuid"
)

type UserRepository struct {
	q *sqlcqueries.Queries
}

func NewUserRepository(q *sqlcqueries.Queries) *UserRepository {
	return &UserRepository{
		q: q,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *models.CreateUserDto) (*uuid.UUID, error) {
	user.ID = uuid.New()

	err := r.q.CreateUser(ctx, sqlcqueries.CreateUserParams{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		FullName: &user.FullName,
	})

	if err != nil {
		return nil, err
	}

	return &user.ID, nil
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]*models.GetUserDto, error) {
	userRow, err := r.q.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]*models.GetUserDto, len(userRow))
	for i, value := range userRow {
		user := &models.GetUserDto{
			ID:        value.ID,
			Email:     value.Email,
			FullName:  *value.FullName,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		users[i] = user
	}

	return users, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.GetUserDto, error) {
	userRow, err := r.q.GetUserByID(ctx, sqlcqueries.GetUserByIDParams{ID: id})
	if err != nil {
		return nil, err
	}

	return &models.GetUserDto{
		ID:        userRow.ID,
		Email:     userRow.Email,
		FullName:  *userRow.FullName,
		CreatedAt: userRow.CreatedAt,
		UpdatedAt: userRow.UpdatedAt,
	}, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	userRow, err := r.q.GetUserByEmail(ctx, sqlcqueries.GetUserByEmailParams{Email: email})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:        userRow.ID,
		Email:     userRow.Email,
		Password:  userRow.Password,
		FullName:  *userRow.FullName,
		CreatedAt: userRow.CreatedAt,
		UpdatedAt: userRow.UpdatedAt,
	}, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	err := r.q.UpdateUser(ctx, sqlcqueries.UpdateUserParams{
		FullName: &user.FullName,
		Email:    user.Email,
		Password: user.Password,
		ID:       user.ID,
	})
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.q.DeleteUser(ctx, sqlcqueries.DeleteUserParams{ID: id})
	return err
}
