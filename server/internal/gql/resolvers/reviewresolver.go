package resolvers

import (
	"context"
	"errors"
	authhandler "github.com/fedoratipper/bitkiosk/server/internal/authentication"
	"github.com/fedoratipper/bitkiosk/server/internal/authentication/session"
	"github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	_ "github.com/fedoratipper/bitkiosk/server/internal/gql/resolvers/transformations"
	tf "github.com/fedoratipper/bitkiosk/server/internal/gql/resolvers/transformations"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/fedoratipper/bitkiosk/server/internal/orm"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/product"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/review"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/user"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils/pointer"
	"github.com/jinzhu/gorm"
)

func (r *mutationResolver) CreateReview(ctx context.Context, input *models.NewReview) (*models.Review, error) {
	authLevel, err := authhandler.GetAuthLevelFromSession(ctx)

	if authLevel != nil && authLevel.AuthLevel >= session.UserAuth {
		return createReview(r, input, uint(authLevel.UID), uint(authLevel.AuthLevel))
	} else {
		if err != nil {
			logger.Error("Unable to resolve auth level with Products resolver. \n" + err.Error())
		}
	}

	return nil, errors.New("not authenticated for query")
}

func (r *queryResolver) LoadReviewsForProduct(ctx context.Context, productSku string, limit *int, offset *int) ([]*models.Review, error) {
	var gqlReturn []*models.Review

	if productSku == "" {
		return nil, errors.New("Please provide a valid SKU for the product")
	}

	db := r.ORM.DB
	dboProduct := product.LoadProductFromSku(productSku, db)

	if dboProduct.Id == 0 {
		return nil, errors.New("Invalid product found for sku " + productSku)
	}

	configLimit := utils.MustGetInt(review.LOOKUP_LIMIT)

	if limit == nil || *limit > configLimit {
		limit = &configLimit
	}

	reviews := review.LoadReviewsForProductWithLimitAndOffset(dboProduct.Id, pointer.DereferenceInt(limit), pointer.DereferenceInt(offset), db)

	for _, dboReview := range reviews {
		gqlReview, err := tf.DBReviewToGQLReview(&dboReview, db)
		if err == nil {
			gqlReturn = append(gqlReturn, gqlReview)
		}
	}

	return gqlReturn, nil
}

func (r *queryResolver) LoadTotalNumberOfReviewsForProduct(ctx context.Context, productSku string) (int, error) {
	if productSku == "" {
		return 0, errors.New("Please provide a valid SKU for the product")
	}

	db := r.ORM.DB
	dboProduct := product.LoadProductFromSku(productSku, db)

	if dboProduct.Id == 0 {
		return 0, errors.New("Invalid product found for sku " + productSku)
	}

	return review.LoadTotalReviewCountForProduct(dboProduct.Id, db), nil
}

func (r *queryResolver) LoadReviewForUserWithProductSku(ctx context.Context, productSku string) (gqlReview *models.Review, err error) {
	authLevel, err := authhandler.GetAuthLevelFromSession(ctx)

	if authLevel != nil && authLevel.AuthLevel >= session.UserAuth {
		db := r.ORM.DB

		reviewByUser, err := findUserSubmittedReview(productSku, uint(authLevel.UID), db)

		if err != nil {
			return nil, err
		}

		if reviewByUser.Id != 0 {
			gqlReview, err = tf.DBReviewToGQLReview(&reviewByUser, db)
		}

		return gqlReview, err
	} else {
		if err != nil {
			logger.Error("Unable to resolve auth level with Products resolver. \n" + err.Error())
		}
	}

	return gqlReview, nil
}

func findUserSubmittedReview(productSku string, userId uint, db *gorm.DB) (reviewByUser review.Review, _ error) {
	if productSku == "" {
		return reviewByUser, errors.New("Please provide a valid SKU for the product")
	}

	dboProduct := product.LoadProductFromSku(productSku, db)

	if dboProduct.Id == 0 {
		return reviewByUser, errors.New("Invalid product found for sku " + productSku)
	}

	reviewByUser = review.LoadProductReviewForUser(dboProduct.Id, userId, db)

	return reviewByUser, nil
}

func createReview(r *mutationResolver, input *models.NewReview, ctxUserId uint, ctxAuthLevel uint) (*models.Review, error){
	var gqlReturn *models.Review

	if input.ProductSku == "" {
		return nil, errors.New("Missing product SKU for review")
	}

	var err error = nil
	tx := r.ORM.DB.Begin()
	defer orm.CommitOrRollBackIfError(tx, err)

	dboUser := user.LoadUserWithEmail(input.Email, tx)

	if dboUser.Id == 0 {
		return nil, errors.New("Unable to find user with " + input.Email + " for review")
	}

	if dboUser.Id != ctxUserId && ctxAuthLevel <= session.UserAuth {
		return nil, errors.New("Insufficient privilege to create a review on behalf of user with " + input.Email)
	}

	dboProduct := product.LoadProductFromSku(input.ProductSku, tx)

	if dboProduct.Id == 0 {
		return nil, errors.New("Unable to find product sku for review")
	}

	dboReview, err := tf.GQLReviewToDBReview(input, dboProduct.Id, dboUser.Id)

	if err != nil {
		return nil, err
	}

	tx, err = dboReview.Create(tx)

	if err == nil {
		gqlReturn, err = tf.DBReviewToGQLReview(dboReview, tx)
	}

	return gqlReturn, err
}
