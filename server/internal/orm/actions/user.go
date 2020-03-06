package actions

import (
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models"
	"github.com/jinzhu/gorm"
)

func GetUserWithEmail(email string, db *gorm.DB) *models.User {
	var user models.User

	db.Where("email = ?", email).Find(&user)

	user.Profile = LoadUserProfile(db, user.ID)

	return &user
}

func GetUserWithId(id uint, db *gorm.DB) *models.User {
	var user models.User

	db.Where("id = ?", id).Find(&user)

	user.Profile = LoadUserProfile(db, user.ID)

	return &user
}
