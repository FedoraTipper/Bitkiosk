package resolvers

import (
	"context"
	"github.com/fedoratipper/bitkiosk/server/internal/gql/models"
)

func (r *mutationResolver) CreateReview(ctx context.Context, input *models.NewReview) (string, error) {
	panic("implement me")
}

func (r *queryResolver) LoadReviewsForProduct(ctx context.Context, productSku string) ([]*models.Review, error) {
	panic("implement me")
}
