package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type User struct {
	BaseModelSoftDelete
	Email     string  `db:"email" gorm:"unique_index;varchar(150);index:user_email_idx"`
	Role	uint	`db:"role" gorm:"not null; default:1"`
}


func (toCreate *User) Create(db *gorm.DB) (*gorm.DB, error) {
	err := toCreate.BeforeCreate(db)

	if err == nil {
		db, err = CreateObject(toCreate, toCreate, db)
	}

	return db, err
}

func (toCreate *User) BeforeCreate(db *gorm.DB) (err error) {
	var users []User

	//Can't use same email
	db.Where("email = ?", toCreate.Email).Find(&users)

	if len(users) > 0 {
		return errors.New("email in use already")
	}

	return nil
}

func (toUpdate *User) BeforeUpdate(db *gorm.DB) (err error) {
	var dbo User

	db.Where("email = ?", toUpdate.Email).Find(&dbo)

	if dbo.Email != "" {
		return errors.New("user not found")
	}

	return
}

