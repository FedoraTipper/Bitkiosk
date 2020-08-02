package review

import (
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/product"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/user"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
)

type Review struct {
	models.BaseModelSoftDelete
	UserID		uint        `gorm:"not null;index:user_review_idx" db:"user_id"`
	ProductID 	uint		`gorm:"not null;index:product_review_idx" db:"product_id"`
	TextReview	string		`gorm:"text" db:"text_review"`
	Rating	 	int			`gorm:"not null;" db:"rating"`
}
func (toCreate *Review) Create(db *gorm.DB) (*gorm.DB, error) {
	return models.CreateObject(toCreate, toCreate, db)
}

func (toCreate *Review) BeforeCreate(db *gorm.DB) error {
	return toCreate.Validate(db, false)
}

func (tr *Review) Validate(db *gorm.DB, toInsert bool) error {
	return validation.ValidateStruct(
		tr,
		validation.Field(&tr.UserID, validation.By(user.ValidateUserExistence(db, toInsert))),
		validation.Field(&tr.ProductID, validation.By(product.ValidateProductExistence(db))),
		validation.Field(&tr.TextReview, validation.Length(0, 4096)),
		validation.Field(&tr.Rating, validation.Max(5)),
		validation.Field(&tr.Rating, validation.Min(0)),
	)
}
