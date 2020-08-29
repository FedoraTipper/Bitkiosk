package product

import (
	"errors"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/user"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
	"regexp"
	"time"
)

type Product struct {
	models.BaseModelSoftDelete
	Name             string     `db:"name" gorm:"varchar(50)"`
	Description		 string		`db:"description" gorm:"text"`
	ShortDescription string     `db:"short_description" gorm:"varchar(255)"`
	Price            float64    `db:"price"`
	Stock            int        `db:"stock"`
	Sku              string     `db:"sku" gorm:"unique_index;index:product_sku_idx"`
	AdminId          int       `db:"admin_id"`
	StartDate        *time.Time `db:"start_date"`
	EndDate          *time.Time `db:"end_date"`
}

func (toCreate *Product) Create(db *gorm.DB) (*gorm.DB, error) {
	return models.CreateObject(toCreate, toCreate, db)
}

func (toCreate *Product) BeforeCreate(db *gorm.DB) error {
	return toCreate.Validate(db, false)
}

func (tp *Product) Validate(db *gorm.DB, toInsert bool) error {
	return validation.ValidateStruct(
		tp,
		validation.Field(&tp.Name, validation.Required, validation.Length(1, 50)),
		validation.Field(&tp.ShortDescription, validation.Required, validation.Length(1, 255)),
		validation.Field(&tp.Description, validation.Required, validation.Length(1, 4096)),
		validation.Field(&tp.AdminId, validation.By(user.ValidateAdminExists(db, "product"))),
		validation.Field(&tp.StartDate, validation.Required),
		validation.Field(&tp.EndDate, validation.By(validateEndDate(tp.StartDate))),
		validation.Field(&tp.Stock, validation.Min(0).Error("Stock count needs to be greater than 0")),
		validation.Field(&tp.Price, validation.Min( float64(0)).Error("Price needs to be greater than 0")),
		validation.Field(&tp.Sku, validation.By(validateSKUUniqueness(db, toInsert, tp.Id))),
		validation.Field(&tp.Sku, validation.By(validateSKUString())),
		)
}

func ValidateProductSkuExistence(db *gorm.DB) validation.RuleFunc {
	return func(value interface{}) error {
		sku, _ := value.(string)

		if sku == "" {
			return errors.New("Empty SKU supplied for product lookup")
		}

		product := LoadProductFromSku(sku, db)

		if product.Id == 0 {
			return errors.New("Unable to find product with SKU " + sku)
		}

		return nil
	}
}

func ValidateProductExistence(db *gorm.DB) validation.RuleFunc {
	return func(value interface{}) error {
		id, _ := value.(int)

		if id == 0 {
			return errors.New("Missing product id")
		}

		product := LoadProductWithId(id, db)

		if product.Id == 0 {
			logger.Errorfn("ValidateProductExistence", errors.New("Unable to find product with id " + string(id)))
			return errors.New("Unable to find product")
		}

		return nil
	}
}

func validateEndDate(startDate *time.Time) validation.RuleFunc {
	return func(value interface{}) error {
		endDate, _ := value.(*time.Time)

		if endDate != nil && startDate.After(*endDate) {
			return errors.New("Start date for products may not be after a given end date")
		}

		return nil
	}
}


func validateSKUString() validation.RuleFunc {
	return func(value interface{}) error {
		sku, _ := value.(string)

		var expression string = `[^\w\d\s-]`

		regex, err := regexp.Compile(expression)

		if err != nil {
			logger.Error("Unable to parse regex: \n " + expression)
			return errors.New("Unable to perform sku validation")
		}

		if len(regex.FindAllString(sku, -1)) > 0 {
			return errors.New("SKU identifiers may not contain special characters. All spaces will be replaced with hyphens")
		}

		return nil
	}
}


func validateSKUUniqueness(db *gorm.DB, toInsert bool, productId int) validation.RuleFunc {
	return func(value interface{}) error {
		sku, _ := value.(string)
		var lookupObj Product

		if !toInsert {
			db.Where("sku = ? and id != ?", sku, productId).First(&lookupObj)

			if lookupObj.Id != 0 {
				return errors.New("SKU identifiers already exists. Please enter a unique SKU for this product")
			}
		}

		return nil
	}
}