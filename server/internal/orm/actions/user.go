package actions

import (
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models"
	"github.com/jinzhu/gorm"
)

func GetUserWithEmail(email string, db *gorm.DB) *models.User {
	var user models.User

	db = db.Where("email = ?", email).Find(&user)

	return &user
}