package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type User struct {
	BaseModelSoftDelete
	FirstName string  `db:"first_name"`
	LastName  string  `db:"last_name"`
	Email     string  `db:"email" gorm:"unique_index;varchar(150)"`
}


func (toCreate *User) BeforeCreate(db *gorm.DB) (errs error) {
	users := []User{}

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