package models

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
)

type UserProfile struct {
	BaseModelSoftDelete
	UserID uint `db:"user_id" gorm:"index:user_id_profile_idx"`
	FirstName *string  `db:"first_name"`
	LastName  *string  `db:"last_name"`
}

func (toCreate *UserProfile) Create(db *gorm.DB) (*gorm.DB, error) {
	err := toCreate.BeforeCreate(db)

	if err == nil {
		db, err = CreateObject(toCreate, toCreate, db)
	}

	return db, err
}

func (toCreate *UserProfile) BeforeCreate(db *gorm.DB) (err error) {
	err = toCreate.Validate()

	if err != nil {
		return err
	}

	var userProfiles []UserProfile

	db.Where("user_id = ?", toCreate.UserID).Find(&userProfiles)

	if len(userProfiles) > 0 {
		return errors.New("User profile already exists for user")
	}

	return nil
}

func (toUpdate *UserProfile) BeforeUpdate(db *gorm.DB) (err error) {
	err = toUpdate.Validate()

	if err != nil {
		return err
	}

	return
}

func (up *UserProfile) Validate() error {
	return validation.ValidateStruct(
			up,
			validation.Field(&up.FirstName, validation.Length(0, 50)),
			validation.Field(&up.LastName, validation.Length(0, 50)),
			validation.Field(&up.UserID, validation.Required.When(up.UserID < 0).Error("UserID needs to be greater than 0")),
		)
}

