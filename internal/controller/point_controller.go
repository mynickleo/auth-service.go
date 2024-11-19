package controller

import (
	"auth-service/internal/interfaces"
	"auth-service/internal/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type PointController struct {
	service interfaces.PointService
}

func NewPointController(s interfaces.PointService) *PointController {
	return &PointController{service: s}
}

func (c *PointController) GetAllPoints(ctx *fiber.Ctx) error {
	points, err := c.service.GetAll(context.Background())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(points)
}

func (c *PointController) UpdatePoint(ctx *fiber.Ctx) error {
	userToken := ctx.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	userIDStr, ok := claims["user.id"].(string)
	if !ok || userIDStr == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token payload"})
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid user ID format"})
	}

	point := models.UserPoints{
		Points: 10,
	}

	point.User_ID = userID
	if err := c.service.Update(context.Background(), &point); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON("")
}
