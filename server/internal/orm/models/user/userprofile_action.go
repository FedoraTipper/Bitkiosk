package user

import (
	"github.com/fedoratipper/bitkiosk/server/internal/redis"
	"github.com/jinzhu/gorm"
	"strconv"
)

func LoadUserProfile(db *gorm.DB, userId uint) *UserProfile {
	var profile UserProfile

	profile = redis.LoadObjectFromCache(UserProfile{}, strconv.Itoa(int(userId))).(UserProfile)

	if profile.ID == 0 {
		db.Where("user_id = ?", userId).Find(&profile)
		if profile.ID != 0 {
			_ = redis.PutObjectInCache(profile, strconv.Itoa(int(userId)))
		}
	}


	return &profile
}