package actions

import (
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models"
	"github.com/fedoratipper/bitkiosk/server/internal/redis"
	"github.com/jinzhu/gorm"
	"strconv"
)

func LoadUserProfile(db *gorm.DB, userId uint) *models.UserProfile {
	var profile models.UserProfile

	profile = redis.LoadObjectFromCache(models.UserProfile{}, strconv.Itoa(int(userId))).(models.UserProfile)

	if profile.ID == 0 {
		db.Where("user_id = ?", userId).Find(&profile)
		if profile.ID != 0 {
			_ = redis.PutObjectInCache(profile, strconv.Itoa(int(userId)))
		}
	}


	return &profile
}