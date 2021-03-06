package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BaseModel struct {
	Id        int       `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time  `gorm:"index;not null;default:CURRENT_TIMESTAMP"` // (My|Postgre)SQL
	UpdatedAt *time.Time `gorm:"index"`
}

// BaseModelSoftDelete defines the common columns that all db structs should
// hold, usually. This struct also defines the fields for GORM triggers to
// detect the entity should soft delete
type BaseModelSoftDelete struct {
	BaseModel
	DeletedAt *time.Time `sql:"index"`
}

func CreateObject(objToCreate interface{}, objToReturn interface{}, db *gorm.DB) (dbToReturn *gorm.DB, err error) {
	dbToReturn = db.Create(objToCreate).First(objToReturn)

	if dbToReturn.Error != nil {
		err = dbToReturn.Error
	}

	return dbToReturn, err
}
