package actions

import (
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models"
	"github.com/jinzhu/gorm"
)

func LoadUserProfile(db *gorm.DB, userId uint) *models.UserProfile {
	var profile models.UserProfile

	db.Where("user_id = ?", userId).Find(&profile)

	return &profile
}