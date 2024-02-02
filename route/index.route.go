package route

import (
	"restful-api-joglo-dev/config"
	"restful-api-joglo-dev/controller"
	"restful-api-joglo-dev/middleware"
	"restful-api-joglo-dev/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(route *fiber.App) {
	route.Static("/public", config.ProjectRootPath+"/public/asset")

	route.Post("/api/user/login", controller.LoginUserController)

	route.Get("/api/user", middleware.AuthMiddleware, controller.UserControllerGetAll)
	route.Get("/api/user/:id", controller.UserControllerGetById)
	route.Put("/api/user/:id", controller.UserControllerUpdate)
	route.Put("/api/user/:id/email", controller.UserControllerUpdate)
	route.Post("/api/user", controller.UserControllerCreate)
	route.Delete("/api/user/:id", controller.UserControllerDelete)

	book := route.Group("api/book")
	book.Post("/", utils.HandleSingleFile, controller.BookControllerCreate)
	book.Post("/files", utils.HandleMultipleFile, controller.ImageControllerCreate)
	book.Post("/files/:id", utils.HandleMultipleFile, controller.ImageControllerDelete)
}
