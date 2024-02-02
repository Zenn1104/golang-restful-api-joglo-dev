package controller

import (
	"fmt"
	"restful-api-joglo-dev/database"
	"restful-api-joglo-dev/model/entity"
	"restful-api-joglo-dev/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BookControllerCreate(ctx *fiber.Ctx) error {
	book := new(request.BookCreateRequest)
	if err := ctx.BodyParser(book); err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "BAD REQUEST",
			"data":   err.Error(),
		})
	}

	validate := validator.New()
	err := validate.Struct(book)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "BAD REQUEST",
			"data":   err.Error(),
		})
	}

	// handle file upload
	var filenameString string

	filename := ctx.Locals("filename")
	if filename == nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "BAD REQUEST",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filename)
	}

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filenameString,
	}

	err = database.DB.Create(&newBook).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":   fiber.StatusInternalServerError,
			"status": "INTERNAL SERVER ERROR",
			"data":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusCreated,
		"status": "OK",
		"data":   newBook,
	})
}
