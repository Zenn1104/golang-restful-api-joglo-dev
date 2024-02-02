package migrations

import (
	"restful-api-joglo-dev/database"
	"restful-api-joglo-dev/model/entity"
)

func InitMigrations() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Book{}, &entity.Category{}, &entity.Image{})
	if err != nil {
		panic(err)
	}
}
