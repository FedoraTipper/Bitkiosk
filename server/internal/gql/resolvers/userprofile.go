package resolvers

import (
	"context"
	"errors"
	authhandler "github.com/fedoratipper/bitkiosk/server/internal/authentication"
	"github.com/fedoratipper/bitkiosk/server/internal/authentication/session"
	gqlModels "github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	tf "github.com/fedoratipper/bitkiosk/server/internal/gql/resolvers/transformations"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/fedoratipper/bitkiosk/server/internal/orm"
	dbm "github.com/fedoratipper/bitkiosk/server/internal/orm/models"
)

func (r *mutationResolver) UpdateUserProfile(ctx context.Context, input gqlModels.UpdatedProfile) (*gqlModels.UserProfile, error) {
	authLevel, err := authhandler.GetAuthLevelFromSession(ctx)

	if authLevel == nil {
		return nil, errors.New("not authenticated for query")
	}

	if authLevel.AuthLevel == session.UserAuth && authLevel.UID == int(authLevel.UID) {
		return updateUserProfile(r, input)
	} else if authLevel.AuthLevel == session.NoAuth	{
		return updateUserProfile(r, input)
	} else {
		if err != nil {
			logger.Error("Unable to resolve auth level with Users resolver. \n" + err.Error())
		}
	}

	return nil, errors.New("not authenticated for query")
}

func (r *queryResolver) UserProfile(ctx context.Context, userId *int) (*gqlModels.UserProfile, error) {
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

func getUserProfile(r *queryResolver, uid int) (*gqlModels.UserProfile, error) {
	var userProfile dbm.UserProfile
	db := r.ORM.DB.New()

	db.Where("user_id = ?", uid).First(&userProfile)
	closeErr := db.Close()

	if closeErr != nil {
		logger.Errorfn("getUserProfile", closeErr)
	}

	gqlUserProfile, err := tf.DBUserProfileToGQLUserProfile(&userProfile)

	if err != nil {
		return &gqlModels.UserProfile{}, err
	}

	return gqlUserProfile, nil

}

func updateUserProfile(r *mutationResolver, input gqlModels.UpdatedProfile) (*gqlModels.UserProfile, error) {

	db := r.ORM.DB.New().Begin()

	dbo , err := tf.GQLUpdateUserProfileToDBUserProfile(&input, db)

	if err != nil {
		return nil, err
	}

	db = db.Save(&dbo)

	if db.Error != nil {
		return nil, errors.New("unable to update user profile")
	}

	gql, err := tf.DBUserProfileToGQLUserProfile(dbo)

	err = orm.CommitOrRollBackIfError(db, err)

	return gql, err
}
