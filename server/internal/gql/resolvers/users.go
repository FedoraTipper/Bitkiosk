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
)

// CreateUser creates a record
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (string, error) {
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

	if authLevel != nil && authLevel.AuthLevel == session.AdminAuth {
		return userList(r)
	} else {
		if err != nil {
			logger.Error("Unable to resolve auth level with Users resolver. \n" + err.Error())
		}
	}

	return nil, errors.New("not authenticated for query")
}

// ## Helper functions
func userCreate(r *mutationResolver, input models.NewUser) (string, error) {
	var gqlReturn string

	userDbo, err := tf.GQLInputUserToDBUser(&input, false)

	if err != nil {
		return "", err
	}

	userDbo.Role = session.UserAuth

	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()

	db, err = userDbo.Create(db)

	if err == nil {
		tokenDigest := digest.GetDigest(input.Token, uint(input.AuthMethodID))

		db, err = (&dbm.AuthenticationMatrix{UserID: userDbo.ID, AuthMethodID: uint(input.AuthMethodID), Token: tokenDigest}).Create(db)

		if err == nil {
			userProfileDbo := &dbm.UserProfile{
				UserID:      userDbo.ID,
				FirstName:   input.FirstName,
				LastName:    input.LastName,
			}

			db, err = userProfileDbo.Create(db)

			if err == nil {
				gqlReturn = "success"
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
	//TODO: Comeback to this.
	entity := "users"
	var res []*models.User

	db := r.ORM.DB.New()

	var dbRecords = []dbm.User{}
	db.Find(&dbRecords)

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

