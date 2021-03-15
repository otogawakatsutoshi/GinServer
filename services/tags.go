package services

import (
	"https://github.com/KatsutoshiOtogawa/GinServer/infrastructure"
	"https://github.com/KatsutoshiOtogawa/GinServer/models"
)

func FetchAllTags() ([]models.Tag, error) {
	database := infrastructure.GetDb()
	var tags []models.Tag
	err := database.Preload("Images", "tag_id IS NOT NULL").Find(&tags).Error
	return tags, err
}
