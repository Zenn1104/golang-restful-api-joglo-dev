package controller

import (
	"log"
	"restful-api-joglo-dev/database"
	"restful-api-joglo-dev/model/entity"
	"restful-api-joglo-dev/model/request"
	"restful-api-joglo-dev/model/response"
	"restful-api-joglo-dev/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserControllerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	err := database.DB.Find(&users)
	if err != nil {
		log.Println(&err)
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   users,
	})
}

func UserControllerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "BAD REQUEST",
			"data":   err.Error(),
		})
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "BAD REQUEST",
			"data":   err.Error(),
		})
	}

	newUser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Address:  user.Address,
		Phone:    user.Phone,
	}

	password, err := utils.HashingPassword(user.Password)
	if err != nil {
		ctx.JSON(fiber.Map{
			"code":   fiber.StatusInternalServerError,
			"status": "INTERNAL SERVER ERROR",
			"data":   err.Error(),
		})
	}

	newUser.Password = password

	err = database.DB.Create(&newUser).Error
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusInternalServerError,
			"status": "INTERNAL SERVER ERROR",
			"data":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusCreated,
		"status": "OK",
		"data":   newUser,
	})
}

func UserControllerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		ctx.JSON(fiber.Map{
			"code":   fiber.StatusNotFound,
			"status": "NOT FOUND",
			"data":   err.Error(),
		})
	}

	response := response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Address:   user.Address,
		Phone:     user.Phone,
		CratedAt:  user.CratedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   response,
	})
}

func UserControllerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "BAD REQUEST",
			"data":   err.Error(),
		})
	}

	userId := ctx.Params("id")
	var user entity.User

	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusNotFound,
			"status": "NOT FOUND",
			"data":   err.Error(),
		})
	}

	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	user.Address = userRequest.Address
	user.Phone = userRequest.Phone
	err = database.DB.Save(&user).Error
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusInternalServerError,
			"status": "INTERNAL SERVER ERROR",
			"data":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   user,
	})
}

func UserControllerUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(request.UserEmailRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "BAD REQUEST",
			"data":   err.Error(),
		})
	}

	userId := ctx.Params("id")
	var user entity.User
	var isEmail entity.User

	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusNotFound,
			"status": "NOT FOUND",
			"data":   err.Error(),
		})
	}

	err = database.DB.First(&isEmail, "email = ?", userRequest.Email).Error
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusNotFound,
			"status": "NOT FOUND",
			"data":   err.Error(),
		})
	}

	user.Email = userRequest.Email
	err = database.DB.Save(&user).Error
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusInternalServerError,
			"status": "INTERNAL SERVER ERROR",
			"data":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   user,
	})
}

func UserControllerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user entity.User

	err := database.DB.Debug().First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusNotFound,
			"status": "NOT FOUND",
			"data":   err.Error(),
		})
	}

	err = database.DB.Debug().Delete(&user).Error
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusInternalServerError,
			"status": "INTERNAL SERVER ERROR",
			"data":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
	})
}
