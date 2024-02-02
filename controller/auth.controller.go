package controller

import (
	"restful-api-joglo-dev/database"
	"restful-api-joglo-dev/model/entity"
	"restful-api-joglo-dev/model/request"
	"restful-api-joglo-dev/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginUserController(ctx *fiber.Ctx) error {
	loginRequest := new(request.AuthRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "BAD REQUEST",
			"data":   err.Error(),
		})
	}

	validate := validator.New()
	err := validate.Struct(loginRequest)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "BAD REQUEST",
			"data":   err.Error(),
		})
	}

	var user entity.User
	err = database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		ctx.JSON(fiber.Map{
			"code":   fiber.StatusUnauthorized,
			"status": "UNAUTHORIZED",
			"data":   err.Error(),
		})
	}

	// check validasi password
	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "WRONG PASSWORD",
		})
	}

	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	if user.Email == "melani@gmail.com" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, err := utils.GenerateToken(&claims)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusUnauthorized,
			"status": "UNAUHTORIZED",
		})
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"token":  token,
	})
}
