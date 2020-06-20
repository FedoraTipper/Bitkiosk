package jobs

import (
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/user"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var (
	uname                  = "Test User"
	fname                  = "Test"
	lname                  = "User"
	nname                  = "Foo Bar"
	description            = "This is the first user ever!"
	location               = "His house, maybe?"
	firstUser   *user.User = &user.User{
		Email:     "test@test.com",
		Role: 0,
	}
)

// SeedUsers inserts the first users
var SeedUsers *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_USERS",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstUser).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstUser).Error
	},
}
