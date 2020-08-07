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
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/user"
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

func (r *queryResolver) UserProfile(ctx context.Context, email *string) (*gqlModels.User, error) {
	authLevel, err := authhandler.GetAuthLevelFromSession(ctx)

	if authLevel == nil || authLevel.AuthLevel == session.NoAuth || err != nil {
		return nil, errors.New("not authenticated for query")
	}

	var userToFind *user.User

	// TODO MOVE THIS SOMEWHERE ELSE
	if email != nil && *email != "" {
		if authLevel.AuthLevel == session.AdminAuth {
			userToFind = user.LoadUserWithEmail(*email, r.ORM.DB)
		} else {
			return nil, errors.New("Unable to access user information")
		}
	} else {
		userToFind = user.LoadUserWithId(uint(authLevel.UID), r.ORM.DB)
	}

	if userToFind == nil {
		return nil, errors.New("Unable to find user")
	} else {
		gqlReturn, err := tf.DBUserToGQLUser(userToFind)

		return gqlReturn, err
	}

	return nil, errors.New("not authenticated for query")
}

//func getUserProfile(r *queryResolver, user *dbm.User) (*gqlModels.User, error) {
//	var userProfile dbm.User
//	db := r.ORM.DB.New()
//
//	db.Where("user_id = ?", &user.ID).First(&userProfile)
//
//	gqlUserProfile, err := tf.DBUserProfileToGQLUserProfile(&userProfile, user)
//
//	if err != nil {
//		return &gqlModels.User{}, err
//	}
//
//	return gqlUserProfile, nil
//
//}

func updateUserProfile(r *mutationResolver, input gqlModels.UpdatedProfile) (*gqlModels.UserProfile, error) {

	tx := r.ORM.DB.New().Begin()

	dbo , err := tf.GQLUpdateUserProfileToDBUserProfile(&input, tx)

	if err != nil {
		return nil, err
	}

	tx = tx.Save(&dbo)

	if tx.Error != nil {
		return nil, errors.New("unable to update user profile")
	}

	gql, err := tf.UpdatedDBUserProfileToGQLUserProfile(dbo, input.Email)

	orm.CommitOrRollBackIfErrorAndCloseSession(tx, err)

	return gql, err
}
