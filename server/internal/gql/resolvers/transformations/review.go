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

func GQLReviewToDBReview(i *gql.NewReview, productId int, userId int) (*review.Review, error) {

	isAnonymousReview := false; if i.Anonymous != nil {
		isAnonymousReview = *i.Anonymous
	}

	dboReview := &review.Review{
		UserID:              userId,
		ProductID:			 productId,
		TextReview:          i.TextReview,
		Rating:              i.Rating,
		Anonymous:			 isAnonymousReview,
	}

	return dboReview, nil
}

func DBReviewToGQLReview(i *review.Review, db *gorm.DB) (*gql.Review, error) {
	userDisplayName := review.DisplayNameAnonymous

	if !i.Anonymous {
		userProfile := user.LoadUserProfile(i.UserID, db)

		if userProfile.Id == 0 {
			logger.Errorfn("DBReviewToGQLReview", errors.New("Unable to find user profile with Id " + string(i.UserID) + " for review " + string(i.Id)))
			return nil, errors.New("Unable to find user profile for review")
		}

		userDisplayName = *userProfile.FirstName
	}

	product := product2.LoadProductWithId(i.ProductID, db)

	if product.Id == 0 {
		logger.Errorfn("DBReviewToGQLReview", errors.New("Unable to find product with Id " + string(i.ProductID) + " for review " + string(i.Id)))
		return nil, errors.New("Unable to find product for review")
	}

	gqlReview := gql.Review{
		UserDisplayName:   userDisplayName,
		ProductSku: product.Sku,
		TextReview: i.TextReview,
		Rating:     i.Rating,
		Anonymous:	i.Anonymous,
		CreatedAt:   *date.FormatToSqlDate(&i.CreatedAt),
	}

	return &gqlReview, nil
}
