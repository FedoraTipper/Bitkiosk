package resolvers

import (
	"context"
	"github.com/fedoratipper/bitkiosk/server/internal/gql/models"
)


func (r *queryResolver) Authenticate(ctx context.Context, authDetails models.LoginDetails) (*models.AuthResponse, error) {
	return nil, nil
}