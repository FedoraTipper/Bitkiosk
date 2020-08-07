// Package orm provides `GORM` helpers for the creation, migration and access
// on the project's database
package orm

import (
	log "github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/migration"
	"time"

	"github.com/fedoratipper/bitkiosk/server/pkg/utils"
	"github.com/jinzhu/gorm"
	//Imports the database dialect of choice
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var autoMigrate, logMode, seedDB bool
var dsn, dialect string

// ORM struct to holds the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

func init() {
	dialect = utils.MustGet("GORM_DIALECT")
	dsn = utils.MustGet("GORM_CONNECTION_DSN")
	seedDB = utils.MustGetBool("GORM_SEED_DB")
	logMode = utils.MustGetBool("GORM_LOGMODE")
	autoMigrate = utils.MustGetBool("GORM_AUTOMIGRATE")
}

// Factory creates a db connection with the selected dialect and connection string
func Factory() (*ORM, error) {
	db, err := gorm.Open(dialect, dsn)
	if err != nil {
		log.Panic("[ORM] err: ", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(time.Hour)

	orm := &ORM{
		DB: db,
	}
	// Log every SQL command on dev, @prod: this should be disabled?
	db.LogMode(logMode)
	// Automigrate tables
	if autoMigrate {
		err = migration.ServiceAutoMigration(orm.DB)
	}
	log.Info("[ORM] Database connection initialized.")
	return orm, err
}

func CommitOrRollBackIfError(db *gorm.DB, err error)  {
	if err == nil{
		db = db.Commit()
	} else {
		db = db.Rollback()
	}

	if db.Error != nil {
		log.Error(db.Error)
	}
}

func CommitOrRollBackIfErrorAndCloseSession(db *gorm.DB, err error) {
	CommitOrRollBackIfError(db, err)
	CloseDbConnectionLogIfError(db)
}

func CloseDbConnectionLogIfError(db *gorm.DB) {
	dbCloseErr := db.Close()

	if dbCloseErr != nil {
		log.Error("Unable to close db connection", dbCloseErr.Error())
	}
}
