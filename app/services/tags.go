package services

import (
	"github.com/KatsutoshiOtogawa/GinServer/infrastructure"
	"github.com/KatsutoshiOtogawa/GinServer/models"
)

func FetchAllTags() ([]models.Tag, error) {
	database := infrastructure.GetDb()
	var tags []models.Tag
	err := database.Preload("Images", "tag_id IS NOT NULL").Find(&tags).Error
	return tags, err
}
