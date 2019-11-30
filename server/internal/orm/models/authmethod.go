package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type AuthMethod struct {
	BaseModelSoftDelete
	MethodId uint `db:"method_id" gorm:"not null"`
	Name string `db:"name" gorm:"not null"`
	TTL float32 `db:"ttl" gorm:"default:10.0"`// hours till refresh is needed
}



func (toCreate *AuthMethod) BeforeCreate(db *gorm.DB) (errs error) {
	authMethods := []AuthMethod{}

	//Can't use same email
	db.Where("name = ? OR method_id = ?", toCreate.Name, toCreate.MethodId).Find(&authMethods)

	if len(authMethods) > 0 {
		return errors.New("Authentication methods with same unique data already exist")
	}

	return
}