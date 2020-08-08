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

func (r *queryResolver) LoadReviewsForProduct(ctx context.Context, productSku string) ([]*models.Review, error) {
	var gqlReturn []*models.Review

	if productSku == "" {
		return nil, errors.New("Please provide a valid SKU for the product")
	}

	db := r.ORM.DB
	dboProduct := product.LoadProductFromSku(productSku, db)

	if dboProduct.ID == 0 {
		return nil, errors.New("Invalid product found for sku " + productSku)
	}

	reviews := review.LoadReviewsForProduct(dboProduct.ID, db)

	for _, dboReview := range reviews {
		gqlReview, err := tf.DBReviewToGQLReview(&dboReview, db)
		if err == nil {
			gqlReturn = append(gqlReturn, gqlReview)
		}
	}

	return gqlReturn, nil
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

	if dboUser.ID == 0 {
		return nil, errors.New("Unable to find user with " + input.Email + " for review")
	}

	if dboUser.ID != ctxUserId && ctxAuthLevel <= session.UserAuth {
		return nil, errors.New("Insufficient privilege to create a review on behalf of user with " + input.Email)
	}

	dboProduct := product.LoadProductFromSku(input.ProductSku, tx)

	if dboProduct.ID == 0 {
		return nil, errors.New("Unable to find product sku for review")
	}

	dboReview, err := tf.GQLReviewToDBReview(input, dboProduct.ID, dboUser.ID)

	if err != nil {
		return nil, err
	}

	tx, err = dboReview.Create(tx)

	if err == nil {
		gqlReturn, err = tf.DBReviewToGQLReview(dboReview, tx)
	}

	return gqlReturn, err
}
