package resolvers

import (
	"context"
	"errors"
	authhandler "github.com/fedoratipper/bitkiosk/server/internal/authentication"
	"github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	tf "github.com/fedoratipper/bitkiosk/server/internal/gql/resolvers/transformations"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	dbm "github.com/fedoratipper/bitkiosk/server/internal/orm/models"
)

func (r *mutationResolver) UpdateUserProfile(ctx context.Context, input models.UpdatedProfile) (*models.UserProfile, error) {
	panic("implement me")
}

func (r *queryResolver) UserProfile(ctx context.Context, userId *int) (*models.UserProfile, error) {
	authLevel, err := authhandler.GetAuthLevelFromJWT(ctx)

	if authLevel == nil {
		return nil, errors.New("not authenticated for query")
	}

	if authLevel.AuthLevel == authhandler.UserAuth && *userId == int(authLevel.UID) {
		return getUserProfile(r, authLevel.UID), nil
	} else if authLevel.AuthLevel == authhandler.AdminAuth	{
		return getUserProfile(r, authLevel.UID), nil
	} else {
		if err != nil {
			logger.Error("Unable to resolve auth level with Users resolver. \n" + err.Error())
		}
	}

	return nil, errors.New("not authenticated for query")
}

func createUserProfile() {

}


func getUserProfile(r *queryResolver, uid uint) (*models.UserProfile, error) {
	var userProfile dbm.UserProfile
	db := r.ORM.DB.New()

	db.Where("user_id = ?", uid).First(&userProfile)
	db.Close()

	if userProfile.ID == 0 {
		return nil, errors.New("unable to find user profile")
	}

	for _, dbRec := range dbRecords {
		if rec, err := tf.DBUserToGQLUser(&dbRec); err != nil {
			logger.Errorfn("userProfile", err)
		} else {
			res = append(res, rec)
		}
	}

	return res, nil

}
