package utils

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const DefaultPathAssetImage = "./public/covers/"

func HandleSingleFile(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("cover")
	title := ctx.Get("title")
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "BAD REQUEST",
		})
	}

	var filename *string
	if file != nil {
		err := checkContentType(file, "image/jpeg", "image/jpg", "image/png")
		if err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"code":   fiber.StatusUnprocessableEntity,
				"status": err.Error(),
			})
		}
		filename = &file.Filename
		ext := filepath.Ext(*filename)
		fmt.Println(file.Header)
		newfilename := fmt.Sprintf("%simages%s", file.Filename, ext)
		fmt.Println(title)

		err = ctx.SaveFile(file, fmt.Sprintf("public/covers/%s", newfilename))
		if err != nil {
			return ctx.JSON(fiber.Map{
				"code":   fiber.StatusBadRequest,
				"status": "BAD REQUEST",
			})
		}
	} else {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "BAD REQUEST",
		})
	}

	if filename != nil {
		ctx.Locals("filename", *filename)
	} else {
		ctx.Locals("filename", nil)
	}

	return ctx.Next()
}

func HandleMultipleFile(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "BAD REQUEST",
		})
	}

	files := form.File["image"]

	var filenames []string
	for i, file := range files {
		var filename string
		if file != nil {
			ext := filepath.Ext(file.Filename)
			filename = fmt.Sprintf("%d-%s%s", i, "images", ext)

			err = ctx.SaveFile(file, fmt.Sprintf("public/covers/%s", filename))
			if err != nil {
				return ctx.JSON(fiber.Map{
					"code":   fiber.StatusBadRequest,
					"status": "BAD REQUEST",
				})
			}
		} else {
			return ctx.JSON(fiber.Map{
				"code":   fiber.StatusBadRequest,
				"status": "BAD REQUEST",
			})
		}

		if filename != "" {
			filenames = append(filenames, filename)
		}
	}

	ctx.Locals("filenames", filenames)

	return ctx.Next()
}

func HandleRemoveFile(filename string, path ...string) error {
	if len(path) > 0 {
		err := os.Remove(path[0] + filename)
		if err != nil {
			log.Println("failed to Remove File")
			return err
		}
	} else {
		err := os.Remove(DefaultPathAssetImage + filename)
		if err != nil {
			log.Println("failed to Remove File")
			return err
		}
	}

	return nil
}

func checkContentType(file *multipart.FileHeader, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		for _, contentType := range contentTypes {
			contentTypeFile := file.Header.Get("Content-Type")
			if contentTypeFile == contentType {
				return nil
			}
		}
		return errors.New("not allowed type of file")
	} else {
		return errors.New("not found content type to be checking")
	}
}
