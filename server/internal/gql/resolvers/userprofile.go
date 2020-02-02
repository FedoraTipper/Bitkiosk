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
	"github.com/fedoratipper/bitkiosk/server/internal/orm/actions"
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

func (r *queryResolver) UserProfile(ctx context.Context, email *string) (*gqlModels.UserProfile, error) {
	authLevel, err := authhandler.GetAuthLevelFromSession(ctx)

	if authLevel == nil || authLevel.AuthLevel == session.NoAuth {
		return nil, errors.New("not authenticated for query")
	}

	var userToFind *dbm.User

	userToFind = actions.GetUserWithEmail(*email, r.ORM.DB.New())

	if userToFind == nil {
		return nil, errors.New("Unable to find user with email " + *email)
	}

	if authLevel.AuthLevel == session.UserAuth && int(userToFind.ID) == authLevel.UID {
		return getUserProfile(r, int(userToFind.ID))
	} else if authLevel.AuthLevel == session.AdminAuth	{
		return getUserProfile(r, int(userToFind.ID))
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
