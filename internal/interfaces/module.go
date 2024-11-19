package interfaces

import (
	"github.com/gofiber/fiber/v2"
)

type JWTModule interface {
	JWTGuard() fiber.Handler
	CheckUserGuard(c *fiber.Ctx) error
	GenerateToken(userId string) (string, error)
}

type UserModule interface {
	Initialization() error
	GetRepo() UserRepository
}

type PointModule interface {
	Initialization() error
	GetRepo() UserPointsRepository
}

type ReadyModule interface {
	Initialization() error
}

type AuthModule interface {
	Initialization() error
}
