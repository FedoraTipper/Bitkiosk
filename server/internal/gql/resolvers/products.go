package resolvers

import (
	"context"
	"github.com/fedoratipper/bitkiosk/server/internal/gql/models"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input *models.NewProduct) (*models.Product, error) {
	panic("implement me")
}