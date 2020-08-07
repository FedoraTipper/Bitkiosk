package user

import (
	"github.com/jinzhu/gorm"
)

func LoadUserWithEmail(email string, db *gorm.DB) *User {
	var user User

	db.Where("email = ?", email).Preload("UserProfile").Find(&user)

	return &user
}

func LoadUserWithId(id uint, db *gorm.DB) *User {
	var user User

	//cacheObject := redis.LoadObjectFromCache(models.User{}, strconv.Itoa(int(id)))
	//if cacheObject != nil {
	//	user = reflect.ValueOf(cacheObject).Interface().(models.User)
	//}

	//if user.ID == 0 {
	db.Where("id = ?", id).Preload("UserProfile").Find(&user)
		//if user.ID != 0 {
		//	_ = redis.PutObjectInCache(user, strconv.Itoa(int(user.ID)))
		//}
	//}


	return &user
}