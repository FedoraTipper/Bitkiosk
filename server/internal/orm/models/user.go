package models

import (
	"errors"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jinzhu/gorm"
)

type User struct {
	BaseModelSoftDelete
	Email       string  `db:"email" gorm:"unique_index;varchar(150);index:user_email_idx"`
	Role        uint	`db:"role" gorm:"not null; default:1"`
	UserProfile *UserProfile
}

func (toCreate *User) Create(db *gorm.DB) (*gorm.DB, error) {
	return CreateObject(toCreate, toCreate, db)
}

func (toCreate *User) BeforeCreate(db *gorm.DB) (err error) {
	return toCreate.Validate(db, true)
}

func (toUpdate *User) BeforeUpdate(db *gorm.DB) (err error) {
	return toUpdate.Validate(db, false)
}

func (u *User) Validate(db *gorm.DB, toInsert bool) error {
	return validation.ValidateStruct(
			u,
			validation.Field(&u.Email, validation.Required, is.Email),
			validation.Field(&u.Email, validation.By(validateEmailUniqueness(db, toInsert, u.ID))),
			validation.Field(&u.Role, validation.Required.When(u.Role < 0).Error("Role needs to be greater than 0")),
		)
}

func validateEmailUniqueness(db *gorm.DB, toInsert bool, id uint) validation.RuleFunc {
	return func(value interface{}) error {
		email, _ := value.(string)
		var lookupObj User

		if toInsert {
			db.Where("email = ?", email).Find(&lookupObj)
		} else {
			db.Where("email = ? and id != ?", email, id).First(&lookupObj)
		}

		if lookupObj.ID != 0 && toInsert {
			return errors.New("Email address already in use")
		} else if lookupObj.ID != 0 && !toInsert {
			return errors.New("New email address already in use")
		}

		return nil
	}
}