package models

import (
	"errors"
	"github.com/fedoratipper/bitkiosk/server/internal/authentication/session"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
	"time"
)

type Product struct {
	BaseModelSoftDelete
	Name string `db:"name"`
	Description string `db:"description"`
	Price float64 `db:"price"`
	Stock uint `db:"stock"`
	Sku string `db:"sku" gorm:"index:product_sku_idx"`
	AdminId uint `db:"admin_id"`
	StartDate *time.Time `db:"start_date"`
	EndDate *time.Time `db:"end_date"`
}

func (toCreate *Product) Create(db *gorm.DB) (*gorm.DB, error) {
	return CreateObject(toCreate, toCreate, db)
}

func (toCreate *Product) BeforeCreate(db *gorm.DB) (error) {
	return toCreate.Validate(db, false)
}

func (tp *Product) Validate(db *gorm.DB, toInsert bool) error {
	return validation.ValidateStruct(
		tp,
		validation.Field(&tp.Name, validation.Required, validation.Length(1, 50)),
		validation.Field(&tp.Description, validation.Required, validation.Length(1, 255)),
		validation.Field(&tp.AdminId, validation.By(validateAdminExists(db))),
		validation.Field(&tp.StartDate, validation.Required),
		validation.Field(&tp.EndDate, validation.Required.When(tp.EndDate != nil && tp.StartDate.After(*tp.EndDate)).Error("Start date for products may not be after a given end date")),
		validation.Field(&tp.Stock, validation.Required.When(tp.Stock > 0).Error("Stock count needs to be greater than 0")),
		validation.Field(&tp.Sku, validation.By(validateSKUUniqueness(db, toInsert, tp.ID))),
		)
}

func validateAdminExists(db *gorm.DB) validation.RuleFunc {
	return func(value interface{}) error {
		adminId, _ := value.(uint)
		var lookupObj User

		db.Where("id = ? and role > ?", adminId, session.AdminAuth).First(&lookupObj)

		if lookupObj.ID != 0 {
			return errors.New("Invalid admin assigned to product")
		}

		return nil
	}
}

func validateSKUUniqueness(db *gorm.DB, toInsert bool, productId uint) validation.RuleFunc {
	return func(value interface{}) error {
		sku, _ := value.(string)
		var lookupObj Product

		if !toInsert {
			db.Where("name = ? and id != ?", sku, productId).First(&lookupObj)

			if lookupObj.ID != 0 {
				return errors.New("SKU name already exists. Please enter a unique SKU for this product")
			}
		}

		return nil
	}
}