package models

import (
	"github.com/google/uuid"
)

type UserPoints struct {
	ID      uuid.UUID `json:"id"`
	User_ID uuid.UUID `json:"user_id" validate:"required"`
	Points  int       `json:"points" validate:"required"`
}

type GetUserPointsDto struct {
	ID            uuid.UUID `json:"id"`
	User_ID       uuid.UUID `json:"user_id"`
	User_FullName string    `json:"user_full_name"`
	Points        int       `json:"points"`
}
