package models

import (
	"github.com/fedoratipper/bitkiosk/server/internal/orm/DBResult"
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

func (toCreate *AuthenticationMatrix) CommitToDb(db *gorm.DB) (authenticationMatrixToReturn *AuthenticationMatrix, dbToReturn *gorm.DB, result *DBResult.DBResult) {
	result = DBResult.NewResult()

	dbToReturn = db.Create(toCreate).First(authenticationMatrixToReturn)

	if db.Error != nil {
		result = result.AddErrorToResult(db.Error)
	}

	return authenticationMatrixToReturn, dbToReturn, result
}
