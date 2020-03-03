package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type AuthenticationMatrix struct {
	BaseModelSoftDelete
	UserID       uint       `gorm:"not null;index:user_auth_matrix_idx" db:"user_id"`
	User         User       `gorm:"-"`
	AuthMethodID uint       `gorm:"not null" db:"auth_method_id"`
	AuthMethod   AuthMethod `gorm:"-"`
	Token        string     `gorm:"not null" db:"token"`
}

func (toCreate *AuthenticationMatrix) Create(db *gorm.DB) (*gorm.DB, error) {
	err := toCreate.BeforeCreate(db)

	if err == nil {
		db, err = CreateObject(toCreate, toCreate, db)
	}

	return db, err
}

func (toCreate *AuthenticationMatrix) BeforeCreate(db *gorm.DB) (err error) {
	var authMatrices []AuthenticationMatrix

	if GetAuthMethod(int(toCreate.AuthMethodID)) == nil {
		return errors.New("Unable to find authentication method")
	}

	//Can't use same email
	db.Where("user_id = ? and auth_method_id = ?", toCreate.UserID, toCreate.AuthMethodID).Find(&authMatrices)

	if len(authMatrices) > 0 {
		return errors.New("Authentication matrix for user already exists")
	}

	return nil
}
