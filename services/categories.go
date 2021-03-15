package services

import (
	"https://github.com/KatsutoshiOtogawa/GinServer/infrastructure"
	"https://github.com/KatsutoshiOtogawa/GinServer/models"
)

func FetchAllCategories() ([]models.Category, error) {
	database := infrastructure.GetDb()
	var categories []models.Category
	err := database.Preload("Images", "category_id IS NOT NULL").Find(&categories).Error
	return categories, err
}
