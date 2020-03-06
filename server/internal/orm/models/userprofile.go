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
	return CreateObject(toCreate, toCreate, db)
}

func (toCreate *UserProfile) BeforeCreate(db *gorm.DB) (err error) {
	return toCreate.Validate(db, true)
}

func (toUpdate *UserProfile) BeforeUpdate(db *gorm.DB) (err error) {
	return toUpdate.Validate(db, false)
}

func (up *UserProfile) Validate(db *gorm.DB, toInsert bool) error {
	return validation.ValidateStruct(
			up,
			validation.Field(&up.FirstName, validation.Length(0, 50)),
			validation.Field(&up.LastName, validation.Length(0, 50)),
			validation.Field(&up.UserID, validation.By(validateUserExistence(db, toInsert))),
			validation.Field(&up.UserID, validation.By(validateUserProfileConstraint(db, toInsert, up.ID))),
			)
}

func validateUserExistence(db *gorm.DB, toInsert bool) validation.RuleFunc {
	return func(value interface{}) error {
		userId, _ := value.(uint)
		var lookupObj User

		if toInsert {
			db.Where("id = ?", userId).First(&lookupObj)

			if lookupObj.ID != 0 {
				return errors.New("Unable to find user")
			}
		}
		return nil
	}
}

// One profile per user - Used when updating
func validateUserProfileConstraint(db *gorm.DB, toInsert bool, profileId uint) validation.RuleFunc {
	return func(value interface{}) error {
		userId, _ := value.(uint)
		var lookupObj User

		if !toInsert {
			db.Where("user_id = ? and id != ?", userId, profileId).First(lookupObj)

			if lookupObj.ID != 0 {
				return errors.New("User profile already exists for user")
			}
		}

		return nil
	}
}

