package migration

import (
	"fmt"
	log "github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/migration/jobs"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/auth"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/product"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/review"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/user"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func updateMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&user.User{},
		&auth.AuthenticationMatrix{},
		&auth.AuthMethod{},
		&user.UserProfile{},
		&product.Product{},
		&review.Review{},
	).Error
}

func updateConstraints(db *gorm.DB) error {
	if err := db.Model(&auth.AuthenticationMatrix{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return err
	}
	if err := db.Model(&user.UserProfile{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return err
	}
	if err := db.Model(&product.Product{}).AddForeignKey("admin_id", "users(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return err
	}
	if err := db.Model(&review.Review{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return err
	}
	return nil
}

// ServiceAutoMigration migrates all the tables and modifications to the connected source
func ServiceAutoMigration(db *gorm.DB) error {
	// Keep a list of migrations here
	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		log.Info("[Migration.InitSchema] Initializing database schema")
		if err := updateMigration(db); err != nil {
			return fmt.Errorf("[Migration.InitSchema]: %v", err)
		}

		if err := updateConstraints(db); err != nil {
			return fmt.Errorf("[Contraints.InitSchema]: %v", err)
		}

		// Add more jobs, etc here
		return nil
	})
	m.Migrate()

	if err := updateMigration(db); err != nil {
		return err
	}

	if err := updateConstraints(db); err != nil {
		return err
	}

	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		jobs.SeedUsers,
	})
	return m.Migrate()
}

