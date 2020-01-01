package date

import (
	"errors"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
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

func ParseSqlDate(date string) (*time.Time, error){
	t, err := time.Parse(ISOLayout, date)

	if err != nil {
		logger.Errorfn("ParseSqlDate", err)
		// replace error with user friendly version
		err = errors.New("Unable to parse date  (" + date + "). Please ensure date is in YYYY-mm-DD format.")
	}

	return &t, err
}
