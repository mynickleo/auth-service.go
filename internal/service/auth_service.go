package service

import (
	"auth-service/internal/interfaces"
	"auth-service/internal/models"
	"auth-service/internal/utils"
	"context"
	"errors"
)

type AuthService struct {
	repoPoints interfaces.UserPointsRepository
	repoUser   interfaces.UserRepository
	jwtModule  interfaces.JWTModule
}

func NewAuthService(repoPoints interfaces.UserPointsRepository, repoUser interfaces.UserRepository, jwtModule interfaces.JWTModule) *AuthService {
	return &AuthService{
		repoPoints: repoPoints,
		repoUser:   repoUser,
		jwtModule:  jwtModule,
	}
}

func (s *AuthService) Register(ctx context.Context, dto *models.CreateUserDto) (string, error) {
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return "", err
	}

	dto.Password = hashedPassword

	id, err := s.repoUser.Create(ctx, dto)
	if err != nil {
		return "", err
	}

	points := models.UserPoints{
		User_ID: *id,
	}
	err = s.repoPoints.Create(ctx, &points)
	if err != nil {
		return "", err
	}

	token, err := s.jwtModule.GenerateToken(id.String())
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Login(ctx context.Context, dto *models.LoginDto) (string, error) {
	user, err := s.repoUser.GetByEmail(ctx, dto.Email)
	if err != nil {
		return "", err
	}

	ok := utils.CheckPasswordHash(dto.Password, user.Password)
	if !ok {
		return "", errors.New("password don't match")
	}

	token, err := s.jwtModule.GenerateToken(user.ID.String())
	if err != nil {
		return "", err
	}

	return token, nil
}
