package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	BaseModelSoftDelete
	Email     string  `db:"email" gorm:"unique_index;varchar(150);index:user_email_idx"`
	Role	uint	`db:"role" gorm:"not null; default:1"`
}

type UserProfile struct {
	BaseModelSoftDelete
	UserID uint `db:"user_id" gorm:"index:user_id_profile_idx"`
	FirstName string  `db:"first_name"`
	LastName  string  `db:"last_name"`
	DateOfBirth time.Time `db:"date_of_birth"`
}

func (toCreate *User) BeforeCreate(db *gorm.DB) (errs error) {
	var users []User

	//Can't use same email
	db.Where("email = ?", toCreate.Email).Find(&users)

	if len(users) > 0 {
		return errors.New("email in use already")
	}

	return
}

func (toUpdate *User) BeforeUpdate(db *gorm.DB) (errs error) {
	var dbo User

	db.Where("email = ?", toUpdate.Email).Find(&dbo)

	if dbo.Email != "" {
		return errors.New("user not found")
	}

	return
}