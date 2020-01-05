package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type UserProfile struct {
	BaseModelSoftDelete
	UserID uint `db:"user_id" gorm:"index:user_id_profile_idx"`
	FirstName *string  `db:"first_name"`
	LastName  *string  `db:"last_name"`
	DateOfBirth *time.Time `db:"date_of_birth"`
}

func (toCreate *UserProfile) Create(db *gorm.DB) (*gorm.DB, error) {
	err := toCreate.BeforeCreate(db)

	if err == nil {
		db, err = CreateObject(toCreate, toCreate, db)
	}

	return db, err
}

func (toCreate *UserProfile) BeforeCreate(db *gorm.DB) (err error) {
	var userProfiles []UserProfile

	//Can't use same email
	db.Where("user_id = ?", toCreate.UserID).Find(&userProfiles)

	if len(userProfiles) > 0 {
		return errors.New("user profile already exists for user")
	}

	return nil
}

func (toUpdate *UserProfile) BeforeUpdate(db *gorm.DB) (err error) {
	return
}
