package controller

import (
	"log"
	"restful-api-joglo-dev/database"
	"restful-api-joglo-dev/model/entity"
	"restful-api-joglo-dev/model/request"
	"restful-api-joglo-dev/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ImageControllerCreate(ctx *fiber.Ctx) error {
	image := new(request.ImageCreateRequest)
	if err := ctx.BodyParser(image); err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "BAD REQUEST",
			"data":   err.Error(),
		})
	}

	validate := validator.New()
	err := validate.Struct(image)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "BAD REQUEST",
			"data":   err.Error(),
		})
	}

	// handle file upload

	filenames := ctx.Locals("filenames")
	log.Println("filenames ::", filenames)
	if filenames == nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "BAD REQUEST",
		})
	} else {
		filenamesData := filenames.([]string)
		for _, filename := range filenamesData {
			newImage := entity.Image{
				Image:      filename,
				CategoryID: image.CategoryId,
			}

			err = database.DB.Create(&newImage).Error
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":   fiber.StatusInternalServerError,
					"status": "INTERNAL SERVER ERROR",
					"data":   err.Error(),
				})
			}

		}
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusCreated,
		"status": "OK",
	})
}

func ImageControllerDelete(ctx *fiber.Ctx) error {
	imageId := ctx.Params("id")
	var image entity.Image

	err := database.DB.Debug().First(&image, "id = ?", imageId).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":   fiber.StatusNotFound,
			"status": "NOT FOUND",
		})
	}

	err = utils.HandleRemoveFile(image.Image)
	if err != nil {
		log.Println("fail to delete files")
	}

	err = database.DB.Debug().Delete(&image).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":   fiber.StatusInternalServerError,
			"status": "INTERNAL SERVER ERROR",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
	})
}
