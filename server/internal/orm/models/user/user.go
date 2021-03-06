package user

import (
	"errors"
	"github.com/fedoratipper/bitkiosk/server/internal/authentication/session"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jinzhu/gorm"
)

type User struct {
	models.BaseModelSoftDelete
	Email       string `db:"email" gorm:"unique_index;varchar(150);index:user_email_idx"`
	Role        int   `db:"role" gorm:"not null; default:1"`
	UserProfile *UserProfile
}

func (toCreate *User) Create(db *gorm.DB) (*gorm.DB, error) {
	return models.CreateObject(toCreate, toCreate, db)
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
			validation.Field(&u.Email, validation.By(ValidateEmailUniqueness(db, toInsert, u.Id))),
			validation.Field(&u.Role, validation.Required.When(u.Role < 0).Error("Role needs to be greater than 0")),
		)
}

func ValidateEmailUniqueness(db *gorm.DB, toInsert bool, id int) validation.RuleFunc {
	return func(value interface{}) error {
		email, _ := value.(string)
		var lookupObj User

		if toInsert {
			db.Where("email = ?", email).Find(&lookupObj)
		} else {
			db.Where("email = ? and id != ?", email, id).First(&lookupObj)
		}

		if lookupObj.Id != 0 && toInsert {
			return errors.New("Email address already in use")
		} else if lookupObj.Id != 0 && !toInsert {
			return errors.New("New email address already in use")
		}

		return nil
	}
}

func ValidateUserExistence(db *gorm.DB, toInsert bool) validation.RuleFunc {
	return func(value interface{}) error {
		userId, _ := value.(int)
		var lookupObj User

		if toInsert {
			lookupObj = *LoadUserWithId(userId, db)

			if lookupObj.Id == 0 {
				return errors.New("Unable to find user")
			}
		}
		return nil
	}
}

func ValidateAdminExists(db *gorm.DB, objectName string) validation.RuleFunc {
	return func(value interface{}) error {
		adminId, _ := value.(int)
		var lookupObj User

		db.Where("id = ? and role > ?", adminId, session.AdminAuth).First(&lookupObj)

		if lookupObj.Id != 0 {
			return errors.New("Invalid admin assigned to " + objectName)
		}

		return nil
	}
}
