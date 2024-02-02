package middleware

import (
	"restful-api-joglo-dev/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":   fiber.StatusUnauthorized,
			"status": "UNAUTHORIZED",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":   fiber.StatusUnauthorized,
			"status": "UNAUTHORIZED",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		ctx.JSON(fiber.Map{
			"code":   fiber.StatusForbidden,
			"status": "FORBIDDEN",
		})
	}

	ctx.Locals("userInfo", claims)
	// ctx.Locals("role", claims["role"])

	return ctx.Next()
}

func PermissionMiddleware(ctx *fiber.Ctx) error {
	return ctx.Next()
}
