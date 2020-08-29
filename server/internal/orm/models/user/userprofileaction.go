package user

import (
	"github.com/fedoratipper/bitkiosk/server/internal/redis"
	"github.com/jinzhu/gorm"
	"strconv"
)

func LoadUserProfile(userId int, db *gorm.DB) *UserProfile {
	var profile UserProfile

	//profile = redis.LoadObjectFromCache(UserProfile{}, strconv.Itoa(int(userId))).(UserProfile)

	if profile.Id == 0 {
		db.Where("user_id = ?", userId).Find(&profile)
		if profile.Id != 0 {
			_ = redis.PutObjectInCache(profile, strconv.Itoa(userId))
		}
	}


	return &profile
}