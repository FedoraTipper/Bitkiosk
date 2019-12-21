// Package orm provides `GORM` helpers for the creation, migration and access
// on the project's database
package orm

import (
	log "github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/DBResult"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/migration"

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

func CommitOrRollBackIfError(db *gorm.DB, dbResult *DBResult.DBResult) (result *DBResult.DBResult){
	result = dbResult

	if result.IsOk() {
		db.Commit()
	} else {
		db.Rollback()
	}

	return result.AddErrorToResult(db.Error)
}

func CommitOrRollBackIfErrorAndCloseSession(db *gorm.DB, dbResult *DBResult.DBResult) (result *DBResult.DBResult){
	result = dbResult

	if result.IsOk() {
		db.Commit()
	} else {
		db.RollbackUnlessCommitted()
	}

	CloseDbConnectionLogIfError(db)

	return result.AddErrorToResult(db.Error)
}

func CloseDbConnectionLogIfError(db *gorm.DB) {
	dbCloseErr := db.Close()

	if dbCloseErr != nil {
		log.Error("Unable to close db connection", dbCloseErr.Error())
	}
}

func CreateObject(objToCreate interface{}, objToReturn interface{}, db *gorm.DB) (dbToReturn *gorm.DB, result *DBResult.DBResult) {
	result = DBResult.NewResult()
	dbToReturn = db.Create(objToCreate).First(objToReturn)

	if db.Error != nil {
		result = result.AddErrorToResult(db.Error)
	}

	return dbToReturn, result
}
