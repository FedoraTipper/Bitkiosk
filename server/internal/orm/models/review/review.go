package review

import (
	"errors"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/product"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/user"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
)

type Review struct {
	models.BaseModelSoftDelete
	UserID		int        `gorm:"not null;index:user_review_idx" db:"user_id"`
	ProductID 	int		`gorm:"not null;index:product_review_idx" db:"product_id"`
	TextReview	string		`gorm:"text" db:"text_review"`
	Rating	 	int			`gorm:"not null;" db:"rating"`
	Anonymous   bool		`gorm:"default:False;"`
}

const (
	DisplayNameAnonymous string = "Anonymous"
	LookupLimit          string = "REVIEW_LIMIT"
)

func (toCreate *Review) Create(db *gorm.DB) (*gorm.DB, error) {
	return models.CreateObject(toCreate, toCreate, db)
}

func (toCreate *Review) BeforeCreate(db *gorm.DB) error {
	return toCreate.Validate(db, true)
}

func (tr *Review) Validate(db *gorm.DB, toInsert bool) error {
	err := validation.ValidateStruct(
		tr,
		validation.Field(&tr.UserID, validation.By(user.ValidateUserExistence(db, toInsert))),
		validation.Field(&tr.ProductID, validation.By(product.ValidateProductExistence(db))),
		validation.Field(&tr.TextReview, validation.Length(0, 4096)),
		validation.Field(&tr.Rating, validation.Max(5)),
		validation.Field(&tr.Rating, validation.Min(0)),
	)

	if err == nil {
		err = validation.Validate(userAndProductPayload{userId: tr.UserID, productId: tr.ProductID}, validation.By(ValidateReviewUniqueness(db, toInsert)))
	}

	return err
}

type userAndProductPayload struct {
	userId    int
	productId int
}

func ValidateReviewUniqueness(db *gorm.DB, toInsert bool) validation.RuleFunc {
	return func(value interface{}) error {
		payload, _ := value.(userAndProductPayload)

		if payload.productId == 0 {
			return errors.New("Missing product for review")
		}

		if payload.userId == 0 {
			return errors.New("Missing user for review")
		}

		review := LoadProductReviewForUser(payload.productId, payload.userId, db)

		if review.Id != 0 && toInsert {
			return errors.New("Only one review per product is allowed")
		} else if review.Id == 0 && !toInsert {
			return errors.New("Unable to find current review to update")
		}

		return nil
	}
}
