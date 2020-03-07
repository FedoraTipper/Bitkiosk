package actions

import (
	dbm "github.com/fedoratipper/bitkiosk/server/internal/orm/models"
	"github.com/jinzhu/gorm"
)

func LoadActiveProducts(db *gorm.DB) []dbm.Product{
	var activeProducts []dbm.Product

	db.Where("start_date <= now() and (end_date is NULL or (start_date != end_date and end_date > now()))").Find(&activeProducts)

	return activeProducts
}