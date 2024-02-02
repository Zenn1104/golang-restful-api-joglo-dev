package main

import (
	"restful-api-joglo-dev/database"
	"restful-api-joglo-dev/database/migrations"
	"restful-api-joglo-dev/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()
	migrations.InitMigrations()
	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":8080")
}
