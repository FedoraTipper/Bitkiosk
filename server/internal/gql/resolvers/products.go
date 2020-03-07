package resolvers

import (
	"context"
	"errors"
	authhandler "github.com/fedoratipper/bitkiosk/server/internal/authentication"
	"github.com/fedoratipper/bitkiosk/server/internal/authentication/session"
	"github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	tf "github.com/fedoratipper/bitkiosk/server/internal/gql/resolvers/transformations"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/fedoratipper/bitkiosk/server/internal/orm"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/actions"
)

func (r *queryResolver) LoadActiveProducts(ctx context.Context, limit *int, offset *int) ([]*models.Product, error) {
	var gqlReturn []*models.Product

	db := r.ORM.DB.New().Begin()

	products := actions.LoadActiveProducts(db)

	for _,product := range products {
		if gqlProduct, err := tf.DBProductToGQLProduct(&product, db); err == nil {
			gqlReturn = append(gqlReturn, gqlProduct)
		} else {
			return nil, err
		}
	}

	return gqlReturn, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input *models.NewProduct) (*models.Product, error) {
	authLevel, err := authhandler.GetAuthLevelFromSession(ctx)

	if authLevel != nil && authLevel.AuthLevel == session.AdminAuth {
		return createProduct(r, input, uint(authLevel.UID))
	} else {
		if err != nil {
			logger.Error("Unable to resolve auth level with Products resolver. \n" + err.Error())
		}
	}

	return nil, errors.New("not authenticated for query")
}

func createProduct(r *mutationResolver, input *models.NewProduct, adminId uint) (*models.Product, error){
	var gqlReturn *models.Product

	productDbo, err := tf.GQLProductToDBProduct(input, adminId)

	if err != nil {
		return nil, err
	}

	db := r.ORM.DB.New().Begin()

	db, err = productDbo.Create(db)

	if err == nil {
		gqlReturn, err = tf.DBProductToGQLProduct(productDbo, db)
	}

	orm.CommitOrRollBackIfError(db, err)

	return gqlReturn, err
}