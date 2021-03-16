package services

import (
	"github.com/KatsutoshiOtogawa/GinServer/infrastructure"
	"github.com/KatsutoshiOtogawa/GinServer/models"
)

func FetchAllCategories() ([]models.Category, error) {
	database := infrastructure.GetDb()
	var categories []models.Category
	err := database.Preload("Images", "category_id IS NOT NULL").Find(&categories).Error
	return categories, err
}
