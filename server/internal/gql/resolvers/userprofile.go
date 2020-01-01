package resolvers

import (
	"context"
	"errors"
	authhandler "github.com/fedoratipper/bitkiosk/server/internal/authentication"
	"github.com/fedoratipper/bitkiosk/server/internal/authentication/session"
	"github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	tf "github.com/fedoratipper/bitkiosk/server/internal/gql/resolvers/transformations"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	dbm "github.com/fedoratipper/bitkiosk/server/internal/orm/models"
)

func (r *mutationResolver) UpdateUserProfile(ctx context.Context, input models.UpdatedProfile) (*models.UserProfile, error) {
	panic("implement me")
}

func (r *queryResolver) UserProfile(ctx context.Context, userId *int) (*models.UserProfile, error) {
	authLevel, err := authhandler.GetAuthLevelFromSession(ctx)

	if authLevel == nil {
		return nil, errors.New("not authenticated for query")
	}

	if authLevel.AuthLevel == session.UserAuth && *userId == int(authLevel.UID) {
		return getUserProfile(r, authLevel.UID)
	} else if authLevel.AuthLevel == session.AdminAuth	{
		return getUserProfile(r, authLevel.UID)
	} else {
		if err != nil {
			logger.Error("Unable to resolve auth level with Users resolver. \n" + err.Error())
		}
	}

	return nil, errors.New("not authenticated for query")
}

func getUserProfile(r *queryResolver, uid int) (*models.UserProfile, error) {
	var userProfile dbm.UserProfile
	db := r.ORM.DB.New()

	db.Where("user_id = ?", uid).First(&userProfile)
	closeErr := db.Close()

	if closeErr != nil {
		logger.Errorfn("getUserProfile", closeErr)
	}

	gqlUserProfile, err := tf.DBUserProfileToGQLUserProfile(&userProfile)

	if err != nil {
		return &models.UserProfile{}, err
	}

	return gqlUserProfile, nil

}
