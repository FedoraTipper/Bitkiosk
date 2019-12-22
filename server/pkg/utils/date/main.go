package date

import (
	"errors"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/DBResult"
	"time"
)

const (
	ISOLayout  = "2006-01-02"
)

func FormatToSqlDate(t *time.Time) *string {
	if t == nil {
		return nil
	}

	date := t.Format(ISOLayout)

	return &date
}


func ParseSqlDate(date string) (time.Time, *DBResult.DBResult){
	dbResult := DBResult.NewResult()

	t, err := time.Parse(ISOLayout, date)

	if err != nil {
		logger.Errorfn("ParseSqlDate", err)
		dbResult = dbResult.AddErrorToResult(errors.New("Unable to parse date  (" + date + "). Please ensure date is in YYYY-mm-DD format."))
	}

	return t, dbResult
}
