package resolvers

import (
	"context"
	"errors"
	authhandler "github.com/fedoratipper/bitkiosk/server/internal/authentication"
	"github.com/fedoratipper/bitkiosk/server/internal/authentication/session"
	"github.com/fedoratipper/bitkiosk/server/internal/digest"
	"github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	tf "github.com/fedoratipper/bitkiosk/server/internal/gql/resolvers/transformations"
	logger "github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/fedoratipper/bitkiosk/server/internal/orm"
	dbm "github.com/fedoratipper/bitkiosk/server/internal/orm/models"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils/date"
	"time"
)

// CreateUser creates a record
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	return userCreate(r, input)
}

// UpdateUser updates a record
func (r *mutationResolver) UpdateUser(ctx context.Context, id uint, input models.NewUser) (*models.User, error) {
	return userUpdate(r, input, id)
}

// DeleteUser deletes a record
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	return userDelete(r, id)
}

// Users lists records
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	authLevel, err := authhandler.GetAuthLevelFromSession(ctx)

	if authLevel.AuthLevel == session.AdminAuth {
		return userList(r)
	} else {
		if err != nil {
			logger.Error("Unable to resolve auth level with Users resolver. \n" + err.Error())
		}
	}

	return nil, errors.New("not authenticated for query")
}

// ## Helper functions
func userCreate(r *mutationResolver, input models.NewUser) (*models.User, error) {
	var gqlReturn *models.User

	userDbo, err := tf.GQLInputUserToDBUser(&input, false)
	if err != nil {
		return nil, err
	}

	userDbo.Role = session.UserAuth

	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()

	err = userDbo.BeforeCreate(db)

	if err != nil {
		db, err = orm.CreateObject(userDbo, userDbo, db)

		if err == nil {
			gqlReturn, _ = tf.DBUserToGQLUser(userDbo)

			tokenDigest := digest.GetDigest(input.Token, uint(input.AuthMethodID))

			authenticationMatrixDbo := &dbm.AuthenticationMatrix{UserID: userDbo.ID, AuthMethodID: uint(input.AuthMethodID), Token: tokenDigest}

			db, err = orm.CreateObject(authenticationMatrixDbo, authenticationMatrixDbo, db)

			var dateOfBirth *time.Time
			dateOfBirth, err = date.ParseSqlDate(*input.DateOfBirth)

			if err == nil {
				userProfileDbo := &dbm.UserProfile{
					UserID:      userDbo.ID,
					FirstName:   input.FirstName,
					LastName:    input.LastName,
					DateOfBirth: dateOfBirth,
				}

				db, err = orm.CreateObject(userProfileDbo, userProfileDbo, db)

				if err == nil {
					if gqlUserProfile, err := tf.DBUserProfileToGQLUserProfile(userProfileDbo); err == nil {
						gqlReturn.UserProfile = gqlUserProfile
					}
				}
			}
		}
	}

	err = orm.CommitOrRollBackIfError(db, err)

	return gqlReturn, err
}

func userUpdate(r *mutationResolver, input models.NewUser, ids ...uint) (*models.User, error) {
	dbo, err := tf.GQLInputUserToDBUser(&input, false)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()

	var errBeforeUpdate = dbo.BeforeUpdate(db)

	if errBeforeUpdate != nil {
		return nil, errBeforeUpdate
	}

	db = db.Model(&dbo).Update(dbo).First(dbo) // Or update it

	gql, err := tf.DBUserToGQLUser(dbo)

	if err != nil {
		return nil, err
	}
	db = db.Commit()

	return gql, db.Error
}

func userDelete(r *mutationResolver, id string) (bool, error) {
	return false, nil
}

func userList(r *queryResolver) ([]*models.User, error) {
	entity := "users"
	var res []*models.User

	db := r.ORM.DB.New()

	var dbRecords = []dbm.User{}
	db.Find(&dbRecords)

	db.Close()

	if dbRecords == nil {
		return nil, nil
	}

	for _, dbRec := range dbRecords {
		if rec, err := tf.DBUserToGQLUser(&dbRec); err != nil {
			logger.Errorfn(entity, err)
		} else {
			res = append(res, rec)
		}
	}

	return res, nil
}

