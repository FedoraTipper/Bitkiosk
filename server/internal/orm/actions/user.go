package actions

import (
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models"
	"github.com/jinzhu/gorm"
)

func LoadUserWithEmail(email string, db *gorm.DB) *models.User {
	var user models.User

	db.Where("email = ?", email).Preload("UserProfile").Find(&user)

	return &user
}

func LoadUserWithId(id uint, db *gorm.DB) *models.User {
	var user models.User

	db.Where("id = ?", id).Preload("UserProfile").Find(&user)

	return &user
}
