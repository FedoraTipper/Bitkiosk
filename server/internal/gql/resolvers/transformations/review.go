package transformations

import (
	"errors"
	gql "github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	product2 "github.com/fedoratipper/bitkiosk/server/internal/orm/models/product"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/review"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/user"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils/date"
	"github.com/jinzhu/gorm"
)

func GQLReviewToDBReview(i *gql.NewReview, productId uint, userId uint) (*review.Review, error) {
	dboReview := &review.Review{
		UserID:              userId,
		ProductID:			 productId,
		TextReview:          i.TextReview,
		Rating:              i.Rating,
	}

	return dboReview, nil
}

func DBReviewToGQLReview(i *review.Review, db *gorm.DB) (*gql.Review, error) {
	userProfile := user.LoadUserProfile(i.UserID, db)

	if userProfile.ID == 0 {
		logger.Errorfn("DBReviewToGQLReview", errors.New("Unable to find user profile with ID " + string(i.UserID) + " for review " + string(i.ID)))
		return nil, errors.New("Unable to find user profile for review")
	}

	product := product2.LoadProductWithId(i.ProductID, db)

	if product.ID == 0 {
		logger.Errorfn("DBReviewToGQLReview", errors.New("Unable to find product with ID " + string(i.ProductID) + " for review " + string(i.ID)))
		return nil, errors.New("Unable to find product for review")
	}

	gqlReview := gql.Review{
		UserName:   *userProfile.FirstName,
		ProductSku: product.Sku,
		TextReview: i.TextReview,
		Rating:     i.Rating,
		CreateAt:   *date.FormatToSqlDate(&i.CreatedAt),
	}

	return &gqlReview, nil
}